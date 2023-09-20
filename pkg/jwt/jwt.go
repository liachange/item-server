package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
	"item-server/pkg/app"
	"item-server/pkg/config"
	"item-server/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrHeaderEmpty            = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        = errors.New("请求头中 Authorization 格式有误")
	ErrTokenSignatureInvalid  = errors.New("令牌签名无效")
)

type JWT struct {
	SignKey    []byte        // 秘钥
	MaxRefresh time.Duration // 刷新token的最大过期时间
}

// UserInfo 自定义用户信息
type UserInfo struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

// CustomJWTClaims 自定义Payload信息
type CustomJWTClaims struct {
	UserInfo
	ExpireAtTime int64 `json:"expire_time"` // 过期时间

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtpkg.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.Get("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

// ParserToken 解析 Token，中间件中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*CustomJWTClaims, error) {
	// 从header获取token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 1.解析token
	token, err := jwt.parseTokenString(tokenString)

	// 2. 解析出错
	if err != nil {
		if errors.Is(err, jwtpkg.ErrTokenExpired) {
			return nil, ErrTokenExpired
		} else if errors.Is(err, jwtpkg.ErrTokenMalformed) {
			return nil, ErrTokenMalformed
		} else if errors.Is(err, jwtpkg.ErrTokenSignatureInvalid) {
			return nil, ErrTokenSignatureInvalid
		} else {
			return nil, ErrTokenInvalid
		}
	}

	if claims, ok := token.Claims.(*CustomJWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 1. 从 Header 里获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		if !errors.Is(err, jwtpkg.ErrTokenExpired) {
			return "", err
		}
	}

	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*CustomJWTClaims)

	// 5. 检查是否过了『最大允许刷新的时间』
	t := app.TimeNowInTimezone().Add(-jwt.MaxRefresh).Unix()
	// 首次签名时间 > (当前时间 - 最大允许刷新时间)
	if claims.IssuedAt.Unix() > t {
		claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(jwt.expireAtTime())
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken 生成  Token，在登录成功时调用
func (jwt *JWT) IssueToken(info UserInfo) string {
	// 构造自定义Payload信息
	expireTime := jwt.expireAtTime()
	claims := CustomJWTClaims{
		// 用户信息
		UserInfo: UserInfo{
			UserID:   info.UserID,
			UserName: info.UserName,
		},
		// 过期时间
		ExpireAtTime: expireTime.Unix(),
		RegisteredClaims: jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(app.TimeNowInTimezone()), // 签名生效时间
			IssuedAt:  jwtpkg.NewNumericDate(app.TimeNowInTimezone()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtpkg.NewNumericDate(expireTime),              // 签名过期时间
			Issuer:    config.GetString("app.name"),                   // 签名颁发者
		},
	}

	// 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims CustomJWTClaims) (string, error) {
	// 使用HS256算法进行token生成
	t := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return t.SignedString(jwt.SignKey)
}

// token过期时间
func (jwt *JWT) expireAtTime() time.Time {
	timezone := app.TimeNowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return timezone.Add(expire)
}

// getTokenFromHeader 使用 jwtpkg.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

// parseTokenString 解析token
func (jwt *JWT) parseTokenString(token string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(token, &CustomJWTClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

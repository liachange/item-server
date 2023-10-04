package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 中间件，跨域问题
func Cors() gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins:  []string{"http://localhost:3010", "http://localhost:3000", "https://www.baidu.com"},
		AllowMethods:  []string{"PUT", "POST", "DELETE", "GET"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "token", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}
	return cors.New(config)
}

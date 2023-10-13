package helpers

import (
	"crypto/rand"
	"fmt"
	"github.com/spf13/cast"
	"io"
	mathrand "math/rand"
	"reflect"
	"time"
	"unsafe"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

// RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	src := mathrand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))

	//letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//b := make([]byte, length)
	//for i := range b {
	//	b[i] = letters[mathrand.Intn(len(letters))]
	//}
	//return string(b)

}

// TimeNow 获取系统时间
func TimeNow() time.Time {
	return cast.ToTime(time.Now().Format("2006-01-02 15:04:05"))
}

// TimeUnix 时间字符串转为时间戳
func TimeUnix(str string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	return tt.Unix()
}

// TimeStr 时间戳转时间字符串
func TimeStr(unix int64, format string) (formatTimeStr string) {
	switch format {
	case "day":
		formatTimeStr = time.Unix(unix, 0).Format("2006-01-02")
		return
	case "second":
		formatTimeStr = time.Unix(unix, 0).Format("2006-01-02 15:04:05")
		return
	default:
		formatTimeStr = cast.ToString(unix)
	}
	return
}

// TimeFormat 时间格式化
func TimeFormat(str time.Time, format string) (formatTimeStr string) {
	switch format {
	case "day":
		formatTimeStr = str.Format("2006-01-02")
		return
	case "second":
		formatTimeStr = str.Format("2006-01-02 15:04:05")
		return
	default:
		formatTimeStr = cast.ToString(str)
	}
	return
}

// TimeParse 时间格式验证
func TimeParse(str string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", str)
	if err == nil {
		return true
	}
	return false
}

// ReqSelect 更新过滤
func ReqSelect(obj interface{}) []string {
	str := make([]string, 0)
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tagValue := t.Field(i).Tag.Get("select")
		if tagValue != "" {
			value := v.Field(i).Interface()
			if !Empty(value) {
				str = append(str, tagValue)
			}
		}
	}
	fmt.Println(str)
	return str
}

func ReqFilter(req interface{}) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tagValue := t.Field(i).Tag.Get("filter")
		if tagValue != "" {
			value := v.Field(i).Interface()
			if !Empty(value) {
				m[tagValue] = value
			}
		}
	}
	fmt.Println(m)
	return m
}

func IdVerify(v uint64) bool {
	if len(cast.ToString(v)) >= 7 {
		return true
	}
	return false
}

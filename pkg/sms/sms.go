package sms

import (
	"item-server/pkg/config"
	"sync"
)

// Message 是短信的结构体
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// Response 服务商返回状态
type Response struct {
	Code          string
	Message       string
	BizId         string
	RequestId     string
	Phone         string
	SignName      string
	TemplateCode  string
	TemplateParam string
}

// SMS 是我们发送短信的操作类
type SMS struct {
	Driver Driver
}

// once 单例模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})

	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}

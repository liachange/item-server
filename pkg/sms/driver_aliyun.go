package sms

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"item-server/pkg/logger"
)

// Aliyun 实现 sms.Driver interface
type Aliyun struct{}

// Send 实现 sms.Driver interface 的 Send 方法
func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	client, err := CreateClient(
		tea.String(config["access_key_id"]),
		tea.String(config["access_key_secret"]),
	)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}

	logger.DebugJSON("短信[阿里云]", "配置信息", config)

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}

	// 发送参数
	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(config["sign_name"]),
		TemplateCode:  tea.String(message.Template),
		TemplateParam: tea.String(string(templateParam)),
	}
	// 其他运行参数
	runtime := &util.RuntimeOptions{}

	result, err := client.SendSmsWithOptions(sendSmsRequest, runtime)

	//失败
	if err != nil {
		var errs = &tea.SDKError{}
		if _t, ok := err.(*tea.SDKError); ok {
			errs = _t
		} else {
			errs.Message = tea.String(err.Error())
		}

		var r dysmsapi.SendSmsResponseBody
		err = json.Unmarshal([]byte(*errs.Data), &r)
		logger.LogIf(err)

		logger.ErrorString("短信[阿里云]", "发信失败", err.Error())
		return false
	}
	logger.DebugJSON("短信[阿里云]", "请求内容", sendSmsRequest)
	logger.DebugJSON("短信[阿里云]", "接口响应", result)
	resultJSON, err := json.Marshal(result.Body)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析响应 JSON 错误", err.Error())
		return false
	}
	response := Response{
		Phone:         phone,
		SignName:      config["sign_name"],
		TemplateCode:  message.Template,
		TemplateParam: string(templateParam),
		RequestId:     tea.StringValue(result.Body.RequestId),
		BizId:         tea.StringValue(result.Body.BizId),
		Code:          tea.StringValue(result.Body.Code),
		Message:       tea.StringValue(result.Body.Message),
	}
	if response.Code == "OK" {
		logger.DebugString("短信[阿里云]", "发信成功", "")
		return true
	} else {
		logger.ErrorString("短信[阿里云]", "服务商返回错误", string(resultJSON))
		return false
	}
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi.Client, _err error) {
	config := &openapi.Config{
		// AccessKey ID
		AccessKeyId: accessKeyId,
		// AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}

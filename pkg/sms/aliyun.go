package sms

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	aliyunUtil "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 可以上官网查看示例 https://next.api.aliyun.com/api/Dysmsapi/2017-05-25/SendSms?params={}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

// SendVerifyCode
func SendSms(req dysmsapi20170525.SendSmsRequest) (_err error) {
	// TODO your key，from config
	client, _err := CreateClient(tea.String("LTAI5t5rnGb9eAXXVDWsJSjR"), tea.String("XY10yYM57gLMLYWg7NWOSbKPOZWA2p"))
	if _err != nil {
		return _err
	}

	defer func() {
		if r := tea.Recover(recover()); r != nil {
			_err = r
		}
	}()

	runtime := &aliyunUtil.RuntimeOptions{}
	result, _err := client.SendSmsWithOptions(&req, runtime)
	if _err != nil {
		return _err
	}

	if *result.Body.Code != "OK" {
		_err = errors.New(result.String())
		return
	}

	return _err
}

func getVerifyCodeReq(phoneNumber, code string) (req dysmsapi20170525.SendSmsRequest) {
	// TODO SignName TemplateCode
	return dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("AkitaPlanet"),
		TemplateCode:  tea.String("SMS_465345782"),
		PhoneNumbers:  tea.String(phoneNumber),
		TemplateParam: tea.String(`{"applet_code":"` + code + `"}`),
	}
}

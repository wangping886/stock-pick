package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendSms() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAIoP845ZKWXwZ4", "qweLCjAM90IdBuUB2lq2xQ7TzSceYq")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = "15811269507"
	request.TemplateParam = "{\"code\":\"2222\"}"
	request.TemplateCode = "SMS_180961335"
	request.SignName = "stockpick"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(11, err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

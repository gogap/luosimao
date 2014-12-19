package main

import (
	"encoding/json"
	"fmt"

	"github.com/gogap/luosimao"
)

var voiceAuth luosimao.Authorization
var smsAuth luosimao.Authorization

func main() {

	voiceAuth = luosimao.Authorization{UserName: "api", Password: ""}
	smsAuth = luosimao.Authorization{UserName: "api", Password: ""}

	send_voice("13400000000", 123)
	send_sms("13400000000", "您的验证码是:1234 【公司签名】")
	voice_status()
	sms_status()
}

func send_voice(mobile string, code int32) {
	sender := luosimao.NewVoiceSender(voiceAuth, luosimao.JSON)
	resp, err := sender.Send(luosimao.VoiceRequest{Mobile: mobile, Code: code}, 1000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, _ := json.Marshal(resp)
		fmt.Println(string(b))
		fmt.Println(resp.ErrorDescription())
	}
}

func voice_status() {
	sender := luosimao.NewVoiceSender(voiceAuth, luosimao.JSON)
	resp, err := sender.Status(1000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, _ := json.Marshal(resp)
		fmt.Println(string(b))
		fmt.Println(resp.ErrorDescription())
	}
}

func send_sms(mobile, message string) {
	sender := luosimao.NewSMSSender(smsAuth, luosimao.JSON)
	resp, err := sender.Send(luosimao.SMSRequest{Mobile: mobile, Message: message}, 1000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, _ := json.Marshal(resp)
		fmt.Println(string(b))
		fmt.Println(resp.ErrorDescription())
	}
}

func sms_status() {
	sender := luosimao.NewSMSSender(smsAuth, luosimao.JSON)
	resp, err := sender.Status(1000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, _ := json.Marshal(resp)
		fmt.Println(string(b))
		fmt.Println(resp.ErrorDescription())
	}
}

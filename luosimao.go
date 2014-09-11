package luosimao

type ProtocalType int32

const (
	JSON  ProtocalType = 1
	XML   ProtocalType = 2
	JSONP ProtocalType = 3
)

const (
	SMSServerURL   = "https://sms-api.luosimao.com/v1/"
	VoiceServerURL = "http://voice-api.luosimao.com/v1/"
)

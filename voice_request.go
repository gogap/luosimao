package luosimao

type VoiceRequest struct {
	Mobile string `json:"mobile"`
	Code   int32  `json:"code,string"`
}

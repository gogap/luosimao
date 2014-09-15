package luosimao

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/parnurzeal/gorequest"
)

type VoiceSender struct {
	sendUrl   string
	statusUrl string
	auth      Authorization
	Proxy     string
}

func NewVoiceSender(auth Authorization, proto ProtocalType) *VoiceSender {
	sender := new(VoiceSender)
	switch proto {
	case JSON:
		{
			sender.sendUrl = VoiceServerURL + "verify.json"
			sender.statusUrl = VoiceServerURL + "status.json"
		}
	case XML:
		{
			sender.sendUrl = VoiceServerURL + "verify.xml"
			sender.statusUrl = VoiceServerURL + "status.xml"
		}
	case JSONP:
		{
			sender.sendUrl = VoiceServerURL + "verify.jsonp"
			sender.statusUrl = VoiceServerURL + "status.jsonp"
		}
	}
	sender.auth = auth

	return sender
}

func (p *VoiceSender) Send(req VoiceRequest, timeout int64) (response Response, err error) {
	to := time.Duration(timeout)

	request := gorequest.New().Timeout(to * time.Millisecond).Proxy(p.Proxy)

	strCode := fmt.Sprintf("%04d", req.Code)

	params := url.Values{}
	params.Add("mobile", req.Mobile)
	params.Add("code", strCode)

	_, body, errs := request.Post(p.sendUrl).Set("Authorization", p.auth.BasicAuthorization()).Set("Content-Type", "application/x-www-form-urlencoded").Send(params.Encode()).End()

	if errs != nil && len(errs) > 0 {
		err = errs[0]
	}

	if e := json.Unmarshal([]byte(body), &response); e != nil {
		err = e
		return
	}

	return
}

func (p *VoiceSender) Status(timeout int64) (status Status, err error) {
	to := time.Duration(timeout)

	request := gorequest.New().Timeout(to * time.Millisecond).Proxy(p.Proxy)

	_, body, errs := request.Post(p.statusUrl).Set("Authorization", p.auth.BasicAuthorization()).Set("Content-Type", "application/x-www-form-urlencoded").End()

	if errs != nil && len(errs) > 0 {
		err = errs[0]
	}

	if e := json.Unmarshal([]byte(body), &status); e != nil {
		err = e
		return
	}

	return
}

package luosimao

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/parnurzeal/gorequest"
)

type SMSSender struct {
	sendUrl   string
	statusUrl string
	auth      Authorization
	Proxy     string
}

func NewSMSSender(auth Authorization, proto ProtocalType) *SMSSender {
	sender := new(SMSSender)
	switch proto {
	case JSON:
		{
			sender.sendUrl = SMSServerURL + "send.json"
			sender.statusUrl = SMSServerURL + "status.json"
		}
	case XML:
		{
			sender.sendUrl = SMSServerURL + "send.xml"
			sender.statusUrl = SMSServerURL + "status.xml"
		}
	}
	sender.auth = auth

	return sender
}

func (p *SMSSender) Send(req SMSRequest, timeout int64) (response Response, err error) {
	to := time.Duration(timeout)

	request := gorequest.New().Timeout(to * time.Millisecond).Proxy(p.Proxy)

	params := url.Values{}
	params.Add("mobile", req.Mobile)
	params.Add("message", req.Message)

	_, body, errs := request.Post(p.sendUrl).Set("Authorization", p.auth.BasicAuthorization()).Set("Content-Type", "application/x-www-form-urlencoded").Send(params.Encode()).End()

	if err = errors_to_error(errs); err != nil {
		return
	}

	if e := json.Unmarshal([]byte(body), &response); e != nil {
		err = e
		return
	}

	return
}

func (p *SMSSender) Status(timeout int64) (status Status, err error) {
	to := time.Duration(timeout)

	request := gorequest.New().Timeout(to * time.Millisecond).Proxy(p.Proxy)

	_, body, errs := request.Post(p.statusUrl).Set("Authorization", p.auth.BasicAuthorization()).Set("Content-Type", "application/x-www-form-urlencoded").End()

	if err = errors_to_error(errs); err != nil {
		return
	}

	if e := json.Unmarshal([]byte(body), &status); e != nil {
		err = e
		return
	}

	return
}

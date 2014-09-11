package luosimao

import (
	"encoding/base64"
)

type Authorization struct {
	UserName string
	Password string
}

func (p *Authorization) BasicAuthorization() string {
	credential := base64.StdEncoding.EncodeToString([]byte(p.UserName + ":" + p.Password))
	return "Basic " + credential
}

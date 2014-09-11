package luosimao

import (
	"encoding/xml"
)

type Response struct {
	XMLName xml.Name  `json:"-" xml:"response"`
	Error   ErrorCode `json:"error" xml:"error"`
	Message string    `json:"message" xml:"msg"`
}

func (p *Response) ErrorDescription() string {
	return GetErrorDescription(p.Error)
}

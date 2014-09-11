package luosimao

import (
	"encoding/xml"
)

type Status struct {
	XMLName xml.Name  `json:"-" xml:"response"`
	Error   ErrorCode `json:"error" xml:"error"`
	Deposit int64     `json:"deposit,string" xml:"deposit"`
}

func (p *Status) ErrorDescription() string {
	return GetErrorDescription(p.Error)
}

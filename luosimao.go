package luosimao

import (
	"errors"
	"fmt"
)

type ProtocalType int32

const (
	JSON ProtocalType = 1
	XML  ProtocalType = 2
)

const (
	SMSServerURL   = "https://sms-api.luosimao.com/v1/"
	VoiceServerURL = "http://voice-api.luosimao.com/v1/"
)

func errors_to_error(errs []error) error {
	if errs == nil || len(errs) == 0 {
		return nil
	}

	errString := ""
	for i, e := range errs {
		errString = errString + fmt.Sprintf("error %d: %s \n", i, e.Error())
	}

	return errors.New(errString)
}

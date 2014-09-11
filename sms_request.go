package luosimao

type SMSRequest struct {
	Mobile   string `json:"mobile"`
	Message  string `json:"message,string"`
	CallBack string `json:"cb,omitempty"`
}

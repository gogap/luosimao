package luosimao

type ErrorCode int32

const (
	SMS_ERROR_OK                     ErrorCode = 0
	SMS_ERROR_VALIDATE_ERROR         ErrorCode = -10
	SMS_ERROR_BALANCE_NOT_ENOUGH     ErrorCode = -20
	SMS_ERROR_EMPTY_MESSAGE          ErrorCode = -30
	SMS_ERROR_CONTENT_SENSITIVE_WORD ErrorCode = -31
	SMS_ERROR_NO_SIGNATURE           ErrorCode = -32
	SMS_ERROR_WRONG_MOBILE_NUMBER    ErrorCode = -40
	SMS_ERROR_IP_NOT_IN_WHITE_LIST   ErrorCode = -50
)

var (
	errorDescriptions map[ErrorCode]string
)

func init() {
	errorDescriptions = make(map[ErrorCode]string)

	errorDescriptions[SMS_ERROR_OK] = "成功"
	errorDescriptions[SMS_ERROR_VALIDATE_ERROR] = "验证信息失败"
	errorDescriptions[SMS_ERROR_BALANCE_NOT_ENOUGH] = "余额不足"
	errorDescriptions[SMS_ERROR_EMPTY_MESSAGE] = "内容为空"
	errorDescriptions[SMS_ERROR_CONTENT_SENSITIVE_WORD] = "内容存在敏感词"
	errorDescriptions[SMS_ERROR_NO_SIGNATURE] = "内容缺少签名信息"
	errorDescriptions[SMS_ERROR_WRONG_MOBILE_NUMBER] = "错误的手机号"
	errorDescriptions[SMS_ERROR_IP_NOT_IN_WHITE_LIST] = "请求发送IP不在白名单内"
}

func GetErrorDescription(errCode ErrorCode) string {
	if err, exist := errorDescriptions[errCode]; exist {
		return err
	}
	return "未知错误码"
}

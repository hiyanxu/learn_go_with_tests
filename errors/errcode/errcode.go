package errcode

import "fmt"

// 定义错误码
const (
	ERR_CODE_OK             = 0 // OK
	ERR_CODE_INVALID_PARAMS = 1 // 无效参数
	ERR_CODE_TIMEOUT        = 2 // 超时
)

// 定义错误码与描述信息的map
var mapErrDesc = map[int]string{
	ERR_CODE_OK:             "OK",
	ERR_CODE_INVALID_PARAMS: "无效参数",
	ERR_CODE_TIMEOUT:        "超时",
}

// 根据错误码获取描述信息
func GetDesc(errCode int) string {
	if desc, ok := mapErrDesc[errCode]; ok {
		return desc
	}

	return fmt.Sprintf("error code: %d", errCode)
}

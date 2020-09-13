package errcode

import "fmt"

/**
当前写法的缺点：
1、错误描述信息需要写两次：注释一次、map里面一次
*/

type ErrCode int // 对错误code专门定义类型

//go:generate stringer -type ErrCode -linecomment
const (
	ERR_CODE_OK             ErrCode = 0 // OK
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 无效参数
	ERR_CODE_TIMEOUT        ErrCode = 2 // 超时
)

// 定义错误码与描述信息的map
var mapErrDesc = map[ErrCode]string{
	ERR_CODE_OK:             "OK",
	ERR_CODE_INVALID_PARAMS: "无效参数",
	ERR_CODE_TIMEOUT:        "超时",
}

// 根据错误码获取描述信息
func GetDesc(errCode ErrCode) string {
	if desc, ok := mapErrDesc[errCode]; ok {
		return desc
	}

	return fmt.Sprintf("error code: %d", errCode)
}

// 对ErrCode类型定义String方法，返回字符串，不用再调用GetDesc
//func (e ErrCode) String() string {
//	return GetDesc(e)
//}

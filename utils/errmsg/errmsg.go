package errmsg

// code = 1000   用户模块的错误
// code = 2000   文章模块的错误
// code = 3000   分类模块的错误
const (
	SUCCESS                = 200
	ERROR                  = 500
	ErrorUsernameUsed      = 1001
	ErrorPasswordWorng     = 1002 //用户密码错误
	ErrorUserNotExist      = 1003 //用户不存在
	ErrorTokenExist        = 1004 //用户携带token不存在
	ErrorTokenRuntime      = 1005 // token超时
	ErrorTokenWrong        = 1006 // token不合法
	ErrorTokenTypeWrong    = 1007
	ErrorUserNopermissions = 1008 // 没有权限
	ErrorArtNotExist       = 2001
	ErrorCateNameUsed      = 3001
	ErrorCateNotExit       = 3002
	ErrDingUser            = 4002
)

// 字典

var Codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ErrorUsernameUsed:      "用户名已存在!",
	ErrorPasswordWorng:     "用户密码错误!",
	ErrorUserNotExist:      "用户不存在!",
	ErrorUserNopermissions: "用户无权限",
	ErrorTokenExist:        "Token不存在!",
	ErrorTokenRuntime:      "Token已过期!",
	ErrorTokenWrong:        "Token不合法/格式错误!",
	ErrorTokenTypeWrong:    "Token格式错误!",
	ErrorArtNotExist:       "文章不存在",
	ErrorCateNameUsed:      "Tags分类已存在!",
	ErrorCateNotExit:       "查询不到该分类！",
	ErrDingUser:            "钉钉 ACCESS_TOKEN 获取失败！",
}

// 输出错误信息

func GetErrMsg(code int) string {
	return Codemsg[code]
}

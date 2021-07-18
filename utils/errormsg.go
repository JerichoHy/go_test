package utils

const (
	SUCCESS        = 200
	ERROR          = 500
	IS_EXIST       = 10001
	IS_NOT_EXIST   = 10002
	WRONG_PASSWORD = 10003
	NOT_LOGIN      = 10004
	SESSION_ERROR  = 10005
)

var codeMsg = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	IS_EXIST:       "is exist",
	IS_NOT_EXIST:   "is not exist",
	WRONG_PASSWORD: "wrong password",
	NOT_LOGIN:      "not login",
	SESSION_ERROR:  "session ERROR",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

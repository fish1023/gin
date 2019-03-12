package retcode

var msg = map[int]string{
	SUCCESS:           "ok",
	NETWORK_EXCEPTION: "网络异常，请重试",
	SYSTEM_EXCEPTION:  "系统异常，请重试",
	DB_EXCEPTION:      "数据库异常，请重试",
	PARAM_ERROR:       "参数错误",
	UPDATE_FAILED:     "更新失败",
	PROCESS_LOCK:      "操作频繁请稍后再试",
	ILLEGAL_REQUEST:   "非法请求",
    SPAM_ACTION:       "操作太频繁，请稍后再试",
    CODE_EMPTY:        "请重新获取验证码",
}

func GetMsg(code int) string {
	msg, ok := msg[code]
	if ok {
		return msg
	}
	return string(msg[NETWORK_EXCEPTION])
}

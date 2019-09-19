package e

//map 是引用类型，可以使用如下声明：（[keytype] 和 valuetype 之间允许有空格，但是 gofmt 移除了空格）

//var map1 map[keytype]valuetype
//var map1 map[string]int

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int, msg string) string {

	if msg != "" {
		return msg
	}

	// value, ok := map1[key1] // 如果key1存在则ok == true，否则ok为false
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

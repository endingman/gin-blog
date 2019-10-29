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
	// 保存图片失败
	ERROR_UPLOAD_SAVE_IMAGE_FAIL: "保存图片失败",
	// 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL: "检查图片失败",
	// 校验图片错误，图片格式或大小有问题
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",

	ERROR_CHECK_EXIST_ARTICLE_FAIL: "检查文章是否存在失败",
	ERROR_GET_ARTICLE_FAIL:         "获取单个文章失败",
	ERROR_GET_TAG_FAIL:             "获取单个标签失败",
	ERROR_GET_TAGS_FAIL:            "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:           "统计标签失败",
	ERROR_EXPORT_TAG_FAIL:          "导出标签失败",
	ERROR_EXIST_TAG_FAIL:           "获取已存在标签失败",
	ERROR_DELETE_TAG_FAIL:          "删除标签失败",
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

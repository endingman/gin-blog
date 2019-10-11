package util

import (
	"gin-blog/pkg/setting"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("Page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}

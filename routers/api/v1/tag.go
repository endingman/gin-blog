package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/export"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"gin-blog/service/tag_service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//这里使用的并不是原来的名字，使用更加简洁的resful api
//类似laravel里面的接口方法定义
func GetTags(c *gin.Context) {
	var msg string
	// c.Query可用于获取?name=test&state=1这类URL参数，而c.DefaultQuery则支持设置一个默认值
	// code变量使用了e模块的错误编码，这正是先前规划好的错误码，方便排错和识别记录
	// util.GetPage保证了各接口的page处理是一致的
	// c *gin.Context是Gin很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的JSON和呈现JSON响应
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, msg, nil)
		return
	}

	count, err := tagService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, msg, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, msg, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

func AddTag(c *gin.Context) {
	var msg string
	appG := app.Gin{C: c}
	name := c.PostForm("name")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	createdBy := c.PostForm("created_by")

	// "github.com/astaxie/beego/validation"包使用
	// ~~~习惯了laravel的写法，go的验证太麻烦没有更好的写法或者更一目了然跟业务分离的写法吗？
	// 没有跟laravel那么简洁的验证，现在的写法相当于是把验证跟业务路基放在一起了，没有分离request跟业务
	// 不知道有没有更好的验证，如果有人看到这个东西有更好的话，希望告知一下
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	if exist, _ := models.ExistTagByName(name); !exist {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, msg, nil)
		return
	}

	_ = models.AddTag(name, state, createdBy)

	appG.Response(http.StatusOK, e.SUCCESS, msg, make(map[string]string))

}

func UpdateTag(c *gin.Context) {
	var msg string
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	state := -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	if exist, _ := models.ExistTagByID(id); !exist {
		appG.Response(http.StatusOK, e.SUCCESS, msg, nil)
		return

	}

	data := make(map[string]interface{})
	data["modified_by"] = modifiedBy
	if name != "" {
		data["name"] = name
	}
	if state != -1 {
		data["state"] = state
	}
	_ = models.EditTag(id, data)

	appG.Response(http.StatusOK, e.SUCCESS, msg, make(map[string]string))
}

func ShowTag(c *gin.Context) {
	var msg string
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("Id 必须大于 1")

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	if exist, _ := models.ExistArticleByID(id); !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, msg, nil)
		return
	}

	tag, err := models.GetTag(id)

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, msg, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, msg, tag)
}

func DestroyTag(c *gin.Context) {
	var msg string
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	tagService := tag_service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, msg, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, msg, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, msg, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, msg, nil)
}

func ExportTag(c *gin.Context) {
	var msg string
	appG := app.Gin{C: c}
	name := c.PostForm("name")
	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:  name,
		State: state,
	}

	filename, err := tagService.Export()
	if err != nil {
		logging.Error(err)
		appG.Response(http.StatusOK, e.ERROR_EXPORT_TAG_FAIL, msg, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, msg, map[string]string{
		"export_url":      export.GetExcelFullUrl(filename),
		"export_save_url": export.GetExcelPath() + filename,
	})
}

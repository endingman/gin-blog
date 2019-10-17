package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"gin-blog/service/article_service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//这里使用的并不是原来的名字，使用更加简洁的resful api
//类似laravel里面的接口方法定义
func GetArticles(c *gin.Context) {
	var msg string
	appG := app.Gin{c}

	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}

	state := -1
	// 链接参数条件state
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	tagId := -1
	// 链接参数条件tag_id
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	// 数据获取
	data["lists"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	appG.Response(http.StatusOK, e.SUCCESS, msg, data)

}

func GetArticle(c *gin.Context) {
	var msg string
	appG := app.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("Id 必须大于 1")

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, msg, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, msg, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, msg, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, msg, article)
}

func AddArticle(c *gin.Context) {
	var msg string
	appG := app.Gin{c}

	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	if exist, _ := models.ExistTagByID(tagId); !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, msg, nil)
		return
	}

	data := make(map[string]interface{})
	data["tag_id"] = tagId
	data["title"] = title
	data["desc"] = desc
	data["content"] = content
	data["created_by"] = createdBy
	data["state"] = state

	models.AddArticle(data)

	appG.Response(http.StatusOK, e.SUCCESS, msg, make(map[string]interface{}))
}

func UpdateArticle(c *gin.Context) {
	var msg string
	appG := app.Gin{c}
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	// PostForm===》post的body传值
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")

	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	if exist, _ := models.ExistArticleByID(id); !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, msg, nil)
	}
	if exist, _ := models.ExistTagByID(tagId); !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, msg, nil)
	}

	data := make(map[string]interface{})

	data["tag_id"] = tagId //过了验证ID必定大于0

	if title != "" {
		data["title"] = title
	}
	if desc != "" {
		data["desc"] = desc
	}
	if content != "" {
		data["content"] = content
	}

	data["modified_by"] = modifiedBy
	models.EditArticle(id, data)

	appG.Response(http.StatusOK, e.SUCCESS, msg, make(map[string]interface{}))
}

func DestroyArticle(c *gin.Context) {
	var msg string
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("Id必须大于0")

	if valid.HasErrors() {
		msg = app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, msg, nil)
		return
	}

	if exist, _ := models.ExistArticleByID(id); !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, msg, nil)
	}

	models.DeleteArticle(id)
	appG.Response(http.StatusOK, e.SUCCESS, msg, nil)
}

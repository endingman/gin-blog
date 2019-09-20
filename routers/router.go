package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PATCH("/tags/:id", v1.UpdateTag)
		//获取指定标签
		apiv1.GET("/tags/:id", v1.ShowTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DestoryTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PATCH("/articles/:id", v1.UpdateArticle)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DestoryArticle)
	}

	return r
}

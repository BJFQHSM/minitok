package main

import (
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/handler"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/gin-gonic/gin"
)

func main() {

	rpc.Init()
	r := gin.New()

	r.Use(logger())

	douyin := r.Group("/douyin")
	user1 := douyin.Group("/user")
	user1.POST("/login/", handler.Login)
	user1.POST("/register/", handler.Register)
	user1.GET("/", handler.QueryUserInfo)

	publish1 := douyin.Group("/publish")
	publish1.POST("/action/", handler.PublishAction)
	publish1.GET("/list/", handler.QueryPublishList)

	favorite1 := douyin.Group("/favorite")
	favorite1.POST("/action/", handler.FavoriteAction)
	favorite1.GET("/list/", handler.QueryFavoriteList)

	comment1 := douyin.Group("/comment")
	comment1.POST("/action/", handler.CommentAction)
	comment1.GET("/list/", handler.QueryCommentList)

	feed1 := douyin.Group("/feed")
	feed1.GET("", handler.Feed)
	//feed1.GET("/", handler.Feed)

	relation1 := douyin.Group("/relation")
	relation1.POST("/action/", handler.RelationAction)
	relation1.GET("/follow/list/", handler.QueryFollowList)
	relation1.GET("/follower/list/", handler.QueryFollowerList)

	douyin.GET("/static/", handler.Video)

	if err := http.ListenAndServe(":8080", r); err != nil {
		util.LogFatalf("API bind error, err = %+v\n", err)
	}
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求处理
		path := c.Request.URL.RequestURI()
		log.Println(path)
		c.Next()
	}
}

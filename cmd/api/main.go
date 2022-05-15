package main

import (
	"github.com/bytedance2022/minimal_tiktok/cmd/api/handler"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//

func main() {
	rpc.Init()
	r := gin.New()

	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.POST("/login", handler.Login)
	user1.POST("/register", handler.Register)
	user1.POST("", handler.QueryInfo)

	publish1 := v1.Group("/publish")
	publish1.POST("/action", handler.PublishAction)
	publish1.POST("/list", handler.QueryPublishList)

	favorite1 := v1.Group("/favorite")
	favorite1.POST("/action", handler.FavoriteAction)
	favorite1.POST("/list", handler.QueryFavoriteList)

	comment1 := v1.Group("/comment")
	comment1.POST("/action", handler.CommentAction)
	comment1.POST("/list", handler.QueryCommentList)

	feed1 := v1.Group("/feed")
	feed1.POST("", handler.Feed)

	relation1 := v1.Group("/relation")
	relation1.POST("/action", handler.RelationAction)
	relation1.POST("/follow/list", handler.QueryFollowList)
	relation1.POST("/follower/list", handler.QueryFollowerList)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

//func login(c *gin.Context) {
//	resp, err := rpc.UserClient.Login(context.Background(), &user.LoginRequest{})
//	if err != nil {
//		fmt.Printf("error %v", err)
//	}
//	fmt.Printf("%+v\n", resp)
//	c.JSON(http.StatusOK, resp)
//}

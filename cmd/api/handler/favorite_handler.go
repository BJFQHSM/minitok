package handler

import (
	"context"
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FavoriteAction godoc
// @Summary      like video action
// @Description
// @Tags         favorite
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Param        video_id body int true "video_id"
// @Param        action_type body int true "1 - like 2 - unlike"
// @Success      200 {object} biz.FavoriteActionResponse
// @Failure      500 {object} biz.FavoriteActionResponse
// @Router       /favorite/action [post]
func FavoriteAction(c *gin.Context) {
	var req biz.FavoriteActionRequest
	req.Token = c.Query("token")
	//根据token获取用户的id

	req.UserId = 1001
	tmp,err := strconv.Atoi(c.Query("video_id"))
	req.VideoId = int64(tmp)
	tmp,err = strconv.Atoi(c.Query("action_type"))
	req.ActionType = int32(tmp)
	fmt.Printf("req=%#v\n", req)
	if err != nil {
		// todo
		log.Printf("参数报错：%v:", err)
		var resp *biz.FavoriteActionResponse
		resp.StatusCode=1
		msg:="Params have errors!"
		resp.StatusMsg=&msg
		c.JSON(http.StatusBadRequest, resp)
	}
	resp, err := rpc.BizClient.FavoriteAction(context.Background(), &req)
	if err != nil {
		// todo
		log.Printf("点赞操作报错：%v:", err)
	}
	fmt.Println("点赞操作", resp)
	c.JSON(http.StatusOK, resp)
}

type ListResponse struct {
	VideoList  []*biz.Video `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
	StatusCode int32    `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string
}

// QueryFavoriteList godoc
// @Summary      get favorite list
// @Description  get favorite list by userId
// @Tags         favorite
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Success      200 {object} biz.QueryFavoriteListResponse
// @Failure      500 {object} biz.QueryFavoriteListResponse
// @Router       /favorite/list [get]
func QueryFavoriteList(c *gin.Context) {
	req := biz.QueryFavoriteListRequest{}
	tmp, err := strconv.Atoi(c.Query("user_id"))
	req.UserId = int64(tmp)
	req.Token = c.Query("token")
	if err != nil {
		// todo
		log.Printf("参数报错：%v:", err)
	}
	resp, err := rpc.BizClient.QueryFavoriteList(context.Background(), &req)
	if err != nil {
		// todo
		log.Printf("点赞列表报错：%v:", err)
	}
	res:=ListResponse{}
	res.StatusCode=resp.StatusCode
	res.VideoList=resp.VideoList
	c.JSON(http.StatusOK, res)
}

package handler

import (
	"context"
	"log"
	"net/http"

	"fmt"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	var req biz.FavoriteActionRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)
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
		resp.StatusCode = 1
		msg := "Params have errors!"
		resp.StatusMsg = &msg
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
	StatusCode int32        `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string
}

func QueryFavoriteList(c *gin.Context) {
	var req biz.QueryFavoriteListRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)

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
	c.JSON(http.StatusOK, resp)
}

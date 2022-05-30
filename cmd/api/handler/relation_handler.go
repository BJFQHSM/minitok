package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

// RelationAction godoc
// @Summary      follow and unfollow
// @Description  follow and unfollow
// @Tags         relation
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Param        to_user_id body int true "to_user_id"
// @Param        action_type body int true "1 - follow 2 - unfollow"
// @Success      200 {object} biz.RelationActionResponse
// @Failure      500 {object} biz.RelationActionResponse
// @Router       /relation/action [post]
func RelationAction(c *gin.Context) {
	var req biz.RelationActionRequest
	err := c.BindJSON(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.RelationAction(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

func QueryFollowList(c *gin.Context) {
	var req biz.QueryFollowListRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryFollowList(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

func QueryFollowerList(c *gin.Context) {
	var req biz.QueryFollowerListRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryFollowerList(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

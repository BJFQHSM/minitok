package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type FeedService interface {
	DoService() *biz.FeedResponse
}

func NewFeedService(ctx context.Context, r *biz.FeedRequest) FeedService {
	return &feedServiceImpl{Req: r, Ctx: ctx, Resp: &biz.FeedResponse{}}
}

type feedServiceImpl struct {
	Req  *biz.FeedRequest
	Resp *biz.FeedResponse
	Ctx  context.Context
}

func (s *feedServiceImpl) DoService() *biz.FeedResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.feed(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *feedServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.LatestTime < 0 {
		return errors.New("illegal params: latest_time cannot lower than 0")
	}
	return nil
}

func (s *feedServiceImpl) feed() error {
	latestTime := s.Req.LatestTime
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}

	//从数据库查找数据
	t := TimeStampToTime(latestTime)
	vdos, err := dal.QueryVideosByTime(t)
	if err != nil {
		return err
	}

	//获取下次的最新时间
	nextTime := time.Now().Unix()
	if len(vdos) > 0 {
		nextTime = vdos[len(vdos)-1].PublishDate.Unix()
	}

	videos := []*biz.Video{}
	for i := 0; i < len(vdos); i++ {
		videos = append(videos, MongoVdoToBizVdo(vdos[i], s.Req.UserIdFromToken))
	}

	s.Resp.VideoList = videos
	s.Resp.NextTime = &nextTime

	return nil
}

func (s *feedServiceImpl) buildResponse(err error) {
	if err != nil {
		errMsg := err.Error()
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 1
	} else {
		errMsg := "SUCCESS"
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 0
	}
}

//将video.go中的Video转化为biz.pb.go中的video类型
func MongoVdoToBizVdo(vdo *dal.Video, tokenId int64) *biz.Video {
	var err error
	res := &biz.Video{}
	res.Id = vdo.VideoId
	//查询当前登录用户信息
	user, err := dal.QueryUserByID(context.TODO(), tokenId)
	if err != nil {
		log.Println(err)
		return nil
	}
	res.Author, err = QueryUserInfoByUID(context.TODO(), vdo.UserId, tokenId)
	if err != nil {
		log.Printf("%+v", err)
		return nil
	}
	res.PlayUrl = vdo.PlayUrl
	res.CoverUrl = vdo.CoverUrl
	res.FavoriteCount = vdo.FavoriteCount
	res.CommentCount = int64(len(vdo.Comments))
	res.Title = vdo.Title
	//判断当前用户是否点赞
	f1 := false
	for i := 0; i < len(user.FavoriteList); i++ {
		if vdo.VideoId == user.FavoriteList[i] {
			f1 = true
			break
		}
	}
	res.IsFavorite = f1
	return res
}

func TimeStampToTime(stamp int64) time.Time {
	tm := time.Unix(stamp, 0)
	t := tm.Format("2006-01-02 15:04:05")
	timeLayout := "2006-01-02 15:04:05"                    //转化所需模板
	loc, _ := time.LoadLocation("Local")                   //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, t, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

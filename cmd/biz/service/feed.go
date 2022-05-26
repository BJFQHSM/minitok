package service

import (
	"context"
	"log"
	"time"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) Feed(req *biz.FeedRequest) ([]*biz.Video, int64, error) {
	latestTime := req.LatestTime
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}

	//从数据库查找数据
	t:=TimeStampToTime(latestTime)
	vdos, err := dal.QueryVideosByTime(t)
	if err!=nil{
		return []*biz.Video{},time.Now().Unix(),err
	}

	//获取下次的最新时间
	nextTime := time.Now().Unix()
	if len(vdos)>0{
		nextTime = vdos[len(vdos)-1].PublishDate.Unix()
	}

	videos := []*biz.Video{}
	for i:=0;i<len(vdos);i++{
		videos=append(videos,MongoVdoToBizVdo(vdos[i]))
	}

	return videos, nextTime, nil
}

//将video.go中的Video转化为biz.pb.go中的video类型
func MongoVdoToBizVdo(vdo *dal.Video) *biz.Video {
	res:=&biz.Video{}
	res.Id=vdo.VideoId
	//查询当前登录用户信息
	user, err:=dal.QueryUserByID(context.TODO(),vdo.UserId)
	if err!=nil{
		log.Println(err)
		return nil
	}
	//校验用户是否已关注
	f:=false
	for i:=0;i<len(user.Follows);i++{
		
	}
	u1:=biz.User{
		Id:            user.UserId,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      &f,
	}
	res.Author=&u1
	res.PlayUrl=vdo.PlayUrl
	res.CoverUrl=vdo.CoverUrl
	res.FavoriteCount=vdo.FavoriteCount
	res.CommentCount=int64(len(vdo.Comments))
	//判断当前用户是否点赞
	f1:=false
	for i:=0;i<len(user.FavoriteList);i++{
		if vdo.VideoId==user.FavoriteList[i]{
			f1=true
			break
		}
	}
	res.IsFavorite=&f1
	return res
}

func TimeStampToTime(stamp int64)time.Time{
	tm := time.Unix(stamp, 0)
	t := tm.Format("2006-01-02 15:04:05")
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, t, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

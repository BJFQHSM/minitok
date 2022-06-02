package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"log"
	"time"
)

type FeedService interface {
	DoService() *biz.FeedResponse
}


func NewFeedService(ctx context.Context, r *biz.FeedRequest) FeedService {
	return &feedServiceImpl{Req: r, Ctx: ctx, Resp: &biz.FeedResponse{}}
}

type feedServiceImpl struct {
	Req *biz.FeedRequest
	Resp *biz.FeedResponse
	Ctx context.Context
}

func (s *feedServiceImpl) DoService() *biz.FeedResponse {
	// mock
	//s.Resp = &biz.FeedResponse{
	//	VideoList: []*biz.Video{
	//		{Id: 2, Author: &biz.User{Id: 2}, CommentCount: 1, FavoriteCount: 2, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/6a87950831471cb0691b0f3dc2ae4428/628f35a7/video/tos/cn/tos-cn-ve-15c001-alinc2/cc660d533592437cb3377017d949ee13/?a=6383&ch=26&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=660&bt=660&cs=0&ds=3&ft=OyFYlOZZI0J.125TmVQbfzo57usylqG7Uag&mime_type=video_mp4&qs=0&rc=OWU7aGRlOjdkaDw1NmZnNEBpM3hweTU6ZmZvPDMzNGkzM0AxMTMuXzZiXzYxXzZjXzAvYSNkbWxzcjRvMGVgLS1kLS9zcw%3D%3D&l=202205261502500102020551523700C98B"},
	//		{Id: 3, Author: &biz.User{Id: 3}, CommentCount: 2, FavoriteCount: 3, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/8f4a23e58e5e011ff35f6f9a67d73dd8/628f3470/video/tos/cn/tos-cn-ve-15-alinc2/a71d679eb2d84b0cb5dbab33f68d94e1/?a=6383&ch=224&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1349&bt=1349&cs=0&ds=6&ft=5q_lc5mmnPD12Nuw3q.-UxHoFuYKc3wv25Na&mime_type=video_mp4&qs=0&rc=Omg8ZzkzN2Y7PDNoaWZnNUBpM3dkd2Q6Zmc3ODMzNGkzM0AzYjZjLWBgXi4xYWIzL2IzYSNuZi1fcjQwbHFgLS1kLTBzcw%3D%3D&l=202205261502460102081020853A00A5CB"},
	//		{Id: 3, Author: &biz.User{Id: 4}, CommentCount: 2, FavoriteCount: 3, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/8f4a23e58e5e011ff35f6f9a67d73dd8/628f3470/video/tos/cn/tos-cn-ve-15-alinc2/a71d679eb2d84b0cb5dbab33f68d94e1/?a=6383&ch=224&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1349&bt=1349&cs=0&ds=6&ft=5q_lc5mmnPD12Nuw3q.-UxHoFuYKc3wv25Na&mime_type=video_mp4&qs=0&rc=Omg8ZzkzN2Y7PDNoaWZnNUBpM3dkd2Q6Zmc3ODMzNGkzM0AzYjZjLWBgXi4xYWIzL2IzYSNuZi1fcjQwbHFgLS1kLTBzcw%3D%3D&l=202205261502460102081020853A00A5CB"},
	//	},
	//	StatusCode: 0,
	//}

	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams() ; err != nil {
			break
		}

		if err = s.feed() ; err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *feedServiceImpl) validateParams() error {
	return nil
}

func (s *feedServiceImpl) feed() error {
	latestTime := s.Req.LatestTime
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}

	//从数据库查找数据
	t:=TimeStampToTime(latestTime)
	vdos, err := dal.QueryVideosByTime(t)
	if err!=nil{
		return err
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
	res.Title = vdo.Title
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

func TimeStampToTime(stamp int64) time.Time{
	tm := time.Unix(stamp, 0)
	t := tm.Format("2006-01-02 15:04:05")
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, t, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

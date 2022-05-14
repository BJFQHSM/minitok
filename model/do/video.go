package do

type Video struct {
	Id int64 `bson:"id"`
	User User `bson:"user"`
	PlayUrl string `bson:"play_url"`
	CoverUrl string `bson:"cover_url"`
	FavoriteCount int64 `bson:"favorite_count"`
	CommentCount int64 `bson:"comment_count"`
	Comments []Comment `bson:"comments"`
	PublishTime string `bson:"publish_time"`
}

type Comment struct {
	Id int64 `bson:"id"`
	UserId int64 `bson:"user_id"`
	Content string `bson:"content"`
	CreateDate string `bson:"create_date"`
}


package do

type User struct {
	Id            int64   `bson:"id"`
	Name          string  `bson:"name"`
	FollowCount   int64   `bson:"follow_count"`
	FollowerCount int64   `bson:"follower_count"`
	FollowList    []int64 `bson:"follow_list"`
	FollowerList  []int64 `bson:"follower_list"`
	PublishList   []int64 `bson:"publish_list"`
	FavoriteList  []int64 `bson:"favorite_list"`
}

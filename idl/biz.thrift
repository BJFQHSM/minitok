// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the  License.
//
namespace go biz


struct User {
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:optional bool is_follow
}

struct Video {
    1:i64 id
    2:User author
    3:string play_url
    4:string cover_url
    5:i64 favorite_count
    6:i64 comment_count
    7:optional bool is_favorite
}

struct FeedRequest {
    1:i64 latest_time
}

struct FeedResponse {
    1:list<Video> video
    2:i64 next_time
    3:i32 status_code
    4:optional string status_msg
}

struct PublishActionRequest {
    1:i64 user_id
    2:string token
    3:list<byte> data
}

struct PublishActionResponse {
    1:i32 status_code
    2:optional string status_msg
}

struct PublishListRequest {
    1:i64 user_id
    2:string token
}

struct PublishListResponse {
    1:list<Video> video_list
    2:i32 status_code
    3:optional string status_msg
}

struct FavoriteActionRequest {
    1:i64 user_id
    2:string token
    3:i64 video_id
    4:i32 action_type
}

struct FavoriteActionResponse {
    1:i32 status_code
    2:optional string status_msg
}

struct FavoriteListRequest {
    1:i64 user_id
    2:string token
}

struct FavoriteListResponse {
    1:list<Video> video_list
    2:i32 status_code
    3:optional string status_msg
}

struct CommentActionRequest {
    1:i64 user_id
    2:string token
    3:i64 video_id
    4:i32 action_type
    5:optional string comment_text
    6:optional i64 comment_id
}

struct CommentActionResponse {
    1:i32 status_code
    2:optional string status_msg
}

struct Comment {
    1:i64 id
    2:User user
    3:string content
    4:string create_date
}

struct CommentListRequest {
    1:i64 user_id
    2:string token
    3:i64 video_id
}

struct CommentListResponse {
    1:list<Comment> comment_list
    2:i32 status_code
    3:optional string status_msg
}

struct RelationActionRequest {
    1:i64 user_id
    2:string token
    3:i64 to_user_id
    4:i32 action_type
}

struct RelationActionResponse {
    1:i32 status_code
    2:optional string status_msg
}

struct FollowListRequest {
    1:i64 user_id
    2:string token
}

struct FollowListResponse {
    1:list<User> user_list
    2:i32 status_code
    3:optional string status_msg
}

struct FollowerListRequest {
    1:i64 user_id
    2:string token
}

struct FollowerListResponse {
    1:list<User> user_list
    2:i32 status_code
    3:optional string status_msg
}

service BizService {
    FeedResponse Feed(1:FeedRequest req)
    PublishActionResponse PublishAction(1:PublishActionRequest req)
    PublishListResponse PublishList(1:PublishListRequest req)
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1:FavoriteListRequest req)
    CommentActionResponse CommentAction(1:CommentActionRequest req)
    CommentListResponse CommentList(1:CommentListRequest req)
    RelationActionResponse RelationAction(1:RelationActionRequest req)
    FollowListResponse FollowList(1:FollowListRequest req)
    FollowerListResponse FollowerList(1:FollowerListRequest req)
}
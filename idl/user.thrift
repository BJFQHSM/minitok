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

namespace go user

struct RegisterRequest {
    1:string username
    2:string password
}

struct RegisterResponse {
    1:i64 user_id
    2:string token
    3:i32 status_code
    4:optional string status_msg
}

struct LoginRequest {
    1:string username
    2:string password
}

struct LoginResponse {
    1:i64 user_id
    2:string token
    3:i32 status_code
    4:optional string status_msg
}

struct User {
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}

struct QueryInfoRequest {
    1:i64 user_id
    2:string token
}

struct QueryInfoResponse {
    1:User user
    2:i32 status_code
    3:optional string status_msg
}

service UserService {
    RegisterResponse Register(1:RegisterRequest req)
    LoginResponse Login(1:LoginRequest req)
    QueryInfoResponse QueryInfo(1:QueryInfoRequest req)
}
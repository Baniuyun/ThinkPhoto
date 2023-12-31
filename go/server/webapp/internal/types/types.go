// Code generated by goctl. DO NOT EDIT.
package types

type Token struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type RegisterRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserId int64 `json:"user_id"`
	Token  Token `json:"token"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId int64 `json:"user_id"`
	Token  Token `json:"token"`
}

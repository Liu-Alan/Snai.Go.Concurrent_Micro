// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
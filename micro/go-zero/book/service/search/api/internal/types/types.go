// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type SearchReply struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type SearchReq struct {
	Name string `form:"name"`
}

package model

type TagReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TagResp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

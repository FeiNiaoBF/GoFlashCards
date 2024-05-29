package model

import "time"

type CradReq struct {
	Id      int       `json:"id"`
	Front   string    `json:"front"`
	Back    string    `json:"back"`
	Know    bool      `json:"know"`
	Tags_id int       `json:"tags_id"`
	AddTime time.Time `json:"add_time"`
}


type CradResp struct {
	Id      int       `json:"id"`
	Front   string    `json:"front"`
	Back    string    `json:"back"`
	Know    bool      `json:"know"`
	Tags_id int       `json:"tags_id"`
	AddTime time.Time `json:"add_time"`
}

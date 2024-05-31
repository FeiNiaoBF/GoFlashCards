package model

type TagInput struct {
	ID   int64    `json:"id" query:"id" form:"id"`
	Name string `json:"name" query:"name" form:"name"`
}

type TagOutput struct {
	ID   int    `json:"id" query:"id" form:"id"`
	Name string `json:"name" query:"name" form:"name"`
}

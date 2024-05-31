package model

type CardInput struct {
	Front  string `json:"front" query:"front" form:"front"`
	Back   string `json:"back" query:"back" form:"back"`
	TagsID int32  `json:"tags_id" query:"tags_id" form:"tags_id"`
}

type CardOutput struct {
	Front  string `json:"front" query:"front" form:"front"`
	Back   string `json:"back" query:"back" form:"back"`
	TagsID int32  `json:"tags_id" query:"tags_id" form:"tags_id"`
	Know   bool   `json:"know" query:"know" form:"know"`
}

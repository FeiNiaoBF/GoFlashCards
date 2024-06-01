package model

type CardInput struct {
	Front  string `json:"front" query:"front" form:"front" validate:"required"`
	Back   string `json:"back" query:"back" form:"back" validate:"required"`
	TagsID int32  `json:"tags_id" query:"tags_id" form:"tags_id" validate:"required"`
}

type CardOutput struct {
	Front  string `json:"front" query:"front" form:"front" validate:"required"`
	Back   string `json:"back" query:"back" form:"back" validate:"required"`
	TagsID int32  `json:"tags_id" query:"tags_id" form:"tags_id" validate:"required"`
	Know   bool   `json:"know" query:"know" form:"know" validate:"required"`
}

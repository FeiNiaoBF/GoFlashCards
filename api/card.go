package api

type CardReq struct {
	Front  string `json:"front" query:"front" form:"front"`
	Back   string `json:"back" query:"back" form:"back"`
	TagsID int32  `json:"tags_id" query:"tags_id" form:"tags_id"`
}

type CardRes struct {
	Front  string `json:"front" query:"front" form:"front"`
	Back   string `json:"back" query:"back" form:"back"`
	TagsID int32  `json:"tags_id" query:"tags_id" form:"tags_id"`
	Konw   bool   `json:"know" query:"know" form:"know"`
}

func CreateCard() {
}

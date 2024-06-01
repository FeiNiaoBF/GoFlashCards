package model

const (
	All = iota
	Know
	Unknown
)

// filter handler
type filterType = int

type Filtype struct {

	Ty   filterType `json:"ty" query:"ty" form:"ty"`
	Name string     `json:"name" query:"name" form:"name"` // know unknow all
}

func NewFilt(name string, ty filterType) *Filtype {
	return &Filtype{
		Name: name,
		Ty:   ty,
	}
}

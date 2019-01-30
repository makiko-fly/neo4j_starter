package types

type ProductIn struct {
	Name         string `json:"name"`
	ImgActivated string `json:"img_activated"`
	ImgNormal    string `json:"img_normal"`
}

type Product ProductIn

type UpdateProductIn struct {
	Id      int64  `json:"id"`
	OldName string `json:"old_name"`
	ProductIn
}

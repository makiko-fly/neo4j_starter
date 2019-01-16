package types

type ProductIn struct {
	Name         string `json:"name"`
	ImgActivated string `json:"img_activated"`
	ImgNormal    string `json:"img_normal"`
}

type Product ProductIn

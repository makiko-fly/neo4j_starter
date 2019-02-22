package types

type ChainIn struct {
	Name string `json:"name"`
}

type GetChainIn struct {
	Name string `json:"name"`
}

type ListChainsIn struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type AddProductToChainIn struct {
	ProductId   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	ChainId     int64  `json:"chain_id"`
	ChainName   string `json:"chain_name"`
}

type RemProductFromChainIn struct {
	AddProductToChainIn
}

type GetProductsOfChainIn struct {
	ChainId   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
}

type UpdateChainIn struct {
	Id         int64  `json:"id"`
	OldName    string `json:"old_name"`
	Name       string `json:"name"`
	PreviewImg string `json:"preview_img'`
	CustomData string `json:"custom_data"`
}

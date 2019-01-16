package types

type StockIn struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Stock StockIn

package dto

type CreateItemRequestBody struct {
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
	Price    int64  `json:"price"`
}

type UpdateItemRequestBody struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
	Price    int64  `json:"price"`
}

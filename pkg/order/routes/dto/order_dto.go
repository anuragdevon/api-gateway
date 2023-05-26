package dto

type CreateOrderRequestBody struct {
	ItemId   int64 `json:"item_id"`
	Quantity int64 `json:"quantity"`
	UserId   int64 `json:"user_id"`
}

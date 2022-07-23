package model

const (
	ActionBuy    = "Buy"
	ActionSell   = "Sell"
	ActionCancel = "Cancel"
)

type OrderRequest struct {
	Quantity int
	Price    float32
}

type OrderResponse struct {
	OrderID string
}

package model

type TradeUsecase interface {
	Start()
	InsertBuyOrder(req *OrderRequest) (res *OrderResponse, err error)
	InsertSellOrder(req *OrderRequest) (res *OrderResponse, err error)
	GetOrderStatus(orderID string) (result *TradeResult)
	CancelOrder(orderID string) (res *OrderResponse, err error)
}

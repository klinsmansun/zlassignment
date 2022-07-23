package model

const (
	ReasonFinished  string = "order finished"
	ReasonCancelled string = "order cancelled"
)

type TradeResult struct {
	OrderID         string
	Action          string
	Price           float32
	TotalQuantity   int
	SucceedQuantity int
	Finished        bool
	Reason          string
}

type ResultUsecase interface {
	Set(id string, data *TradeResult)
	Get(id string) *TradeResult
	Cancel(id string) (*TradeResult, error)
}

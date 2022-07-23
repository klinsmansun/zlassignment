package grpc

import (
	context "context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/klinsmansun/zlassignment/model"
)

type grpcStruct struct {
	tradeUsecase model.TradeUsecase

	UnimplementedOrderServiceServer
}

func RegisterGRPCRoute(server *grpc.Server, tradeUsecase model.TradeUsecase) {
	r := &grpcStruct{
		tradeUsecase: tradeUsecase,
	}
	RegisterOrderServiceServer(server, r)
}

func (g *grpcStruct) Buy(ctx context.Context, req *TradeRequest) (resp *TradeResponse, err error) {
	resp = new(TradeResponse)

	res, e := g.tradeUsecase.InsertBuyOrder(&model.OrderRequest{
		Quantity: int(req.Quantity),
		Price:    req.Price,
	})

	if e != nil {
		resp.RespCode = RespCode_BUSY
	} else {
		resp.RequestID = res.OrderID
	}

	return
}

func (g *grpcStruct) Sell(ctx context.Context, req *TradeRequest) (resp *TradeResponse, err error) {
	resp = new(TradeResponse)

	res, e := g.tradeUsecase.InsertSellOrder(&model.OrderRequest{
		Quantity: int(req.Quantity),
		Price:    req.Price,
	})

	if e != nil {
		resp.RespCode = RespCode_BUSY
	} else {
		resp.RequestID = res.OrderID
	}

	return
}

func (g *grpcStruct) Cancel(ctx context.Context, req *CancelRequest) (resp *CancelResponse, err error) {
	resp = new(CancelResponse)

	_, e := g.tradeUsecase.CancelOrder(req.RequestID)

	if e != nil {
		resp.RespCode = RespCode_BUSY
	}

	return
}
func (g *grpcStruct) CheckOrderResult(ctx context.Context, req *QueryRequest) (resp *QueryResponse, err error) {
	resp = new(QueryResponse)

	if r := g.tradeUsecase.GetOrderStatus(req.RequestID); r != nil {
		resp.Finished = r.Finished
		resp.Reason = r.Reason
		resp.Action = r.Action
		resp.Price = fmt.Sprintf("%.2f", r.Price)
		resp.TotalQuantity = int32(r.TotalQuantity)
		resp.SucceedQuantity = int32(r.SucceedQuantity)
	} else {
		resp.RespCode = RespCode_ORDERNOTEXIST
	}

	return
}

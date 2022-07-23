package core

import (
	"errors"
	"time"

	"github.com/klinsmansun/zlassignment/model"
)

type coreUsecase struct {
	config        *model.Config
	logUsecase    model.LogUsecase
	idGenerator   model.IDGenerator
	resultUsecase model.ResultUsecase

	order     chan *orderItem      // queue request when trading, including buy/sell/cancel
	buyQueue  map[float32][]string // used to match trade
	sellQueue map[float32][]string // used to match trade
}

type orderItem struct {
	id     string
	action string
	price  float32
}

func CreateCoreUsecase(
	config *model.Config,
	logUsecase model.LogUsecase,
	idGenerator model.IDGenerator,
	resultUsecase model.ResultUsecase,
) model.TradeUsecase {
	c := &coreUsecase{
		config:        config,
		logUsecase:    logUsecase,
		idGenerator:   idGenerator,
		resultUsecase: resultUsecase,
		order:         make(chan *orderItem, config.Core.ChannelLength),
		buyQueue:      map[float32][]string{},
		sellQueue:     map[float32][]string{},
	}

	return c
}

func (c *coreUsecase) Start() {
	tradeTicker := time.NewTicker(c.config.Core.TradeInterval * time.Second)
	for {
		// make sure c.matching() gets higher priority
		select {
		case <-tradeTicker.C:
			c.matching()
		default:
			c.acceptQueue()
		}
	}
}

func (c *coreUsecase) InsertBuyOrder(req *model.OrderRequest) (res *model.OrderResponse, err error) {
	id := c.idGenerator.GenerateID()

	select {
	case c.order <- &orderItem{id: id, action: model.ActionBuy, price: req.Price}:
		c.resultUsecase.Set(id, &model.TradeResult{
			OrderID:       id,
			Action:        model.ActionBuy,
			Price:         req.Price,
			TotalQuantity: req.Quantity,
		})
		res = &model.OrderResponse{
			OrderID: id,
		}
	default:
		err = errors.New(model.ErrorSystemBusy)
	}

	return
}

func (c *coreUsecase) InsertSellOrder(req *model.OrderRequest) (res *model.OrderResponse, err error) {
	id := c.idGenerator.GenerateID()

	select {
	case c.order <- &orderItem{id: id, action: model.ActionSell, price: req.Price}:
		c.resultUsecase.Set(id, &model.TradeResult{
			OrderID:       id,
			Action:        model.ActionSell,
			Price:         req.Price,
			TotalQuantity: req.Quantity,
		})
		res = &model.OrderResponse{
			OrderID: id,
		}
	default:
		err = errors.New(model.ErrorSystemBusy)
	}

	return
}

func (c *coreUsecase) GetOrderStatus(orderID string) (result *model.TradeResult) {
	result = c.resultUsecase.Get(orderID)

	return
}

func (c *coreUsecase) CancelOrder(orderID string) (res *model.OrderResponse, err error) {
	select {
	case c.order <- &orderItem{id: orderID, action: model.ActionCancel}:

	default:
		err = errors.New(model.ErrorSystemBusy)
	}

	return
}

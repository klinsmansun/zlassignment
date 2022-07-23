package core

import (
	"time"

	"github.com/klinsmansun/zlassignment/model"
)

func (c *coreUsecase) acceptQueue() (orderAdded int) {
	if orderLen := len(c.order); orderLen > 0 {
		for i := 0; i < orderLen; i++ {
			order := <-c.order
			switch order.action {
			case model.ActionBuy:
				c.buyQueue[order.price] = append(c.buyQueue[order.price], order.id)
				orderAdded++
			case model.ActionSell:
				c.sellQueue[order.price] = append(c.sellQueue[order.price], order.id)
				orderAdded++
			case model.ActionCancel:
				c.resultUsecase.Cancel(order.id)
			}
		}
	} else {
		time.Sleep(time.Millisecond * 100)
	}

	return
}

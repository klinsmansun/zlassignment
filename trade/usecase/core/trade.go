package core

import (
	"fmt"
	"sync"

	"github.com/klinsmansun/zlassignment/model"
)

type removeOrder struct {
	price     float32
	buy, sell int
}

func (c *coreUsecase) tradeOrders(price float32, buy, sell []string, wg *sync.WaitGroup, tbr chan<- *removeOrder) {
	defer wg.Done()
	orderToRemove := &removeOrder{
		price: price,
	}
	totalQuantity := 0
	c.logUsecase.LogDebug(fmt.Sprintf("trade order with price: %.2f", price))

	buyIndex, sellIndex := 0, 0
	for buyIndex < len(buy) && sellIndex < len(sell) {
		// get order result(order info) and check if it is cancelled
		buyOrderResult := c.resultUsecase.Get(buy[buyIndex])
		if buyOrderResult.Finished {
			buyIndex++
			orderToRemove.buy++
			continue
		}
		sellOrderResult := c.resultUsecase.Get(sell[sellIndex])
		if sellOrderResult.Finished {
			sellIndex++
			orderToRemove.sell++
			continue
		}

		buyQuantity := buyOrderResult.TotalQuantity - buyOrderResult.SucceedQuantity
		sellQuantity := sellOrderResult.TotalQuantity - sellOrderResult.SucceedQuantity

		// we can modify value of buyOrderResult/sellOrderResult here directly
		// because by design, this goroutine is the only one to modify them
		switch {
		case buyQuantity < sellQuantity:
			buyOrderResult.SucceedQuantity += buyQuantity
			sellOrderResult.SucceedQuantity += buyQuantity
			buyOrderResult.Finished = true
			buyOrderResult.Reason = model.ReasonFinished
			buyIndex++
			orderToRemove.buy++
			totalQuantity += buyQuantity
		case buyQuantity > sellQuantity:
			buyOrderResult.SucceedQuantity += sellQuantity
			sellOrderResult.SucceedQuantity += sellQuantity
			sellOrderResult.Finished = true
			sellOrderResult.Reason = model.ReasonFinished
			sellIndex++
			orderToRemove.sell++
			totalQuantity += sellQuantity
		default: // finish 2 orders at the same time
			buyOrderResult.SucceedQuantity += sellQuantity
			sellOrderResult.SucceedQuantity += sellQuantity
			buyOrderResult.Finished = true
			buyOrderResult.Reason = model.ReasonFinished
			sellOrderResult.Finished = true
			sellOrderResult.Reason = model.ReasonFinished
			buyIndex++
			sellIndex++
			orderToRemove.buy++
			orderToRemove.sell++
			totalQuantity += buyQuantity
		}
	}

	c.logUsecase.LogDebug(fmt.Sprintf("trade order with price: %.2f completed, succeed quantity: %d", price, totalQuantity))

	tbr <- orderToRemove
}

func (c *coreUsecase) removeFinishedOrder(infoChan <-chan *removeOrder, wg *sync.WaitGroup) {
	defer wg.Done()
	var removeQueue []*removeOrder

	for data := range infoChan {
		removeQueue = append(removeQueue, data)
	}

	// when reaching here, infoChan is already closed, means all trades are made
	// we can modify buyQueue and sellQueue now
	for _, info := range removeQueue {
		c.buyQueue[info.price] = c.buyQueue[info.price][info.buy:]
		c.sellQueue[info.price] = c.sellQueue[info.price][info.sell:]
	}
}

// main function for matching orders
// each price will be handled by a go routine
func (c *coreUsecase) matching() {
	var wg, removeWG sync.WaitGroup

	// this channel is used to gather all finished
	// c.removeFinishedOrder will queue all information and remove finish order from queue after all trades are done
	toBeRemoved := make(chan *removeOrder, 100)

	removeWG.Add(1)
	go c.removeFinishedOrder(toBeRemoved, &removeWG)

	for price, buy := range c.buyQueue {
		if sell := c.sellQueue[price]; len(buy) != 0 && len(sell) != 0 {
			wg.Add(1)
			// each price will be handled by a go routine
			go c.tradeOrders(price, buy, sell, &wg, toBeRemoved)
		}
	}
	wg.Wait()
	// all trades are made, close channel
	// this will trigger c.removeFinishedOrder to actually remove finished orders
	close(toBeRemoved)

	// wait until order queues are all updated
	removeWG.Wait()
}

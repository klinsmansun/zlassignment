package mem

import (
	"errors"
	"sync"

	"github.com/klinsmansun/zlassignment/model"
)

type resultUsecase struct {
	lock sync.RWMutex
	data map[string]*model.TradeResult
}

func CreateResultUsecase() model.ResultUsecase {
	r := &resultUsecase{
		data: map[string]*model.TradeResult{},
	}

	return r
}

func (r *resultUsecase) Set(id string, data *model.TradeResult) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.data[id] = data
}

func (r *resultUsecase) Get(id string) *model.TradeResult {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.data[id]
}

func (r *resultUsecase) Cancel(id string) (result *model.TradeResult, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result = r.data[id]
	if result == nil {
		err = errors.New(model.ErrorOrderNotExist)
		return
	}

	if !result.Finished {
		result.Finished = true
		result.Reason = model.ReasonCancelled
	} else {
		result = nil
		err = errors.New(model.ErrorOrderFinished)
	}

	return
}

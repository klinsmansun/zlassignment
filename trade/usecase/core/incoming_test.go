package core

import (
	"testing"

	"github.com/klinsmansun/zlassignment/model"
)

func Test_coreUsecase_acceptQueue(t *testing.T) {
	type fields struct {
		config        *model.Config
		logUsecase    model.LogUsecase
		idGenerator   model.IDGenerator
		resultUsecase model.ResultUsecase
		order         chan *orderItem
		buyQueue      map[float32][]string
		sellQueue     map[float32][]string
	}
	tests := []struct {
		name           string
		fields         fields
		orders         []*orderItem
		wantOrderAdded int
	}{
		{
			name: "Test1",
			fields: fields{
				order:     make(chan *orderItem, 100),
				buyQueue:  map[float32][]string{},
				sellQueue: map[float32][]string{},
			},
			orders: []*orderItem{
				{action: model.ActionBuy},
				{action: model.ActionSell},
			},
			wantOrderAdded: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &coreUsecase{
				config:        tt.fields.config,
				logUsecase:    tt.fields.logUsecase,
				idGenerator:   tt.fields.idGenerator,
				resultUsecase: tt.fields.resultUsecase,
				order:         tt.fields.order,
				buyQueue:      tt.fields.buyQueue,
				sellQueue:     tt.fields.sellQueue,
			}
			for _, order := range tt.orders {
				c.order <- order
			}
			if gotOrderAdded := c.acceptQueue(); gotOrderAdded != tt.wantOrderAdded {
				t.Errorf("coreUsecase.acceptQueue() = %v, want %v", gotOrderAdded, tt.wantOrderAdded)
			}
		})
	}
}

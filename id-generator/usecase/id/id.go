package id

import (
	"strconv"
	"sync/atomic"

	"github.com/klinsmansun/zlassignment/model"
)

type uuidUsecase struct {
	id int64
}

func CreateIDGenerator() model.IDGenerator {
	g := &uuidUsecase{}

	return g
}

func (u *uuidUsecase) GenerateID() string {
	id := atomic.AddInt64(&u.id, 1)

	return strconv.FormatInt(id, 10)
}

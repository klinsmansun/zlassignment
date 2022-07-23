package model

type IDGenerator interface {
	GenerateID() string
}

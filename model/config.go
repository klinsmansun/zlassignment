package model

import "time"

type LogConfig struct {
	Level int `json:"level"`
}

type CoreConfig struct {
	ChannelLength int           `json:"channelLength"`
	TradeInterval time.Duration `json:"tradeInterval"`
}

type GRPCConfig struct {
	ListenIP   string `json:"listenIP"`
	ListenPort string `json:"listenPort"`
}

type Config struct {
	Version string     `json:"version"`
	Log     LogConfig  `json:"log"`
	Core    CoreConfig `json:"core"`
	GRPC    GRPCConfig `json:"grpc"`
}

type ConfigLoader interface {
	LoadConfig() *Config
}

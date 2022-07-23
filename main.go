package main

import (
	"net"

	"google.golang.org/grpc"

	_config "github.com/klinsmansun/zlassignment/config/repository/file"
	_id "github.com/klinsmansun/zlassignment/id-generator/usecase/id"
	_logger "github.com/klinsmansun/zlassignment/log/usecase/std"
	_request "github.com/klinsmansun/zlassignment/request/delivery/grpc"
	_result "github.com/klinsmansun/zlassignment/result/usecase/mem"
	_trade "github.com/klinsmansun/zlassignment/trade/usecase/core"
)

func main() {
	ConfigLoader := _config.CreateConfigLoader(".", "config", "yml")
	config := ConfigLoader.LoadConfig()
	logger := _logger.CreateLogger(config.Log.Level)
	idGenerator := _id.CreateIDGenerator()
	resultUsecase := _result.CreateResultUsecase()
	tradeUsecase := _trade.CreateCoreUsecase(config, logger, idGenerator, resultUsecase)

	// create grpc server first
	grpcServer := grpc.NewServer()

	_request.RegisterGRPCRoute(grpcServer, tradeUsecase)

	// start trade engine
	go tradeUsecase.Start()

	// start listen request
	lis, err := net.Listen("tcp", config.GRPC.ListenIP+":"+config.GRPC.ListenPort)
	if err != nil {
		logger.LogErr(err)
	}
	logger.LogDebug("server listening at", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logger.LogErr("failed to serve: %v", err)
	}
}

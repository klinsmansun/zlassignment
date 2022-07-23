# zlassignment
## Introduction
### this is sample implementation for zerologix assignment
&nbsp;
## Prerequisite
### the code structure is based on [clean architecture](https://github.com/bxcodec/go-clean-arch), which defines four domain layers: model, usecase, delivery and repository
&nbsp;
## Folder Definition
|Folder|Description|
|------|-----------|
|config|load configuration, currently implements loading config from file|
|id-generator|generate ID for each order(buy/sell), currently implements generating id from a unique number|
|log|generate log message, currently outputs log message to stdout|
|model|this folder defines all data structure and interfaces for other modules to communicate with each other|
|request|accepts order and query requests, currently implements grpc server|
|result|responsible for saving order results, currently implements in memory storage|
|trade|main business logic|
|main.go|main function, handle module initialization and dependency|
|config.yml|config file of this application, check it for comment|

&nbsp;
## GRPC Interface
### Tis application use grpc to accept incoming requests, we can use [BloomRPC](https://github.com/bloomrpc/bloomrpc) as grpc client to test
|GRPC Method|Description|
|-----------|-----------|
|Buy|Insert an new buy order at specific price|
|Sell|Insert an new sell order at specific price|
|Cancel|Cancel an existing order|
|CheckOrderResult|Check status/result of and existing order|

&nbsp;
## Build and Run
```bash
cd ${PROJECT_TOP_FOLDER} # contains main.go and confilg.yml
go run main.go
```
&nbsp;
## Key design points
- Trade buy/sell orders every 5 seconds(this interval is configurable)
- Matching orders has higher priority than receiving new orders
  - when system is busy, existing order will be handled first
  - when system is busy, new orders may be rejected
- The purpose of this design is to avoid too many race conditions
  - when trade is in process, internal queue will only be accessed by trade routine
  - when trade is not processing, new orders will be inserted to internal queue
- Main trade logic is in trade/usecase/core/trade.go, check in-line comment for details

&nbsp;
## TODO
- Unit test coverage
- Error handle

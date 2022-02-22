package main

import (
	_controller "exchange-rate/controller"
	_driver "exchange-rate/driver"
	_repository "exchange-rate/repository"
	_service "exchange-rate/service"

	"github.com/gin-gonic/gin"
)

var (
	client, _                                = _driver.CreateClient()
	repository _repository.MongoDbRepository = _repository.NewRepository(client)
	service    _service.Service              = _service.NewAccountService(repository)
	controller _controller.Controller        = _controller.NewAccountController(service)
)

func main() {
	handleRequests().Run(":5000")
}

func handleRequests() *gin.Engine {
	engine := gin.Default()
	engine.GET("/:rateTo/:rateFrom/:amount/convertRate", controller.ConvertExchangeRatesHandler)
	engine.GET("/getAll", controller.GetAllHandler)
	engine.POST("/pullRate", controller.PullExchangeRateHandler)
	return engine
}

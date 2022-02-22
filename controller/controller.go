package controller

import (
	_service "exchange-rate/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	ConvertExchangeRatesHandler(c *gin.Context)
	GetAllHandler(c *gin.Context)
	PullExchangeRateHandler(c *gin.Context)
}

type controller struct{}

var service _service.Service

func NewAccountController(s _service.Service) Controller {
	service = s
	return &controller{}
}

func (*controller) ConvertExchangeRatesHandler(c *gin.Context) {
	rateTo := c.Param("rateTo")
	rateFrom := c.Param("rateFrom")
	amount := c.Param("amount")

	convertedAmount, _ := strconv.ParseFloat(amount, 64)

	rate, err := service.ConvertExchangeRates(rateTo, rateFrom, convertedAmount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rate)
}

func (*controller) GetAllHandler(c *gin.Context) {
	result, err := service.GetAllRates()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (*controller) PullExchangeRateHandler(c *gin.Context) {
	response, err := service.PullExchangeRates()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR ": err.Error()})
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, response)
}

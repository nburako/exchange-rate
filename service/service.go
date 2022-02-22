package service

import (
	"encoding/json"
	"exchange-rate/config"
	"exchange-rate/model"
	_repository "exchange-rate/repository"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

type Service interface {
	ConvertExchangeRates(rateTo string, rateFrom string, amount float64) (float64, error)
	GetAllRates() ([]model.ExchangeRate, error)
	PullExchangeRates() ([]byte, error)
}

type service struct{}

var repository _repository.MongoDbRepository

func NewAccountService(repo _repository.MongoDbRepository) Service {
	repository = repo
	return &service{}
}

func (*service) ConvertExchangeRates(rateTo string, rateFrom string, amount float64) (float64, error) {
	result1, err := repository.GetById(rateTo)

	if err != nil {
		return 0, err
	}

	result2, err := repository.GetById(rateFrom)

	if err != nil {
		return 0, err
	}

	convertedRate := (result1.Value / result2.Value) * amount

	return convertedRate, nil
}

func (*service) GetAllRates() ([]model.ExchangeRate, error) {
	result, err := repository.GetAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) PullExchangeRates() ([]byte, error) {
	config, err := config.GetConfig("config/config")

	if err != nil {
		return nil, err
	}

	response, err := http.Get(config.ExternalApi.Endpoint)

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var responseJson = &model.Response{}
	json.Unmarshal([]byte(responseData), responseJson)

	var rateList []model.ExchangeRate
	v := reflect.ValueOf(responseJson.Rates)
	for i := 0; i < v.NumField(); i++ {
		rate := model.ExchangeRate{
			CreateTime:   time.Now(),
			CurrencyCode: v.Type().Field(i).Name,
			Value:        v.FieldByName(v.Type().Field(i).Name).Float(),
			SourceCode:   "ExternalApi"}

		rateList = append(rateList, rate)
	}

	mongoErr := repository.CreateMany(rateList)

	if mongoErr != nil {
		return nil, mongoErr
	}

	return responseData, nil
}

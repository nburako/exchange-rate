package repository

import (
	"context"
	"exchange-rate/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbRepository interface {
	CreateMany(list []model.ExchangeRate) error
	GetById(currencyCode string) (model.ExchangeRate, error)
	GetAll() ([]model.ExchangeRate, error)
}

type repo struct{ client *mongo.Client }

func NewRepository(client *mongo.Client) MongoDbRepository {
	return &repo{client: client}
}

func (db *repo) CreateMany(list []model.ExchangeRate) error {
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}

	collection := db.client.Database("dev").Collection("exchange-rate")

	_, err := collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}

	return nil
}

func (db *repo) GetById(currencyCode string) (model.ExchangeRate, error) {
	rate := model.ExchangeRate{}
	filter := bson.D{primitive.E{Key: "CurrencyCode", Value: currencyCode}}

	collection := db.client.Database("dev").Collection("exchange-rate")

	err := collection.FindOne(context.TODO(), filter).Decode(&rate)
	if err != nil {
		return rate, err
	}

	return rate, nil
}

func (db *repo) GetAll() ([]model.ExchangeRate, error) {
	filter := bson.D{{}}
	var rates []model.ExchangeRate

	collection := db.client.Database("dev").Collection("exchange-rate")

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return rates, findError
	}

	for cur.Next(context.TODO()) {
		var t model.ExchangeRate
		err := cur.Decode(&t)
		if err != nil {
			return rates, err
		}
		rates = append(rates, t)
	}

	cur.Close(context.TODO())
	if len(rates) == 0 {
		return rates, mongo.ErrNoDocuments
	}

	return rates, nil
}

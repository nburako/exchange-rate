package model

import (
	"time"
)

type ExchangeRate struct {
	CurrencyCode string    `bson:"CurrencyCode"`
	SourceCode   string    `bson:"SourceCode"`
	Value        float64   `bson:"Value"`
	CreateTime   time.Time `bson:"CrateTime"`
}

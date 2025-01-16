package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Coin struct {
	Symbol string
	Price  float64
	Time   time.Time
}

type CoinsDB struct {
	bun.BaseModel `bun:"table:currency.coins"`

	Id       int    `bun:"id"`
	CoinName string `bun:"coin_name"`
	IsActual string `bun:"is_actual"`
}

package models

import "time"

type Coin struct {
	Symbol string
	Price  float64
	Time   time.Time
}

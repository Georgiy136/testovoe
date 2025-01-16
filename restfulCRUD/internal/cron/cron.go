package cron

import "myapp/clients"

func NewCron(binanceApiClient clients.BinanceApiClient, secondsInterval int64) *Cron {
	return &Cron{
		binanceApiClient: binanceApiClient,
		secondsInterval:  secondsInterval,
	}
}

type Cron struct {
	binanceApiClient clients.BinanceApiClient
	secondsInterval  int64
}

func (c *Cron) Configure() {
}

func (c *Cron) Work() {
	for {

	}
}

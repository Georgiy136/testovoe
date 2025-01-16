package cron

import (
	"fmt"
	"log"
	"myapp/clients"
	"myapp/internal/usecase/repository"
	"myapp/pkg/cache"
	"strings"
	"sync"
	"time"
)

func NewCron(binanceApiClient clients.BinanceApiClient, repo repository.Project, cache cache.Cache, secondsInterval int) *Cron {
	return &Cron{
		binanceApiClient: binanceApiClient,
		secondsInterval:  secondsInterval,
		cache:            cache,
		repo:             repo,
	}
}

type Cron struct {
	binanceApiClient clients.BinanceApiClient
	secondsInterval  int
	cache            cache.Cache
	repo             repository.Project
}

func (c *Cron) Configure() {
	// получить основные койны из БД

	// установливаем в кэш

}

func (c *Cron) Work() {
	for {
		var (
			maxGoroutines = 100
			sem           = make(chan struct{}, maxGoroutines)
			errMsg        = strings.Builder{}
			wg            = &sync.WaitGroup{}
		)
		// получаем названия койнов из памяти
		for coin := range c.cache.GetListCoins() {
			wg.Add(1)
			// выполняем запрос http и получаем актуальную инфо-ию
			go func(coin string) {
				defer wg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()

				res, err := c.binanceApiClient.GetCoin(coin)
				if err != nil {
					errMsg.WriteString(fmt.Sprintf("[ERROR] Get coin %s error: %s\n", coin, err.Error()))
				}
				// запись инфо в БД

			}(coin)
		}

		wg.Wait()

		if errMsg.Len() > 0 {
			log.Println(errMsg.String())
		}
		time.Sleep(time.Second * time.Duration(c.secondsInterval))
	}
}

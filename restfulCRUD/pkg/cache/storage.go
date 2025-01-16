package cache

import (
	"sync"
)

type Cache struct {
	coins map[string]struct{}

	mu sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		coins: make(map[string]struct{}),
	}
}

func (c *Cache) GetListCoins() map[string]struct{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.coins
}

func (c *Cache) AddCoins(coins []string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, coin := range coins {
		c.coins[coin] = struct{}{}
	}
}

func (c *Cache) AddCoin(coin string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.coins[coin] = struct{}{}
}

func (c *Cache) DeleteCoin(coin string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	delete(c.coins, coin)
}

package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myapp/internal/models"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func NewBinanceApiClient() *BinanceApiClient {
	return &BinanceApiClient{}
}

type BinanceApiClient struct{}

const binanceAPI = "https://api.binance.com/api/v4"

func (b *BinanceApiClient) GetCoin(coin string) (*models.Coin, error) { // "BTCUSDT"
	api := "/ticker/price"

	baseURL, _ := url.Parse(binanceAPI + api)
	params := url.Values{}
	params.Add("symbol", coin)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка: статус ответа %s, ответ %s", resp.Status, string(body))
	}

	var result = struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf(" json.Unmarshal error: %w", err)
	}

	price, err := strconv.ParseFloat(result.Price, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat error: %w", err)
	}

	return &models.Coin{
		Symbol: result.Symbol,
		Price:  price,
		Time:   time.Now().UTC(),
	}, nil
}

package coincap

import (
	"fmt"
	"time"
)

type AssetsResponse struct {
	Data []struct {
		ID                string `json:"id"`
		Rank              string `json:"rank"`
		Symbol            string `json:"symbol"`
		Name              string `json:"name"`
		Supply            string `json:"supply"`
		MaxSupply         string `json:"maxSupply"`
		MarketCapUsd      string `json:"marketCapUsd"`
		VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
		PriceUsd          string `json:"priceUsd"`
		ChangePercent24Hr string `json:"changePercent24Hr"`
		Vwap24Hr          string `json:"vwap24Hr"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

func (a *AssetsResponse) Info() []string {
	var DataInfo []string
	for _, v := range a.Data {
		DataInfo = append(DataInfo, fmt.Sprintf("ID: %s Rank: %s Symbol: %s, Name: %s Supply: %s PriseUSD: %s", v.ID, v.Rank, v.Symbol, v.Name, v.Supply, v.PriceUsd))
	}
	return DataInfo
}

func PrintInfo(c *AssetsResponse) {
	for i, v := range c.Data {
		if v.Name == "Ethereum" {
			fmt.Println(c.Info()[i])
		}
	}
	time.Sleep(time.Second * 15)
}

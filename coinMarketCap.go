package main

import (
	"fmt"
	"encoding/json"
)

const CMC_URL = "https://api.coinmarketcap.com/v1/ticker/?limit=100"

type coinMarketCapAsset struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	PriceInUsd  string `json:"price_usd"`
	PriceInBtc  string `json:"price_btc"`
	LastUpdated string `json:"last_updated"`
}

type CoinMarketCapAssetService struct {
	restClient    RestClient
	configuration Configuration
}

func newCoinMarketCapAssetService(restClient RestClient, configuration Configuration) *CoinMarketCapAssetService {
	return &CoinMarketCapAssetService{restClient, configuration}
}

func (service CoinMarketCapAssetService) Retrieve() ([]Asset, error) {
	response, err := service.restClient.GetResponse(CMC_URL)
	if err != nil {
		return nil, err
	}
	var cmcResponse []coinMarketCapAsset
	err = json.Unmarshal(response.Body, &cmcResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body as json. %s, %v", response.Body, err)
	}
	return mapToAssets(cmcResponse), nil
}

func mapToAssets(cmcResponse []coinMarketCapAsset) []Asset {
	assets := make([]Asset, len(cmcResponse))
	for i, cmcAsset := range cmcResponse {
		assets[i] = Asset{
			cmcAsset.Id,
			cmcAsset.Name,
			cmcAsset.Symbol,
			cmcAsset.LastUpdated,
			cmcAsset.PriceInUsd,
			cmcAsset.PriceInBtc,
		}
	}
	return nil
}

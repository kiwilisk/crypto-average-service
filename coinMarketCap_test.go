package main

import (
	"testing"
	"time"
)

type mockedRestClient struct {
}

func (mockedRestClient) GetResponse(url string) (*HttpResponse, error) {
	body := []byte (`[
	{
		"id": "bitcoin",
		"name": "Bitcoin",
		"symbol": "BTC",
		"rank": "1",
		"price_usd": "573.137",
		"price_btc": "1.0",
		"24h_volume_usd": "72855700.0",
		"market_cap_usd": "9080883500.0",
		"available_supply": "15844176.0",
		"total_supply": "15844176.0",
		"percent_change_1h": "0.04",
		"percent_change_24h": "-0.3",
		"percent_change_7d": "-0.57",
		"last_updated": "1472762067"
	}
]`)
	response := &HttpResponse{
		Body:       body,
		StatusCode: 200,
		ReceivedAt: time.Now(),
		Size:       int64(len(body)),
	}
	return response, nil
}

func TestMapEndpointResponseToAssetSlice(t *testing.T) {
	config := Configuration{AssetEndpointURL: "mockUrl"}
	restClient := mockedRestClient{}
	service := CoinMarketCapAssetService{restClient, config}

	assets, err := service.Retrieve()

	if err != nil {
		t.Fatalf("execution failed with error %v", err)
	}
	expectedAsset := Asset{
		"bitcoin",
		"Bitcoin",
		"BTC",
		"1472762067",
		"573.137",
		"1.0"}
	if assets[0] != expectedAsset || len(assets) != 1 {
		t.Fatalf("Asset did not meet expectation %v", err)
	}
}

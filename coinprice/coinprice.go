package coinprice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const coinbaseAPI = "https://api.coinbase.com/v2/prices/%s-%s/spot"

type data struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type cryptoValue struct {
	Data data `json:"data"`
}

func Getcoinprice(coin string, currency string) cryptoValue {
	coinPrice := fmt.Sprintf(coinbaseAPI, coin, currency)
	resp, err := http.Get(coinPrice)
	if err != nil {
		log.Fatalln(err)
	}
	var cResp cryptoValue
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	return cResp
}

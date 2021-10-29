package main

import (
	"fmt"

	"github.com/DMEvanCT/CoinBaseAPI/v1/coinprice"
)

func main() {
	coin := coinprice.Getcoinprice("BTC", "USD")
	fmt.Printf("%s: %s\n", coin.Data.Currency, coin.Data.Amount)

}

package main

import (
	"fmt"

	"github.com/DMEvanCT/CoinBaseAPI/v1/coinprice"
)

func main() {
	bitcoin := coinprice.Getcoinprice("BTC", "USD")
	fmt.Printf("%s: %s\n", bitcoin.Data.Currency, bitcoin.Data.Amount)
	eth := coinprice.Getcoinprice("ETH", "USD")
	fmt.Printf("%s: %s\n", eth.Data.Currency, eth.Data.Amount)
	etheur := coinprice.Getcoinprice("ETH", "EUR")
	fmt.Printf("%s: %s\n", etheur.Data.Currency, etheur.Data.Amount)

}

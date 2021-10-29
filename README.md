# CoinBaseAPI
This is a coinbase Golang API written in Go

## Getcoinprice
EX: Get Bitcoin Price
```golang
	bitcoin := coinprice.Getcoinprice("BTC", "USD")
	fmt.Printf("%s: %s\n", bitcoin.Data.Currency, bitcoin.Data.Amount)
```
EX: Get Eherium Price in USD

```go
	eth := coinprice.Getcoinprice("ETH", "USD")
	fmt.Printf("%s: %s\n", eth.Data.Currency, eth.Data.Amount)
```
ETH in Euro
```go
	etheur := coinprice.Getcoinprice("ETH", "EUR")
	fmt.Printf("%s: %s\n", etheur.Data.Currency, etheur.Data.Amount)
```

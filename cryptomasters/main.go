package main

import (
	"fem/go/cryptomasters/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	print(rate.Price, err)
}

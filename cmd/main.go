package main

import (
	"aivia-test-assignment/internal/api"
	"aivia-test-assignment/internal/services"
	"github.com/aiviaio/go-binance/v2"
	"log"
)

func main() {
	client := binance.NewClient("", "")

	priceRetriever := api.NewBinanceSymbolPricePrinter(client, log.Default())
	pricePrinter := services.NewPricePrinter(log.Default())

	pricePrinter.Print(5, priceRetriever)
}

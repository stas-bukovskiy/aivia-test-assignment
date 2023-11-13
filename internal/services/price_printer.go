package services

import (
	"aivia-test-assignment/internal/api"
	"log"
)

type PricePrinter struct {
	logger *log.Logger
}

func NewPricePrinter(logger *log.Logger) *PricePrinter {
	logger.SetPrefix("PricePrinter: ")
	return &PricePrinter{logger: logger}
}

func (p *PricePrinter) Print(numberOfSymbols int, retriever api.SymbolPriceRetriever) {
	symbolPrice, err := retriever.Retrieve(numberOfSymbols)

	for symbol, price := range symbolPrice {
		p.logger.Printf("%s %s\n", symbol, price)
	}

	if err != nil {
		p.logger.Printf("error retrieving prices: %+v\n", err)
	}
}

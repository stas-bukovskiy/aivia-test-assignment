package services

import "aivia-test-assignment/internal/api"

type SymbolPricePrinter interface {
	Print(numberOfSymbols int, retriever api.SymbolPriceRetriever)
}

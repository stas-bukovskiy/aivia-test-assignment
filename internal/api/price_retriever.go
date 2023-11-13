package api

type SymbolPriceRetriever interface {
	Retrieve(numberOfSymbols int) (map[string]string, error)
}

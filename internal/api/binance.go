package api

import (
	"context"
	"github.com/aiviaio/go-binance/v2"
	"log"
	"sync"
)

type BinanceSymbolPriceRetriever struct {
	client *binance.Client
	logger *log.Logger
}

func NewBinanceSymbolPricePrinter(client *binance.Client, logger *log.Logger) *BinanceSymbolPriceRetriever {
	logger.SetPrefix("BinanceSymbolPriceRetriever: ")
	return &BinanceSymbolPriceRetriever{client: client, logger: logger}
}

func (r *BinanceSymbolPriceRetriever) Retrieve(numberOfSymbols int) (map[string]string, error) {
	// retrieve all symbols from the API
	exchangeInfo, err := r.client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		r.logger.Printf("error retrieving exchange info: %+v\n", err)
		return nil, err
	}

	resWg := sync.WaitGroup{}
	resChan := make(chan map[string]string)

	res := make(map[string]string)
	resWg.Add(1)
	go func() {
		defer resWg.Done()

		// aggregate results
		for symbolPrice := range resChan {
			for symbol, price := range symbolPrice {
				res[symbol] = price
			}
		}
	}()

	workerWg := sync.WaitGroup{}

	// retrieve prices for each symbol in goroutines
	for i := 0; i < numberOfSymbols; i++ {
		workerWg.Add(1)
		go func(symbol binance.Symbol, res chan map[string]string) {
			defer workerWg.Done()

			prices, err := r.client.NewListPricesService().Symbol(symbol.Symbol).Do(context.Background())
			symbolPrice := make(map[string]string)
			if err != nil {
				r.logger.Printf("error retrieving price for symbol %s: %+v\n", symbol.Symbol, err)
				symbolPrice[symbol.Symbol] = "error: " + err.Error()
			} else {
				symbolPrice[symbol.Symbol] = prices[0].Price
			}

			res <- symbolPrice
		}(exchangeInfo.Symbols[i], resChan)
	}

	// wait all goroutines to finish
	workerWg.Wait()
	close(resChan)
	resWg.Wait()

	return res, nil
}

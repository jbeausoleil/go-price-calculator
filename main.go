package main

import (
	"example.com/go-price-calculator/file_manager"
	"example.com/go-price-calculator/prices"
	"fmt"
)

// main is the entry point of the program.
// It reads price data from a file and calculates the tax-included prices for different tax rates.
// The tax-included prices are then written to separate JSON files for each tax rate.
func main() {
	taxRates := []float64{0, .07, 0.10, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := file_manager.New("prices.txt", fmt.Sprintf("results/result_%.0f.json", taxRate*100))
		//cmdm := cmd_manager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index, taxRate := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Printf("Finished writing file: result_%.0f.json\n", taxRate*100)
		}
	}
}

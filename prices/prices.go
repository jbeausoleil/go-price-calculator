package prices

import (
	"example.com/go-price-calculator/conversions"
	"example.com/go-price-calculator/io_manager"
	"fmt"
)

// TaxIncludedPriceJob represents a job for calculating tax-included prices.
// It contains the tax rate, input prices, and a map to store the calculated tax-included prices.
type TaxIncludedPriceJob struct {
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices map[string]string    `json:"tax_included_prices"`
	IOManager         io_manager.IOManager `json:"-"`
}

// ReadData is a method of the TaxIncludedPriceJob type that reads pricing data from a file.
// It opens the file "prices.txt" and reads its content line by line.
// It converts each line to a floating-point number and stores it in the InputPrices field of the TaxIncludedPriceJob instance.
// If any error occurs during file reading or conversion, an error message is displayed, and the method returns without updating the InputPrices field.
func (job *TaxIncludedPriceJob) ReadData() error {

	lines, err := job.IOManager.ReadLines()

	prices, err := conversions.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.ReadData()

	if err != nil {
		//return err
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.1f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice) // required to convert float to string
	}

	job.TaxIncludedPrices = result
	err = job.IOManager.WriteResult(job)

	if err != nil {
		//return err
		errorChan <- err
		return
	}

	doneChan <- true
}

// NewTaxIncludedPriceJob is a function that creates a new instance of TaxIncludedPriceJob.
// It takes an instance of IOManager and a tax rate as parameters and returns a pointer to TaxIncludedPriceJob.
// The IOManager parameter is used for reading and writing data.
// The InputPrices field of the TaxIncludedPriceJob instance is initialized with a slice containing the values 10, 20, and 30.
// The TaxRate field of the TaxIncludedPriceJob instance is set to the provided tax rate.
func NewTaxIncludedPriceJob(iom io_manager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

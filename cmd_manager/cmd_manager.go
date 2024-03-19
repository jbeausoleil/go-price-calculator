package cmd_manager

import "fmt"

// CmdManager represents a command line manager.
type CmdManager struct {
}

// ReadLines reads a series of prices from the user until the user enters "0".
// It prompts the user to enter a price and stores each entered price in a slice.
// Once the user enters "0", the function returns the slice of prices.
// If there is an error while reading the prices, it returns an error.
func (fm CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices.  Confirm every price with ENTER.")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

// WriteResult is a method of CmdManager type that prints the provided data
// to the console.
// The method takes an input argument of type `interface{}` representing
// any data type. The data is printed using the `fmt.Println()` function.
// After printing the data, the method returns `nil` indicating successful execution.
// In case of any error, this method does not return any error.
// Please note that this method does not modify the `CmdManager` object or
// its state in any way.
func (fm CmdManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

// New creates a new instance of CmdManager.
func New() CmdManager {
	return CmdManager{}
}

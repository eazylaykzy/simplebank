package util

import "golang.org/x/exp/slices"

// Constants for all supported currencies
var currencies = []string{"USD", "EUR", "CAD"}

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	return slices.Contains(currencies, currency)
}

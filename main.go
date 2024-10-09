package main

import "fmt"

// exchangeToUSD menghitung total USD setelah menukar sejumlah IDR dan SGD ke USD
func exchangeToUSD(amountIDR float64, exchangeRateUSD_IDR float64, amountSGD float64, exchangeRateUSD_SGD float64) float64 {
	// Konversi IDR ke USD
	usdFromIDR := amountIDR / exchangeRateUSD_IDR
	// Konversi SGD ke USD
	usdFromSGD := amountSGD / exchangeRateUSD_SGD
	// Total USD
	totalUSD := usdFromIDR + usdFromSGD

	return totalUSD
}

func main() {
	// Contoh perhitungan
	amountIDR := 120000.0
	exchangeRateUSD_IDR := 15000.0
	amountSGD := 29.0
	exchangeRateUSD_SGD := 1.5

	totalUSD := exchangeToUSD(amountIDR, exchangeRateUSD_IDR, amountSGD, exchangeRateUSD_SGD)
	fmt.Printf("Total USD: %.2f\n", totalUSD) // Output: 27.33 USD
}

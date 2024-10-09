package main

import (
	"fmt"
	"math"
)

// calcPriceOfExcessBaggage menghitung harga bagasi berlebih
func calcPriceOfExcessBaggage(userBaggage int, freeQuotaBaggage int, domesticPrice int, internationalPrice int, flightType string) int {
	// Hitung berat bagasi berlebih
	excessBaggage := userBaggage - freeQuotaBaggage
	if excessBaggage <= 0 {
		return 0
	}

	// Pilih harga per kg sesuai jenis penerbangan
	if flightType == "domestic" {
		return excessBaggage * domesticPrice
	} else if flightType == "international" {
		return excessBaggage * internationalPrice
	}
	
	return 0
}

// calcFlightTime menghitung waktu penerbangan antara fromCountry dan destinationCountry berdasarkan worldMap
func calcFlightTime(worldMap []string, fromCountry string, destinationCountry string) int {
	var fromIndex, toIndex int = -1, -1

	// Cari indeks dari fromCountry dan destinationCountry di worldMap
	for i, country := range worldMap {
		if country == fromCountry {
			fromIndex = i
		}
		if country == destinationCountry {
			toIndex = i
		}
	}

	// Jika salah satu negara tidak ditemukan, kembalikan 0
	if fromIndex == -1 || toIndex == -1 {
		return 0
	}

	// Hitung perbedaan indeks, ini akan menjadi waktu dalam jam
	return int(math.Abs(float64(toIndex - fromIndex)))
}

func main() {
	// Contoh perhitungan harga bagasi berlebih
	fmt.Println(calcPriceOfExcessBaggage(50, 40, 110000, 160000, "domestic"))      // Output: 1100000
	fmt.Println(calcPriceOfExcessBaggage(50, 40, 110000, 160000, "international")) // Output: 1600000
	fmt.Println(calcPriceOfExcessBaggage(40, 50, 110000, 160000, "international")) // Output: 0

	// Peta dunia disederhanakan
	worldMap1 := []string{"PH", "", "ID", "SG", "MY", "VN", "KH"}

	// Contoh perhitungan waktu penerbangan
	fmt.Println(calcFlightTime(worldMap1, "ID", "KH")) // Output: 4
	fmt.Println(calcFlightTime(worldMap1, "SG", "PH")) // Output: 3
	fmt.Println(calcFlightTime(worldMap1, "KH", "PH")) // Output: 6
}

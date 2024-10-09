package main

import "fmt"

// calcPriceOfExcessBaggage menghitung harga bagasi berlebih berdasarkan berat, kuota, dan jenis penerbangan
func calcPriceOfExcessBaggage(userBaggage int, freeQuotaBaggage int, domesticPrice int, internationalPrice int, flightType string) int {
    // Hitung berat bagasi berlebih
    excessBaggage := userBaggage - freeQuotaBaggage
    if excessBaggage <= 0 {
        return 0 // Tidak ada biaya jika tidak ada bagasi berlebih
    }

    // Tentukan harga berdasarkan jenis penerbangan
    var pricePerKG int
    if flightType == "domestik" {
        pricePerKG = domesticPrice
    } else if flightType == "internasional" {
        pricePerKG = internationalPrice
    } else {
        return 0 // Jenis penerbangan tidak valid
    }

    // Hitung total biaya
    totalPrice := excessBaggage * pricePerKG
    return totalPrice
}

func main() {
    // Contoh penggunaan
    beratBagasi := 50
    kuotaBebasBiaya := 40
    hargaDomestik := 110000
    hargaInternasional := 160000
    jenisPenerbangan := "domestik"

    totalHarga := calcPriceOfExcessBaggage(beratBagasi, kuotaBebasBiaya, hargaDomestik, hargaInternasional, jenisPenerbangan)
    fmt.Printf("Total biaya bagasi berlebih: Rp %d\n", totalHarga) // Output: Rp 1100000
}

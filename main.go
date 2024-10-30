package main

import "fmt"

func main() {
	// Deklarasi Konstanta
	const satuan = "m"
	const satuan_luas = "m2"

	// Deklarasi Variabel
	var panjang, lebar int
	panjang = 10
	lebar = 5

	fmt.Println("Hitung Luas dan Keliling Persegi Panjang")
	fmt.Println("Panjang : ", panjang, satuan)
	fmt.Println("Lebar   : ", lebar, satuan)

	fmt.Println("Luas     = ", hitung_luas(panjang, lebar), satuan_luas)
	fmt.Println("Keliling = ", hitung_keliling(panjang, lebar), satuan)
}

// Function Hitung Luas
func hitung_luas(panjang, lebar int) int {
	return panjang * lebar
}

// Function Hitung Keliling
func hitung_keliling(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

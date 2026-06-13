package main

import "fmt"

func Quiz() {

	//cetak index yang genap, index ganjil jangan di cetak

	title := "Golang the best language"

	for index, letter := range title {
		// Kondisi: jika index habis dibagi 2 (index % 2 == 0) maka index adalah genap
		// % modulo (sisa bagi)
		if index%2 == 0 {
			fmt.Println("index :", index, "letter :", string(letter))
		}
	}

}

func Quiz2() {

	fmt.Println("=== Penjelasan Cara Cek Index Genap ===")
	fmt.Println("Gunakan operator modulo (%) untuk cek sisa bagi")
	fmt.Println("Jika index % 2 == 0 → index adalah GENAP")
	fmt.Println("Jika index % 2 == 1 → index adalah GANJIL")
	fmt.Println()

	Quiz()

	fmt.Println()
	fmt.Println("=== Contoh Lain: Dengan Slice ===")
	buah := []string{"Apel", "Jeruk", "Mangga", "Pisang", "Anggur"}
	for idx, val := range buah {
		if idx%2 == 0 {
			fmt.Printf("Index %d: %s (GENAP)\n", idx, val)
		}
	}

	fmt.Println()
	fmt.Println("=== Contoh Lain: Loop dengan For Biasa ===")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Angka %d adalah genap\n", i)
		}
	}
}

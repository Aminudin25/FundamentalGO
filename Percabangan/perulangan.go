package main

import (
	"fmt"
)

func PerulanganFor() {

	// i = 1
	// pengecekan if i <= 100 = TRUE
	// Cetak println
	// i++ = 1 + 1 = 2
	// i = 2
	// sampai i == 101

	for i := 1; i <= 100; i++ {
		fmt.Println("Saya sedang belajar Go:", i)
	}
}

func PerulanganWhile() {
	// sebenernya di go itu tidak ada perulangan While
	// tapi kita bisa membuat sama dengan While pada umumnya

	i := 1
	for i <= 100 {
		fmt.Println("Saya sedang belajar Go:", i)
		i++
	}

	// kenapa harus menggunakan i++ ??
	// kalo tidak ditambahkan maka i akan terus menjadi 1
	// dan kondisinya tidak akan berhenti
}

func PerulanganArray() {
	// sama dengan for each
	// sama dengan while, di GO tidak ada func for each
	// maka masih menggunakan for

	title := "Golang the best language"

	for index, letter := range title {
		fmt.Println("index :", index, "letter :", string(letter))
	}

	//perulangan ini yang sering digunakan pada saat kita mengerjakan project real
}

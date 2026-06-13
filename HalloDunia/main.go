package main

import (
	"fmt"
	"halodunia/calculator"
)

// go mod init <namaProject> = menginisiasi projek dan nanti akan dibuatkan file go.mod

// go build untuk membuat build dan akan ada file executable, dan ketika kita ingin menjalankan file tersebut kita cukup tuliskan nama executable nya

// Package fungsinya untuk mengelompokan
// sebagai contoh, kita membuat file baru dengan entity.go dan didalamnya terdapat function
// tetapi package nya masih sama, sehingga function tersebut masih kelompok yang sama
// sehingga kita juga bisa menggunakan function tersebut

// func main() {
// 	fmt.Println("Hallo Dunia, kita sedang belajar Pemogramman GO")

// 	// menggunakan function yang ada di file entity.go
// 	kalimat := TestAja()

// 	fmt.Println(kalimat)

// 	// sekarang seharusnya terdapat 2 paragraf yang dimunculkan ketika kita menjalankan programnya

// 	// akan tetapi ketika kita menjalankanya menggunakan perintah : go run main.go , hasilnya tidak akan muncul, karena function TestAja berada di file entity.go, sedangkan kita memanggil hanya main.go
// 	// jadi kita harus memanggil juga file entity.go, sehingga perintahnya seperti berikut : go run main.go entity.go

// 	// kenapa kita menamai package dengan nama "main"
// 	// itu atura GO yang wajib diikuti
// 	// terdapat 2 type penamaan package
// 	// 1.Executable (main)
// 	// 2.Library (selain main)

// 	// kita sudah berhasil membuat package Executable dengan 2 file
// 	// sekarang kita coba contohkan untuk package Library

// 	// ============== Lanjut di Bawah =============
// 	// ===== file entity kita disable =============
// }

// ======= file executable utk package Library ====
func main() {
	// buat file baru dengan package calculator
	fmt.Println("Hallo Dunia")

	// panggil function yang ada di package calculator
	add := calculator.Add(8, 9)
	fmt.Println(add)

	multiplication := calculator.Multiplication(5, 5)
	fmt.Println(multiplication)

	// hasilnya akan print :
	// Hallo Dunia
	// 17
}

// func main sebuah fungsion khusus, method khusus yang dipake untuk executable

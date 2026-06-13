package main

import "fmt"

func PercabanganSwitchCase() {
	number := 2

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("NULL")
	}

	// kasus seperti ini bisa saja menggunakan IF ELSE
	// akan tetapi lebih baik menggunakan switch case ketika mempunyai kondisi yang lebih dari 2

}

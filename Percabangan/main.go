package main

import "fmt"

// if atau percabangan itu sebuah kondisi yang harus dipenuhi
// contoh percabangan Umur

func main() {
	ageOld := 11

	if ageOld > 10 {
		fmt.Println("Kamu boleh bermain Game")
	} else {
		fmt.Println("Kamu belum boleh bermain Game")
	}

	// contoh lagi mengenai pengkondisian else if
	score := 80
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Print(grade)

	// contoh lagi if bersarang
	// rekomendasi jangan sampe lebih dari 2 if
	// contoh

	// Contoh: cek apakah seseorang bisa mendaftar untuk kursus mengemudi
	age := 17
	hasParentPermission := true

	if age >= 16 {
		fmt.Println("Usia memenuhi syarat")
		if hasParentPermission {
			fmt.Println("✓ Sudah dapat izin orang tua")
			fmt.Println("✓ Anda bisa mendaftar kursus mengemudi!")
		} else {
			fmt.Println("✗ Belum dapat izin orang tua")
			fmt.Println("Tidak bisa mendaftar")
		}
	} else {
		fmt.Println("Usia masih terlalu muda, minimal 16 tahun")
	}
}

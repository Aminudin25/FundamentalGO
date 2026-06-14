package main

import "fmt"

func Slice() {
	fmt.Println("=== PERBEDAAN ARRAY vs SLICE ===")

	// ==================== ARRAY ====================
	fmt.Println("📌 ARRAY - Ukuran TETAP/FIXED")
	fmt.Println("-------------------------------------------")

	// Array dengan ukuran 3 (HARUS 3, tidak bisa lebih atau kurang)
	var arrayBuah [3]string
	arrayBuah[0] = "Apel"
	arrayBuah[1] = "Jeruk"
	arrayBuah[2] = "Mangga"

	fmt.Println("Array buah:", arrayBuah)
	fmt.Println("Panjang array:", len(arrayBuah))
	// arrayBuah[3] = "Pisang" // ERROR! Index out of range, karena array hanya 3 elemen
	// arrayBuah = append(arrayBuah, "Pisang") // ERROR! Array tidak bisa pakai append

	fmt.Println()

	// ==================== SLICE ====================
	fmt.Println("📌 SLICE - Ukuran DINAMIS (bisa berubah)")
	fmt.Println("-------------------------------------------")

	// Slice tanpa ukuran (DINAMIS)
	var sliceLanguage []string
	fmt.Println("Slice awal (kosong):", sliceLanguage)
	fmt.Println("Panjang slice:", len(sliceLanguage))

	// Tambah value dengan append (bisa terus bertambah)
	sliceLanguage = append(sliceLanguage, "Go")
	sliceLanguage = append(sliceLanguage, "Python")
	sliceLanguage = append(sliceLanguage, "JavaScript")
	sliceLanguage = append(sliceLanguage, "Java") // Bisa ditambah terus!

	fmt.Println("Slice setelah append:", sliceLanguage)
	fmt.Println("Panjang slice:", len(sliceLanguage))

	fmt.Println()

	// ==================== PERBANDINGAN LANGSUNG ====================
	fmt.Println("📌 PERBANDINGAN LANGSUNG")
	fmt.Println("-------------------------------------------")

	arrayAngka := [4]int{10, 20, 30, 40} // Array dengan size 4
	sliceAngka := []int{10, 20, 30, 40}  // Slice (tanpa size)

	fmt.Printf("Array: %T → %v (size tetap 4)\n", arrayAngka, arrayAngka)
	fmt.Printf("Slice: %T → %v (size bisa berubah)\n", sliceAngka, sliceAngka)

	// Tambah ke slice
	sliceAngka = append(sliceAngka, 50)
	sliceAngka = append(sliceAngka, 60)
	fmt.Println("Slice setelah append:", sliceAngka)

	fmt.Println()

	// ==================== SLICE INISIALISASI LANGSUNG ====================
	fmt.Println("📌 SLICE dengan Inisialisasi Langsung")
	fmt.Println("-------------------------------------------")

	console := []string{"1", "2"}
	fmt.Println("Slice console:", console)
	fmt.Println("Panjang:", len(console))

	// Bisa tambah
	console = append(console, "3")
	console = append(console, "4", "5") // Bisa tambah multiple
	fmt.Println("Slice console setelah append:", console)

}

func Sintaks() {
	fmt.Println("=== PENJELASAN PERBEDAAN SINTAKS ===")
	fmt.Println("Perbedaan Array dan Slice HANYA pada ada/tidaknya ANGKA SIZE!")

	fmt.Println("✅ ARRAY - Ada angka SIZE di bracket:")
	fmt.Println("   [4]int{10, 20, 30, 40}     ← [4] = SIZE TETAP 4 elemen = ARRAY")
	fmt.Println("   [5]string{...}             ← [5] = SIZE TETAP 5 elemen = ARRAY")
	fmt.Println("   [3]bool{...}               ← [3] = SIZE TETAP 3 elemen = ARRAY")

	fmt.Println()

	fmt.Println("✅ SLICE - TIDAK ada angka SIZE di bracket:")
	fmt.Println("   []int{10, 20, 30, 40}      ← [] = TANPA SIZE = SLICE (dinamis)")
	fmt.Println("   []string{...}              ← [] = TANPA SIZE = SLICE (dinamis)")
	fmt.Println("   []bool{...}                ← [] = TANPA SIZE = SLICE (dinamis)")

	fmt.Println()
	fmt.Println("-------------------------------------------")
	fmt.Println("KESIMPULAN:")
	fmt.Println("• [ANGKA]Type  → ARRAY   (size tetap, panjang fixed)")
	fmt.Println("• []Type       → SLICE   (size dinamis, panjang bisa berubah)")
	fmt.Println("-------------------------------------------")

	fmt.Println()
	fmt.Println("CONTOH PRAKTIS:")

	// Jika dihapus angka 4-nya
	fmt.Println("\nBefore (ARRAY):")
	arraySebelum := [4]int{10, 20, 30, 40}
	fmt.Printf("  %T : %v\n", arraySebelum, arraySebelum)
	fmt.Println("  Size tetap 4, tidak bisa dirubah!")

	fmt.Println("\nAfter (SLICE) - Hanya dihapus angka [4]:")
	sliceSetelah := []int{10, 20, 30, 40}
	fmt.Printf("  %T : %v\n", sliceSetelah, sliceSetelah)
	fmt.Println("  Terlihat sama, tapi bisa pakai append:")
	sliceSetelah = append(sliceSetelah, 50, 60, 70)
	fmt.Printf("  Setelah append: %v (jadi 7 elemen!)\n", sliceSetelah)
}

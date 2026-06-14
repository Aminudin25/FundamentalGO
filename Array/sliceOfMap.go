package main

import "fmt"

func SliceOfMap() {
	fmt.Println("=== SLICE OF MAP ===")
	fmt.Println("Slice of Map adalah slice yang berisi MAP")
	fmt.Println("Sintaks: []map[keyType]valueType")

	// ==================== CONTOH 1: Slice of Map Kosong ====================
	fmt.Println("📌 CONTOH 1: Deklarasi Slice of Map Kosong")
	fmt.Println("-------------------------------------------")

	// Deklarasi slice of map (kosong)
	var students []map[string]string

	fmt.Println("Slice of map kosong:", students)
	fmt.Println("Panjang:", len(students))
	fmt.Println()

	// ==================== CONTOH 2: Menambah Map ke Slice ====================
	fmt.Println("📌 CONTOH 2: Menambah Map ke Slice dengan append")
	fmt.Println("-------------------------------------------")

	// Map pertama (siswa 1)
	siswa1 := map[string]string{
		"nama":    "Budi",
		"umur":    "17",
		"kelas":   "XI",
		"sekolah": "SMA Maju",
	}

	// Map kedua (siswa 2)
	siswa2 := map[string]string{
		"nama":    "Andi",
		"umur":    "16",
		"kelas":   "X",
		"sekolah": "SMA Maju",
	}

	// Map ketiga (siswa 3)
	siswa3 := map[string]string{
		"nama":    "Siti",
		"umur":    "17",
		"kelas":   "XI",
		"sekolah": "SMA Maju",
	}

	// Tambahkan ke slice dengan append
	students = append(students, siswa1, siswa2, siswa3)

	fmt.Println("Slice of map setelah append:")
	for i, siswa := range students {
		fmt.Printf("  Siswa %d: %v\n", i+1, siswa)
	}
	fmt.Println()

	// ==================== CONTOH 3: Akses Nilai dari Slice of Map ====================
	fmt.Println("📌 CONTOH 3: Mengakses Nilai dari Slice of Map")
	fmt.Println("-------------------------------------------")

	// Akses map pertama (index 0)
	fmt.Println("Data Siswa Pertama:")
	fmt.Println("  Nama:", students[0]["nama"])
	fmt.Println("  Umur:", students[0]["umur"])
	fmt.Println("  Kelas:", students[0]["kelas"])
	fmt.Println()

	// Akses map kedua (index 1)
	fmt.Println("Data Siswa Kedua:")
	fmt.Println("  Nama:", students[1]["nama"])
	fmt.Println("  Kelas:", students[1]["kelas"])
	fmt.Println()

	// ==================== CONTOH 4: Inisialisasi Langsung ====================
	fmt.Println("📌 CONTOH 4: Inisialisasi Slice of Map Langsung")
	fmt.Println("-------------------------------------------")

	produk := []map[string]interface{}{
		{
			"id":    1,
			"nama":  "Laptop",
			"harga": 5000000,
			"stok":  10,
		},
		{
			"id":    2,
			"nama":  "Mouse",
			"harga": 150000,
			"stok":  50,
		},
		{
			"id":    3,
			"nama":  "Keyboard",
			"harga": 300000,
			"stok":  30,
		},
	}

	fmt.Println("Daftar Produk:")
	for _, p := range produk {
		fmt.Printf("  [%d] %v - Rp%v (Stok: %v)\n", p["id"], p["nama"], p["harga"], p["stok"])
	}
	fmt.Println()

	// ==================== CONTOH 5: Loop dan Modifikasi ====================
	fmt.Println("📌 CONTOH 5: Loop dan Modifikasi Slice of Map")
	fmt.Println("-------------------------------------------")

	// Tambah siswa baru
	siswaBaru := map[string]string{
		"nama":    "Rini",
		"umur":    "16",
		"kelas":   "X",
		"sekolah": "SMA Maju",
	}
	students = append(students, siswaBaru)

	fmt.Println("Daftar semua siswa:")
	for idx, siswa := range students {
		fmt.Printf("  %d. %s (Kelas %s, Umur %s tahun)\n", idx+1, siswa["nama"], siswa["kelas"], siswa["umur"])
	}
	fmt.Println()

	// ==================== CONTOH 6: Cari Data Tertentu ====================
	fmt.Println("📌 CONTOH 6: Mencari Data dalam Slice of Map")
	fmt.Println("-------------------------------------------")

	cariNama := "Andi"
	fmt.Printf("Mencari siswa dengan nama '%s':\n", cariNama)

	for idx, siswa := range students {
		if siswa["nama"] == cariNama {
			fmt.Printf("  Ditemukan di index %d: %v\n", idx, siswa)
		}
	}
	fmt.Println()

	// ==================== CONTOH 7: Praktis - Toko Online ====================
	fmt.Println("📌 CONTOH 7: Praktis - Sistem Toko Online")
	fmt.Println("-------------------------------------------")

	keranjang := []map[string]interface{}{
		{
			"produk":   "Keyboard Mechanical",
			"harga":    500000,
			"jumlah":   1,
			"kategori": "Elektronik",
		},
		{
			"produk":   "Mouse Gaming",
			"harga":    250000,
			"jumlah":   2,
			"kategori": "Elektronik",
		},
		{
			"produk":   "Monitor 24 inch",
			"harga":    1500000,
			"jumlah":   1,
			"kategori": "Display",
		},
	}

	totalHarga := 0
	fmt.Println("📦 ISI KERANJANG:")
	for i, item := range keranjang {
		jumlah := item["jumlah"].(int)
		harga := item["harga"].(int)
		subtotal := jumlah * harga

		fmt.Printf("  %d. %s x%d = Rp%v\n", i+1, item["produk"], jumlah, subtotal)
		totalHarga += subtotal
	}
	fmt.Printf("\n  💰 Total: Rp%v\n", totalHarga)

}

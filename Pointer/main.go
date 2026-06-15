package main

import "fmt"

func main() {
	// ==================== POINTER BASICS ====================
	// Pointer adalah variabel yang menyimpan ALAMAT MEMORY dari variabel lain
	// Bukan nilai nya, tapi alamat dimana nilai itu disimpan di memory
	//
	// OPERATOR POINTER:
	// & (Ampersand) = Ambil ALAMAT dari variabel
	// * (Asterisk)  = Akses NILAI dari alamat (dereference)

	// ==================== CONTOH 1: Pointer String ====================
	nama := "Budi"
	// Variabel biasa: nama menyimpan nilai "Budi"

	ptrNama := &nama
	// Pointer: ptrNama menyimpan ALAMAT dari variabel nama, bukan nilai "Budi"

	fmt.Println("=== Contoh 1: Pointer String ===")
	fmt.Printf("Nilai: nama = %s\n", nama)
	fmt.Printf("Alamat: &nama = %p\n", &nama)
	fmt.Printf("Pointer: ptrNama = %p\n", ptrNama)
	fmt.Printf("Dereference: *ptrNama = %s\n\n", *ptrNama)

	// ==================== CONTOH 2: Pointer Integer ====================
	angka := 42
	// ptrAngka menunjuk ke alamat variabel angka

	ptrAngka := &angka

	fmt.Println("=== Contoh 2: Pointer Integer ===")
	fmt.Printf("Nilai: angka = %d\n", angka)
	fmt.Printf("Pointer: ptrAngka = %p\n", ptrAngka)
	fmt.Printf("Dereference: *ptrAngka = %d\n\n", *ptrAngka)

	// ==================== CONTOH 3: Mengubah Nilai Melalui Pointer ====================
	// PENTING: Mengubah nilai melalui pointer akan mengubah variabel ASLI juga!

	umur := 20
	ptrUmur := &umur

	fmt.Println("=== Contoh 3: Modifikasi Nilai via Pointer ===")
	fmt.Printf("Umur awal: %d\n", umur)

	// Ubah nilai melalui pointer dengan *ptrUmur = nilai_baru
	*ptrUmur = 25

	fmt.Printf("Setelah *ptrUmur = 25:\n")
	fmt.Printf("  umur = %d (BERUBAH!)\n", umur)
	fmt.Printf("  *ptrUmur = %d\n\n", *ptrUmur)

	// ==================== CONTOH 4: Pointer Nil ====================
	// Pointer yang tidak menunjuk ke variabel manapun adalah nil (kosong)

	var ptrKosong *string
	fmt.Println("=== Contoh 4: Pointer Nil ===")
	fmt.Printf("Pointer kosong: %v\n", ptrKosong)
	fmt.Println("(Pointer kosong tidak menunjuk ke variabel apapun)")

	// ==================== CONTOH 5: Pass by Value ====================
	// Kalau pass parameter tanpa pointer, nilai tidak berubah (copy value)

	nilai := 100
	fmt.Println("=== Contoh 5: Pass by Value ===")
	fmt.Printf("Nilai awal: %d\n", nilai)

	ubahNilaiByValue(nilai)
	fmt.Printf("Setelah function: %d (tidak berubah)\n\n", nilai)

	// ==================== CONTOH 6: Pass by Reference ====================
	// Kalau pass pointer, bisa mengubah nilai variabel asli (pass reference)

	nilai2 := 100
	fmt.Println("=== Contoh 6: Pass by Reference (Pointer) ===")
	fmt.Printf("Nilai awal: %d\n", nilai2)

	ubahNilaiByReference(&nilai2)
	fmt.Printf("Setelah function: %d (BERUBAH!)\n\n", nilai2)

	// ==================== CONTOH 7: Pointer ke Struct ====================
	// Bisa membuat pointer yang menunjuk ke struct

	type Siswa struct {
		nama  string
		umur  int
		kelas string
	}

	siswa := Siswa{nama: "Ali", umur: 17, kelas: "XI"}
	ptrSiswa := &siswa

	fmt.Println("=== Contoh 7: Pointer ke Struct ===")
	fmt.Printf("Struct: %v\n", siswa)
	fmt.Printf("Akses field via pointer: ptrSiswa.nama = %s\n", ptrSiswa.nama)

	// Ubah field melalui pointer
	ptrSiswa.umur = 18
	fmt.Printf("Setelah ptrSiswa.umur = 18: siswa.umur = %d\n\n", siswa.umur)

	// ==================== CONTOH 8: Pointer ke Array/Slice ====================
	// Bisa membuat pointer ke array atau slice

	angkaArray := [5]int{1, 2, 3, 4, 5}
	ptrArray := &angkaArray

	fmt.Println("=== Contoh 8: Pointer ke Array ===")
	fmt.Printf("Array: %v\n", angkaArray)
	fmt.Printf("Pointer: %p\n", ptrArray)
	fmt.Printf("Akses via pointer: ptrArray[0] = %d\n", ptrArray[0])

	// ========== PENJELASAN: RETURN vs POINTER ==========
	// PERTANYAAN: Kenapa pakai pointer? Kan bisa pakai return juga?
	// JAWABAN: Dua-duanya bisa! Tapi ada perbedaan dan use case berbeda.

	fmt.Println("\n=== PERBANDINGAN: Return Value vs Pointer ===")

	student := Student{1, "Aminudin", 3.75}
	fmt.Printf("Nama awal: %s\n", student.Name)

	// PENDEKATAN 1: Menggunakan Return (Pass by Value + Return)
	fmt.Println("\n--- Pendekatan 1: Return Value ---")
	student.Name = graduateWithReturn(student)
	// Harus assign ulang hasil return ke variabel
	fmt.Printf("Nama setelah return: %s\n", student.Name)

	// PENDEKATAN 2: Menggunakan Pointer (Pass by Reference)
	fmt.Println("\n--- Pendekatan 2: Pointer (Pass by Reference) ---")
	student2 := Student{2, "Budi", 3.50}
	fmt.Printf("Nama awal: %s\n", student2.Name)
	graduateWithPointer(&student2)
	// Tidak perlu assign ulang, langsung berubah
	fmt.Printf("Nama setelah pointer: %s\n", student2.Name)

	// ===== PENJELASAN PERBEDAAN =====
	fmt.Println("\n=== PERBEDAAN ===")
	fmt.Println("RETURN VALUE:")
	fmt.Println("  • Function harus return nilai")
	fmt.Println("  • Caller harus assign ulang: student.Name = function(...)")
	fmt.Println("  • Hanya bisa ubah 1 field per function (return 1 nilai)")
	fmt.Println("  • Cocok untuk operasi sederhana")

	fmt.Println("\nPOINTER:")
	fmt.Println("  • Function tidak perlu return apapun")
	fmt.Println("  • Perubahan langsung ke variabel asli")
	fmt.Println("  • Bisa ubah multiple fields sekaligus")
	fmt.Println("  • Cocok untuk struct besar atau modifikasi kompleks")
}

// Pass by Value: Parameter adalah COPY dari nilai asli
// Perubahan di dalam function TIDAK mempengaruhi variabel asli
func ubahNilaiByValue(nilai int) {
	nilai = 200
	// Ini hanya mengubah copy, bukan variabel asli
}

// Pass by Reference: Parameter adalah POINTER ke variabel asli
// Perubahan di dalam function MEMPENGARUHI variabel asli
func ubahNilaiByReference(ptr *int) {
	*ptr = 300
	// Ini mengubah variabel asli karena kita mengubah melalui pointer
}

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// PENDEKATAN 1: Return Value (hanya bisa return 1 nilai)
func graduateWithReturn(student Student) string {
	// Return nama yang sudah ditambah "S.T"
	return student.Name + " S.T"
}

// PENDEKATAN 2: Pointer (bisa ubah langsung struct)
func graduateWithPointer(student *Student) {
	// Ubah langsung field dari pointer
	student.Name = student.Name + " S.T"
}

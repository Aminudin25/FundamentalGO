package main

import "fmt"

// =====================================================
// APA ITU INTERFACE?
// =====================================================
// Interface adalah sebuah "kontrak" yang mendefinisikan method apa saja
// yang harus diimplementasikan oleh sebuah tipe data.
//
// Interface TIDAK mengimplementasikan method, hanya MENDEFINISIKAN
// method apa yang harus ada.
//
// Analogi:
// - Variabel biasa: Rumah dengan ukuran tetap
// - Interface: Persyaratan untuk membangun rumah
//   (harus punya 2 jendela, 1 pintu, dll)
//   siapa saja bisa membangun rumah sesuai persyaratan ini
//
// =====================================================
// ATURAN MAIN INTERFACE:
// =====================================================
// 1. Interface HANYA berisi method definition (signature)
// 2. Interface TIDAK boleh punya field/property
// 3. Method interface TIDAK boleh punya body (implementasi)
// 4. Type apa saja bisa implement interface, asalkan punya semua method
// 5. Tidak perlu deklarasi explicit "implements"
//    Go menggunakan implicit interface (duck typing)
// 6. Bisa buat interface kosong (interface{}) - bisa diterima tipe apa saja
//
// =====================================================

// ==================== CONTOH 1: Interface Dasar ====================
// Interface "Writer" mendefinisikan 1 method: Write
type Writer interface {
	Write(data string) string
}

// Struct File implement interface Writer
type File struct {
	name string
}

func (f File) Write(data string) string {
	// Implementasi Write untuk File
	return "File " + f.name + " menulis: " + data
}

// Struct Database juga implement interface Writer
type Database struct {
	host string
}

func (d Database) Write(data string) string {
	// Implementasi Write untuk Database
	return "Database " + d.host + " menulis: " + data
}

// ==================== CONTOH 2: Interface dengan Multiple Methods ====================
// Interface "Shape" mendefinisikan 2 methods: Area dan Perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct Circle implement interface Shape
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// Struct Rectangle juga implement interface Shape
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// ==================== CONTOH 3: Interface Kosong ====================
// Interface kosong (interface{}) bisa menerima value apa saja
// Cocok untuk function yang parameter-nya bisa tipe apa saja
type Container interface{}

// ==================== CONTOH 4: Embedded Interface ====================
// Interface bisa embed interface lain
type Reader interface {
	Read() string
}

type ReadWriter interface {
	Reader // Embed interface Reader
	Write(data string) string
}

type FileRW struct {
	name string
}

func (f FileRW) Read() string {
	return "File " + f.name + " dibaca"
}

func (f FileRW) Write(data string) string {
	return "File " + f.name + " menulis: " + data
}

// =====================================================
// KENAPA HARUS MENGGUNAKAN INTERFACE?
// =====================================================
// 1. POLYMORPHISM: Bisa treat berbagai type dengan cara yang sama
// 2. ABSTRACTION: Interface menyembunyikan detail implementasi
// 3. FLEXIBILITY: Mudah menambah tipe baru tanpa ubah existing code
// 4. TESTABILITY: Mudah membuat mock object untuk testing
// 5. LOOSE COUPLING: Code jadi lebih fleksibel dan maintainable
//
// CONTOH NYATA:
// Bayangkan function printData yang bisa menerima File, Database, Network
// Tanpa interface, harus buat 3 function terpisah
// Dengan interface, hanya 1 function untuk ketiga-tiganya!

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║              PENJELASAN LENGKAP TENTANG INTERFACE           ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	// ==================== CONTOH 1: Interface Writer ====================
	fmt.Println("=== Contoh 1: Interface Writer ===")
	fmt.Println("Interface Writer mewajibkan implementasi method Write()")

	file := File{name: "data.txt"}
	database := Database{host: "localhost"}

	// Kedua struct berbeda, tapi punya method Write yang sama
	fmt.Println(file.Write("Hello World"))
	fmt.Println(database.Write("SELECT * FROM users"))

	// Bisa pakai interface untuk treat kedua-duanya sama
	var writers []Writer
	writers = append(writers, file, database)

	fmt.Println("\nMenggunakan interface Writer:")
	for _, w := range writers {
		fmt.Println(w.Write("Test data"))
	}

	// ==================== CONTOH 2: Interface Shape ====================
	fmt.Println("=== Contoh 2: Interface Shape ===")
	fmt.Println("Interface Shape mewajibkan implementasi Area() dan Perimeter()")

	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 7}

	// Function yang accept interface Shape
	printShapeInfo(circle)
	printShapeInfo(rectangle)

	// Bisa simpan dalam slice Shape
	fmt.Println("\nDaftar shapes:")
	var shapes []Shape
	shapes = append(shapes, circle, rectangle)

	for i, shape := range shapes {
		fmt.Printf("Shape %d - Area: %.2f, Perimeter: %.2f", i+1, shape.Area(), shape.Perimeter())
	}

	// ==================== CONTOH 3: Interface Kosong ====================
	fmt.Println("=== Contoh 3: Interface Kosong (interface{}) ===")
	fmt.Println("Interface{} bisa terima nilai apa saja:")

	var anything interface{}

	anything = "String"
	fmt.Printf("Nilai: %v, Tipe: %T", anything, anything)

	anything = 42
	fmt.Printf("Nilai: %v, Tipe: %T", anything, anything)

	anything = 3.14
	fmt.Printf("Nilai: %v, Tipe: %T", anything, anything)

	anything = true
	fmt.Printf("Nilai: %v, Tipe: %T", anything, anything)

	// ==================== CONTOH 4: Type Assertion ====================
	fmt.Println("=== Contoh 4: Type Assertion ===")
	fmt.Println("Type assertion untuk konversi interface ke tipe asli:")

	var w Writer = File{name: "myfile.txt"}

	// Type assertion
	f, ok := w.(File)
	if ok {
		fmt.Printf("Berhasil konversi ke File: %v", f)
	}

	d, ok := w.(Database)
	if ok {
		fmt.Printf("Berhasil konversi ke Database: %v", d)
	} else {
		fmt.Println("Tidak bisa konversi ke Database")
	}

	// ==================== CONTOH 5: Multiple Interface ====================
	fmt.Println("=== Contoh 5: Type Bisa Implement Multiple Interface ===")

	fileRW := FileRW{name: "document.txt"}

	// FileRW implement interface Reader dan Writer
	var r Reader = fileRW
	var w2 Writer = fileRW

	fmt.Println(r.Read())
	fmt.Println(w2.Write("New content"))

	// ==================== RINGKASAN ====================
	fmt.Println("=== RINGKASAN ===")
	// 	fmt.Println(`
	// INTERFACE ADALAH KONTRAK YANG MENDEFINISIKAN METHOD:

	// Syntax Interface:
	//   type InterfaceName interface {
	//       MethodName1(param1 type1) returnType1
	//       MethodName2(param2 type2) returnType2
	//   }

	// Implementing Interface:
	//   type MyType struct { ... }
	//   func (m MyType) MethodName1(param1 type1) returnType1 { ... }
	//   func (m MyType) MethodName2(param2 type2) returnType2 { ... }

	// KEUNTUNGAN:
	//   ✅ Polymorphism - berbagai type dalam satu interface
	//   ✅ Abstraction - hide implementation details
	//   ✅ Flexibility - mudah extend tanpa ubah existing code
	//   ✅ Testability - mudah buat mock untuk testing
	//   ✅ Loose Coupling - code jadi modular
	// `)
}

// Function yang accept interface Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f", s.Area(), s.Perimeter())
}

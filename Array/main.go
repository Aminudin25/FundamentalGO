package main

import "fmt"

func main() {

	// array yang didefinisikan jumlah value nya
	var languages [5]string
	// isi setiap valuenya
	// array dimulai dari 0
	languages[0] = "GO"
	languages[1] = "Python"
	languages[2] = "JavaScript"
	languages[3] = "C#"
	languages[4] = "Java"

	fmt.Println(languages)

	// karna kita mendefinisikan array dengan jumlah valuenya 5
	// maka array itu tidak bisa menampung lebih dari 5
	// coba kita tambahkan lagi
	// languages[5] = "PHP"
	// terdapat error bahwa array out of bound

	// kita bisa cek juga jumlah value dari aray
	lengthArray := len(languages)
	fmt.Println(lengthArray)
	// outpunya akan menampilkan "5"

	// Bagaimana kita cara membuat Array dinamis utk jumlah value nya ?
	dynamicArray := [...]string{"Satu", "Dua", "Tiga", "Empat"}
	fmt.Println(dynamicArray)

	// kita coba kombinasi dengan fungsi looping untuk mengambil setiap value yang ada di array

	for index, value := range dynamicArray {
		fmt.Println("index :", index, "value :", value)
	}

	Slice()
	Map()
}

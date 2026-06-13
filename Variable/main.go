package main

import "fmt"

// golang itu Static Type, sama seperti C# & Java
// maksudnya apaa ?
// ketika kita sudah mendefinisikan sebuah variable dengan tipe data tertentu
// lalu kita memasukan value nya tidak sesuai dengan tipe data nya, maka akan error
// mangkanya kenapa GO itu Static Type
// bukan Dynamic Type seperti PHP, JavaScript, Rubby dll

// TAPI GO juga bisa membaca tipe data sesuai dengan value nya, jadi ketika value nya diganti menjadi 'String' maka variable nya menjadi String, begitupun ketika kita merubah nya ke tipe data lain.
// CONTOH Pembuatanya :
// DynamicType := 10. ketika kita ganti value nya menjadi STRING, maka tipe data nya juga akan berubah menjadi STRING

func main() {
	// mendefinisikan variable yang langsung di assign value nya
	var name string = "Golang"
	fmt.Println(name)

	// mendefinisikan variable tanpa di assign value nya
	var address string
	// asign value ke dalam variable address
	address = "Jalan Merdeka"
	fmt.Println(address)

	//merubah value dari variable yang sudah di asign
	name = "Aminudin"
	fmt.Println(name)

	//mendefinisikan variable secara dynamic tanpa harus menuliskan tipe data nya
	dynamicType := "String"
	fmt.Println(dynamicType)
}

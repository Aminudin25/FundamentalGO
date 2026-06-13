package main

import "fmt"

// kita membuat file entity.go dengan package main
// sehingga file entity.go masih berhubungan dan atau sama dengan file main.go
// karena package nya masih sama walaupun beda file
// sebagai contoh, kita membuat function di file entity.go
// function ini mengembalikan string, sehingga harus menggunakan return dan fmt.Sprintf
// %s dan %d untuk menyisipkan data string (%s) dan int (%d)

func TestAja() string {
	return fmt.Sprintf("Halo nama saya %s dan umur saya %d", "Aminudin", 26)
}

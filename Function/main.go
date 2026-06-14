package main

import "fmt"

func main() {
	printMyResult()
	sentencePrint("Belajar")
	name := outputPrint("Aminudin")
	fmt.Println(name)

	// karna mengembalikan return (int), maka harus ditampung
	calculatorAdd := add(5, 5)
	fmt.Println(calculatorAdd)

	calculatorX := multiplication(5, 5)
	fmt.Println(calculatorX)

	Quiz()
}

// contoh function sederhana, tidak ada input ataupun output
// di dunia nyata biasanya terdapat input maupun output
func printMyResult() {
	fmt.Println("saya sedang belajar GO")
}

// kita buat yang mempunyai :
// 1. input
// 2. proses

func sentencePrint(sentence string) {
	fmt.Println(sentence)
}

//kita buat yang terdapat output
func outputPrint(name string) string {
	newSentence := "Hallo, " + name + " Selamat Datang"

	return newSentence
}

func add(num1 int, num2 int) int {
	// cara pertama
	// result := num1 + num2
	// return result

	// cara kedua
	return num1 + num2
}

func multiplication(num1, num2 int) int {
	return num1 * num2
}

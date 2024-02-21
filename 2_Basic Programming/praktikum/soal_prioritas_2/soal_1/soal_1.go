package main

import "fmt"

func main() {
	fmt.Println("===========================================================================")
	fmt.Println("Program untuk mencetak sebuah pola berdasarkan faktor bilangan")
	fmt.Println("Jika faktor bilangan merupakan bilangan genap maka gunakan simbol “I” (huruf i besar) dan selain itu menggunakan simbol huruf O.")
	fmt.Println("===========================================================================")

	bilangan := 26

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			if i%2 == 0 {
				fmt.Print("I")
			} else {
				fmt.Print("O")
			}
		}
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Program untuk menentukan grade dari sebuah nilai")

	var nilai int8
	fmt.Print("Masukkan Nilai \t: ")
	_, err := fmt.Scan(&nilai)
	if err != nil {
		fmt.Println("Nilai Invalid")
		return
	}
	fmt.Print("Nilai anda \t: ")
	switch {
	case nilai >= 85 && nilai <= 100:
		fmt.Print("A")
		break
	case nilai >= 70 && nilai <= 84:
		fmt.Print("B")
		break
	case nilai >= 85 && nilai <= 69:
		fmt.Print("C")
		break
	case nilai >= 40 && nilai <= 50:
		fmt.Print("D")
		break
	case nilai >= 0 && nilai <= 39:
		fmt.Print("E")
		break
	default:
		fmt.Println("Nilai Invalid")
	}
}

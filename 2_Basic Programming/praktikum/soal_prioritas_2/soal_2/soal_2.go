package main

import "fmt"

func main() {
	fmt.Println("program untuk menentukan prioritas dari sebuah proyek")
	fmt.Println("======================================================")
	var (
		budget     int
		time       int
		difficulty int
		score      int
		result     string
	)

	fmt.Print("Masukkan Budget: ")
	fmt.Scan(&budget)
	switch {
	case budget >= 0 && budget <= 50:
		score += 4
		break
	case budget >= 51 && budget <= 80:
		score += 3
		break
	case budget >= 81 && budget <= 100:
		score += 2
		break
	case budget > 100:
		score += 1
		break
	default:
		fmt.Println("Nilai Invalid")
		return
	}

	fmt.Print("Masukkan waktu pengerjaaan: ")
	fmt.Scan(&time)
	switch {
	case time >= 0 && time <= 20:
		score += 10
		break
	case time >= 21 && time <= 30:
		score += 5
		break
	case time >= 31 && time <= 50:
		score += 2
		break
	case time > 50:
		score += 1
		break
	default:
		fmt.Println("Nilai Invalid")
		return
	}

	fmt.Print("Masukkan kesulitan: ")
	fmt.Scan(&difficulty)
	switch {
	case difficulty >= 0 && difficulty <= 3:
		score += 10
		break
	case difficulty >= 4 && difficulty <= 6:
		score += 5
		break
	case difficulty >= 8 && difficulty <= 10:
		score += 1
		break
	case difficulty > 10:
		score += 0
		break
	default:
		fmt.Println("Nilai Invalid")
		return
	}

	switch {
	case score >= 17 && score <= 24:
		result = "High"
		break
	case score >= 10 && score <= 16:
		result = "Medium"
		break
	case score >= 3 && score <= 9:
		result = "Low"
		break
	case score <= 2:
		result = "Impossible"
		break
	default:
		fmt.Println("Error")
		return
	}
	fmt.Printf("Output: %s", result)

}

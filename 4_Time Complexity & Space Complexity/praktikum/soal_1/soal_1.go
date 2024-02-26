package main

func primeNumber(number int) bool {

	if number <= 1 {
		return false
	}

	if number == 2 || number == 3 || number == 5 || number == 7 {
		return true
	}

	if number%2 == 0 || number%3 == 0 || number%5 == 0 || number%7 == 0 {
		return false
	}
	return true

}

func main() {
	bil1 := primeNumber(1000000007)
	bil2 := primeNumber(1500450271)

	if bil1 {
		println("bilangan prima")
	} else {
		println("bukan bilangan prima")
	}

	if bil2 {
		println("bilangan prima")
	} else {
		println("bukan bilangan prima")
	}

}

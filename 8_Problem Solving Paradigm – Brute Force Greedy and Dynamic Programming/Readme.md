# Problem Solving Paradigm – Brute Force Greedy and Dynamic Programming :rocket:

> Problem Solving Paradigm mengacu pada pendekatan atau metodologi yang digunakan untuk memecahkan masalah. Ada beberapa paradigma yang biasa digunakan dalam ilmu komputer dan pemrograman. Paradigma itu terdiri dari Brute Force, Divide and Conquer, Greedy, dan Dynamic Programming. Setiap paradigma akan didefinisikan, dan contoh-contoh akan diberikan untuk mengilustrasikan penggunaannya

## Complete Search/Brute Force

> Complete Search, juga dikenal sebagai Brute Force, adalah metode langsung untuk memecahkan masalah dengan mencoba semua solusi yang mungkin. Metode ini melibatkan pengulangan melalui semua solusi yang mungkin dan memilih yang terbaik.

Contoh: **Find Max and Min**
Find Max and Min adalah contoh complete search. Dalam kasus mencari nilai maksimum (max) dan minimum (min) dari suatu himpunan data, complete search akan memeriksa setiap elemen untuk menemukan nilai maksimum dan minimum.

```
func findMaxMin(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

	maxVal := arr[0]
	minVal := arr[0]

	for _, num := range arr {
		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	return maxVal, minVal
}
```

## Divide and Conquer

> Divide and Conquer adalah paradigma pemecahan masalah yang melibatkan pemecahan masalah menjadi bagian-bagian yang lebih kecil dan lebih mudah dikelola, menyelesaikan setiap bagian secara independen, dan kemudian menggabungkan solusi untuk menyelesaikan masalah aslinya.

Contoh: **Power Function**
Power Function adalah contoh dari paradigma Divide and Conquer. Paradigma ini menghitung nilai sebuah bilangan yang dipangkatkan dengan eksponen dengan membagi eksponen tersebut secara rekursif menjadi dua dan menghitung hasilnya untuk setiap bagian, lalu menggabungkan hasilnya.

```
func power(x, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		partial := power(x, n/2)
		return partial * partial
	} else {
		partial := power(x, (n-1)/2)
		return x * partial * partial
	}
}
```

Contoh: **Binary Search**

```
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
```

## Greedy

> Greedy adalah paradigma pemecahan masalah yang melibatkan pembuatan pilihan optimal secara lokal pada setiap langkah dengan harapan menemukan optimal global. Paradigma ini tidak selalu menjamin solusi yang optimal, tetapi dapat menjadi pendekatan yang sederhana dan efisien untuk banyak masalah.

Contoh: **Coin Change**
Coin Change adalah contoh dari paradigma Greedy. Paradigma ini melibatkan pencarian jumlah minimum koin yang dibutuhkan untuk menghasilkan sejumlah uang kembalian. Pendekatan greedy melibatkan pemilihan koin denominasi terbesar pada setiap langkah sampai jumlah total tercapai.

```
func coinChange(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	count := 0
	for _, coin := range coins {
		if amount >= coin {
			count += amount / coin
			amount %= coin
		}
	}
	if amount != 0 {
		return -1
	}
	return count
}
```

## Dynamic Programming

> Dynamic Programming adalah paradigma pemecahan masalah yang melibatkan pemecahan masalah menjadi submasalah yang lebih kecil yang saling tumpang tindih dan menyelesaikan setiap submasalah hanya sekali. Solusi dari submasalah disimpan dan digunakan kembali untuk memecahkan submasalah yang lebih besar, sehingga menghasilkan solusi yang lebih efisien.

Karakteristik:

- Overlapping Subproblems: Masalah dapat dipecah menjadi submasalah yang lebih kecil yang dapat digunakan kembali beberapa kali.

- Optimal Substructurel: Solusi optimal untuk masalah dapat dibangun dari solusi optimal untuk submasalah.

Metode:

- Top-Down with Memoization: Dalam pendekatan ini, masalah diselesaikan secara rekursif, tetapi solusi untuk submasalah disimpan dalam tabel (memoisasi) untuk menghindari perhitungan yang berlebihan.

- Bottom-Up with Tabulation: Dalam pendekatan ini, masalah diselesaikan secara iteratif mulai dari submasalah yang paling kecil dan meningkat ke masalah yang lebih besar. Solusi dari submasalah disimpan dalam sebuah tabel dan digunakan untuk menyelesaikan submasalah yang lebih besar.

Contoh: Fibonacci Sequence
Fibonacci Sequence adalah contoh masalah yang bisa diselesaikan dengan menggunakan Pemrograman Dinamis. Deret ini didefinisikan secara rekursif sebagai F(n) = F(n-1) + F(n-2), dengan kasus dasar F(0) = 0 dan F(1) = 1. Dengan menyimpan solusi-solusi untuk submasalah (nilai-nilai F(n-1) dan F(n-2)) dalam sebuah larik, kita bisa menghindari penghitungan ulang dan meningkatkan efisiensi algoritma.

```
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}
```

# Thank You :star2:

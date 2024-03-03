# String, Recursive, Number Theory, Sorting, Searching :rocket:

## Recursive

> Rekursi adalah teknik di mana sebuah fungsi memanggil dirinya sendiri untuk menyelesaikan masalah. Di Go, seperti dalam banyak bahasa pemrograman lainnya, Anda dapat menggunakan rekursi untuk mengimplementasikan solusi yang sederhana dan elegan untuk beberapa masalah, terutama yang terkait dengan struktur data rekursif seperti pohon atau daftar terhubung.

Kita membutuhkan rekursi karena ada banyak masalah di mana pendekatan rekursif dapat memberikan solusi yang lebih sederhana dan mudah dipahami daripada pendekatan iteratif. Rekursi terutama bermanfaat dalam kasus-kasus di mana masalah yang dihadapi dapat dipecahkan menjadi submasalah yang serupa, dan solusi untuk setiap submasalah dapat digunakan untuk memecahkan masalah yang lebih besar.

**Contoh Penggunaan rekursif:**

1. Pencarian dalam struktur data rekursif: Misalnya, dalam struktur data pohon, pencarian nilai tertentu dapat dilakukan dengan cara rekursif, dimulai dari akar pohon dan kemudian secara rekursif mencari nilai di subpohon kiri dan kanan.

2. Faktorial: Seperti yang telah dijelaskan sebelumnya, menghitung faktorial dari sebuah bilangan adalah contoh klasik penggunaan rekursi.

3. Algoritma Divide and Conquer: Rekursi sering digunakan dalam algoritma divide and conquer, di mana masalah dibagi menjadi submasalah yang lebih kecil, dipecahkan secara rekursif, dan kemudian digabungkan kembali untuk mendapatkan solusi akhir. Contoh dari algoritma ini adalah Merge Sort dan Quick Sort.

Strategi rekursi melibatkan dua konsep kunci: base case (kasus dasar) dan recurrence relation (hubungan rekurens).

1. Base Case (Kasus Dasar): Base case adalah kondisi di mana rekursi berhenti. Ini adalah kasus yang paling sederhana dan biasanya memiliki solusi langsung yang diketahui tanpa memerlukan rekursi. Tanpa base case, rekursi akan terus berlanjut secara tak terbatas, yang dapat menyebabkan stack overflow.

2. Recurrence Relation (Hubungan Rekurens): Recurrence relation adalah cara untuk menggambarkan hubungan antara masalah asli dan submasalah yang lebih kecil. Ini adalah langkah rekursif di mana fungsi memanggil dirinya sendiri dengan argumen yang berbeda untuk memecahkan submasalah.

**Contoh Program:**

```
package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5)) // Output: 120
}
```

## Number Theory

> Number theory adalah cabang matematika yang mempelajari sifat-sifat dasar dari bilangan bulat, termasuk operasi aritmetika, pembagian, faktorisasi, bilangan prima, dan fungsi aritmetika. Teori bilangan mencakup konsep-konsep seperti teorema dasar aritmetika yang menyatakan bahwa setiap bilangan bulat dapat diuraikan menjadi faktor-faktor prima secara unik, serta konsep-konsep lanjutan seperti persamaan diofantin, kongruensi, dan hukum reciprocitas kuadrat. Aplikasi dari teori bilangan meliputi bidang kriptografi, teori bilangan komputasional, dan berbagai bidang matematika lainnya.

**Faktorial:**

```
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

```

**Bilangan Prima:**

```
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    maxDivisor := int(math.Sqrt(float64(n)))
    for i := 2; i <= maxDivisor; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
```

**Greatest Common Divisor / FPB:**

```
func gcd(a int, b int) int {
   if b == 0 {
    return a
   }
   return gcd(b, a % b)
}
```

**Least Common Multiple / KPK:**

```
func lcm(a int, b int) int {
  return a * (b / lcm(a,b))
}
```

## Searching

> Searching adalah proses untuk mencari posisi nilai dari sebuah daftar nilai.

### Linear Search

> Linear search adalah metode pencarian sederhana yang digunakan untuk mencari sebuah elemen dalam array atau daftar linier secara berurutan. Algoritma ini bekerja dengan cara memeriksa setiap elemen satu per satu, mulai dari elemen pertama hingga elemen terakhir, sampai elemen yang dicari ditemukan atau seluruh array telah diperiksa. Jika elemen yang dicari ditemukan, linear search mengembalikan indeks elemen tersebut; jika tidak ditemukan, ia mengembalikan nilai yang menandakan bahwa elemen tidak ada dalam array.

Contoh :

```
func linearSearch(arr []int, target int) int {
    for i, num := range arr {
        if num == target {
            return i
        }
    }
    return -1
```

### Builtins Search

> Package sort di Go menyediakan fungsi SearchInts yang dapat digunakan untuk melakukan pencarian biner pada slice []int yang sudah terurut. Berikut adalah contoh program yang menggunakan sort.SearchInts untuk mencari elemen dalam slice:

Contoh :

```
package main

import (
    "fmt"
    "sort"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    target := 7

    index := sort.SearchInts(numbers, target)

    if index < len(numbers) && numbers[index] == target {
        fmt.Printf("Elemen %d ditemukan di indeks %d\n", target, index)
    } else {
        fmt.Printf("Elemen %d tidak ditemukan\n", target)
    }
}
```

## Sorting

> Sorting adalah proses mengatur kumpulan data dalam urutan tertentu, biasanya berdasarkan aturan atau kriteria tertentu. Tujuan utama pengurutan adalah untuk membuat data lebih mudah dicari, dibaca, atau diproses secara efisien. Dalam konteks Go, sorting mengacu pada proses mengatur elemen-elemen dalam slice atau array dalam urutan tertentu, seperti ascending (dari kecil ke besar) atau descending (dari besar ke kecil). Proses pengurutan ini biasanya dilakukan menggunakan algoritma pengurutan seperti quicksort, mergesort, atau heapsort, yang diimplementasikan dalam package sort di Go.

### Selection Sort - O(n^2)

> Selection sort adalah algoritma pengurutan sederhana yang secara berulang memilih elemen dengan nilai minimum (atau maksimum) dari sisa array yang belum diurutkan, dan menukarnya dengan elemen pertama (atau terakhir) dari sisa array yang belum diurutkan. Algoritma ini memiliki kompleksitas waktu O(n^2), di mana n adalah jumlah elemen dalam array.

Langkah-langkah selection sort adalah sebagai berikut:

1. Mulai dari indeks 0 hingga n-1, cari elemen terkecil dalam array.
2. Tukar elemen terkecil dengan elemen pada indeks pertama.
3. Ulangi langkah 1 dan 2 untuk sisa array, yaitu dari indeks 1 hingga n-1, cari elemen terkecil dalam array (tanpa memperhatikan elemen di indeks 0).
4. Ulangi proses ini hingga seluruh array diurutkan.

```
func selectionSort(arr []int) []int {
    n := len(arr)
    sortedArr := make([]int, n)
    copy(sortedArr, arr)

    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if sortedArr[j] < sortedArr[minIdx] {
                minIdx = j
            }
        }
        if minIdx != i {
            sortedArr[i], sortedArr[minIdx] = sortedArr[minIdx], sortedArr[i]
        }
    }
    return sortedArr
}
```

### Counting Sort - O(n + k)

> Counting sort adalah algoritma pengurutan yang efisien untuk mengurutkan elemen-elemen dalam array atau slice dengan rentang nilai yang terbatas. Algoritma ini bekerja dengan menghitung jumlah kemunculan setiap nilai unik dalam array, kemudian menggunakan informasi tersebut untuk menempatkan setiap elemen ke posisinya yang benar dalam array hasil.

Langkah-langkah counting sort adalah sebagai berikut:

1. Hitung jumlah kemunculan setiap nilai dalam array dan simpan informasi tersebut dalam array frekuensi.
2. Hitung jumlah kumulatif dari frekuensi sehingga kita dapat menentukan posisi akhir setiap nilai dalam array hasil.
3. Iterasi melalui array asli dan tempatkan setiap elemen ke posisi yang benar dalam array hasil, kemudian kurangi nilai frekuensi yang terkait dengan elemen tersebut.
4. Array hasil sekarang berisi elemen-elemen yang sudah diurutkan.

```
func countingSort(arr []int, maxValue int) []int {
    n := len(arr)
    sortedArr := make([]int, n)
    count := make([]int, maxValue+1)

    for i := 0; i < n; i++ {
        count[arr[i]]++
    }

    pos := 0
    for i := 0; i <= maxValue; i++ {
        for j := 0; j < count[i]; j++ {
            sortedArr[pos] = i
            pos++
        }
    }

    return sortedArr
}
```

### Merge Sort - O(log n)

> Merge sort adalah algoritma pengurutan yang menggunakan pendekatan divide and conquer untuk mengurutkan sebuah array. Algoritma ini bekerja dengan cara membagi array menjadi dua bagian yang sama besar, mengurutkan kedua bagian secara terpisah, dan kemudian menggabungkan kedua bagian yang sudah diurutkan tersebut. Merge sort memiliki kompleksitas waktu O(n log n), di mana n adalah jumlah elemen dalam array.

Langkah-langkah merge sort adalah sebagai berikut:

1. Jika panjang array <= 1, kembalikan array tersebut (kasus dasar).
2. Bagi array menjadi dua bagian yang sama besar.
3. Panggil merge sort secara rekursif untuk kedua bagian.
4. Gabungkan kedua bagian yang sudah diurutkan menjadi satu array hasil.

```
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    middle := len(arr) / 2
    left := mergeSort(arr[:middle])
    right := mergeSort(arr[middle:])

    return merge(left, right)
}

func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0

    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}
```

## Builtins Sort

> Package sort di Go menyediakan beberapa fungsi bawaan yang memungkinkan pengguna untuk mengurutkan slice dengan tipe data tertentu. Berikut adalah penjelasan singkat tentang tiga fungsi bawaan yang umum digunakan: sort.Strings, sort.Ints, dan sort.IntsAreSorted.

**1. sort.Strings:** Fungsi ini digunakan untuk mengurutkan slice yang berisi string secara ascending (dari A ke Z). Fungsi ini mengubah slice tersebut secara langsung, tanpa mengembalikan nilai.

```
func main() {
    fruits := []string{"apple", "orange", "banana", "grape"}
    sort.Strings(fruits)
    fmt.Println(fruits)
}
```

**2. sort.Ints:** Fungsi ini digunakan untuk mengurutkan slice yang berisi bilangan bulat secara ascending (dari kecil ke besar). Fungsi ini mengubah slice tersebut secara langsung, tanpa mengembalikan nilai.

```
func main() {
    numbers := []int{4, 2, 7, 1, 5}
    sort.Ints(numbers)
    fmt.Println(numbers) // Output: [1 2 4 5 7]
}
```

**3. sort.IntsAreSorted:** Fungsi ini digunakan untuk memeriksa apakah slice yang berisi bilangan bulat sudah terurut secara ascending. Fungsi ini mengembalikan nilai true jika slice sudah terurut, dan false jika belum terurut.

```
func main() {
    numbers := []int{1, 2, 4, 5, 7}
    fmt.Println(sort.IntsAreSorted(numbers))

    numbers = []int{4, 2, 7, 1, 5}
    fmt.Println(sort.IntsAreSorted(numbers))
}
```

# Thank You :star2:

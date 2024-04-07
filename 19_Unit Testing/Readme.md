# Unit Testing :rocket:

## Apa itu Unit Testing dan Software Testing

Unit testing adalah teknik Software Testing di mana unit atau komponen individu dari perangkat lunak diuji secara terpisah dari bagian lain dari aplikasi. Ini membantu memastikan bahwa setiap unit perangkat lunak berfungsi seperti yang diharapkan.

Software Testing, di sisi lain, adalah istilah yang lebih luas yang mencakup berbagai teknik pengujian dan metodologi yang bertujuan untuk mengevaluasi kualitas suatu produk perangkat lunak.

## Tujuan Pengujian

Tujuan utama pengujian adalah untuk mengidentifikasi dan memperbaiki bug dalam perangkat lunak, memastikan bahwa perangkat lunak memenuhi persyaratan yang ditentukan, dan berfungsi seperti yang diharapkan. Pengujian juga membantu meningkatkan kualitas keseluruhan perangkat lunak dan mengurangi kemungkinan bug dan kesalahan di masa depan.

## Tingkatan Pengujian

Ada berbagai tingkatan pengujian, termasuk:

1. UI Testing: Menguji antarmuka pengguna perangkat lunak.
2. Integration Testing: Menguji integrasi beberapa unit atau komponen.
3. Unit Testing: Menguji unit atau komponen individu dari perangkat lunak.

## Framework untuk Unit Testing

Di Golang, ada beberapa framework yang tersedia untuk unit testing, termasuk:

- Testing: Paket pengujian bawaan dalam Golang, yang menyediakan dukungan untuk menulis dan menjalankan unit test.
- Testify: Toolkit pengujian yang populer untuk Golang yang menyediakan utilitas tambahan dan asertions untuk menulis tes.

## Pattern dalam Unit Testing

Ada dua pola umum untuk mengorganisir unit test:

1. Centralized: Semua file test ditempatkan dalam satu direktori, terpisah dari kode sumber.
2. Simpan File Uji Bersama: File tes ditempatkan dalam direktori yang sama dengan file sumber, dengan tambahan akhiran \_test pada nama file.

## Test File, Test Suite, Test Fixture, dan Test Cases

- Test File: File Go yang berisi unit test, biasanya dinamai \*\_test.go.
- Test Suite: Kumpulan kasus tes atau fungsi tes yang menguji unit atau komponen tertentu.
- Test Fixture: Persiapan awal yang diperlukan untuk menjalankan kasus tes, seperti membuat objek tiruan atau mengatur lingkungan tes.
- Test Cases: Tes individu yang memverifikasi perilaku unit atau komponen tertentu.

## Mocking

Mocking adalah teknik yang digunakan dalam unit testing untuk menggantikan dependensi sebuah unit dengan objek tiruan. Ini memungkinkan unit diuji secara terisolasi tanpa bergantung pada dependensi sebenarnya.

## Hal yang tidak perlu di uji

Beberapa hal yang biasanya tidak diuji dalam unit test meliputi:

- Dependensi eksternal seperti basis data, API, dan layanan eksternal.
- Fungsionalitas spesifik sistem operasi.
- Elemen antarmuka pengguna yang sulit diotomatisasi.

## Coverage dan Cara Membuatnya

Coverage adalah ukuran seberapa banyak kode sumber yang dieksekusi selama pengujian. Di Golang, coverage dapat dihasilkan menggunakan flag -cover dengan perintah go test. Misalnya:

```
go test -cover ./...
```

ini akan menghasilkan laporan coverage yang menunjukkan persentase coverage kode untuk setiap paket dalam proyek.

# Thank You :star2:

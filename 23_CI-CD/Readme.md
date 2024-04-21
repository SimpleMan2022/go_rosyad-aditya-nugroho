# CI/CD :rocket:

## Apa itu CI/CD

CI/CD adalah singkatan dari Continuous Integration dan Continuous Deployment (atau Continuous Delivery). Ini adalah praktik pengembangan perangkat lunak yang memungkinkan tim pengembang untuk secara terus-menerus mengintegrasikan perubahan kode ke dalam repositori kode bersama secara otomatis, serta secara otomatis melakukan proses pengujian dan penyebaran kode ke lingkungan produksi atau staging.

**Continuous Integration (CI)** berfokus pada penggabungan kode secara teratur ke dalam repositori bersama, di mana setiap perubahan kode akan diperiksa dan diuji secara otomatis, sehingga memastikan bahwa kode tersebut tetap berfungsi dan berintegrasi dengan baik dengan bagian lain dari proyek.

**Continuous Deployment (CD) atau Continuous Delivery (CD)** berfokus pada otomatisasi proses penyebaran perangkat lunak ke lingkungan produksi setelah berhasil melewati langkah-langkah pengujian yang didefinisikan. Dengan CI/CD, perubahan perangkat lunak dapat dengan cepat dan aman diterapkan ke dalam lingkungan produksi, meningkatkan kecepatan dan keandalan pengiriman perangkat lunak.

## Tools CI/CD

Ada banyak tools yang dapat digunakan untuk menerapkan CI/CD dalam proyek perangkat lunak. Berikut adalah beberapa contoh tools yang umum digunakan:

**1. Jenkins:** Jenkins adalah salah satu alat CI/CD open-source yang paling populer. Ini memungkinkan otomatisasi berbagai tahap pengembangan perangkat lunak, termasuk pengujian, integrasi, dan penyebaran.
**2. GitLab CI/CD:** GitLab menyediakan fitur CI/CD terintegrasi langsung dalam platform GitLab, memungkinkan pengguna untuk membuat dan mengelola pipa CI/CD mereka tanpa perlu mengintegrasikan alat tambahan.
**3. GitHub Actions:** GitHub Actions adalah layanan CI/CD yang terintegrasi langsung dengan GitHub. Ini memungkinkan pengguna untuk membuat dan mengelola alur kerja otomatis untuk menguji, membangun, dan menerapkan perangkat lunak mereka.

Ada banyak lagi alat CI/CD lainnya yang tersedia, dan pilihan terbaik tergantung pada kebutuhan spesifik dan preferensi tim pengembangan

## CI/CD pada Github Actions

Untuk melakukan CI/CD menggunakan GitHub Actions, Anda perlu membuat file konfigurasi khusus yang disebut dengan GitHub Actions workflow. Berikut langkah-langkah umumnya:

**1.Buat File Konfigurasi Workflow:** Buat file dengan nama .github/workflows/<nama-file>.yml di dalam repositori GitHub Anda. File ini akan berisi konfigurasi untuk alur kerja CI/CD.

Contoh:

```
name: CI/CD Pipeline

on:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: '1.17'

      - name: Build
        run: go build -v ./...

      - name: Test
        run: go test -v ./...

      - name: Deploy
        run: |
          # Your deployment script or commands here
```

**2. Konfigurasi Workflow:** Konfigurasi workflow harus mengatur kapan alur kerja ini harus dijalankan dan langkah-langkah apa yang harus dilakukan. Dalam contoh di atas, workflow akan dijalankan setiap kali ada push ke branch main

**3. Langkah-langkah Alur Kerja:** Setiap langkah dalam alur kerja didefinisikan sebagai steps. Setiap langkah bisa berisi berbagai perintah yang perlu dijalankan, seperti mengambil kode, mengatur lingkungan, membangun, menguji, dan menyebarluaskan aplikasi.

**4. Menambahkan dan Melakukan Perubahan:** Setelah membuat file konfigurasi, simpan dan dorong perubahan ke repositori GitHub Anda. GitHub Actions akan secara otomatis mendeteksi perubahan ini dan menjalankan alur kerja sesuai dengan konfigurasi yang telah Anda tentukan.

**5. Memantau Hasil:** Anda dapat melihat hasil alur kerja CI/CD di tab "Actions" di repositori GitHub Anda. Jika ada kesalahan, Anda dapat melihat log untuk menemukan penyebabnya.

# Thank You :star2:

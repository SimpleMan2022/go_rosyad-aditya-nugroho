# Clean and Hexagonal Architecture :rocket:

## Apa itu Clean Architecture

Clean Architecture adalah sebuah pendekatan dalam pengembangan perangkat lunak yang menekankan pada pemisahan konsep-konsep inti aplikasi dari detail-detail implementasinya. Clean Architecture membagi aplikasi menjadi beberapa lapisan yang saling terisolasi dan independen satu sama lain.

## Mengapa Clean Architecture

Clean Architecture membantu memisahkan konsep bisnis dari implementasi teknisnya, sehingga membuat aplikasi lebih mudah untuk dimodifikasi dan diuji. Dengan Clean Architecture, aplikasi juga lebih modular dan lebih mudah untuk dikelola.

## Perbedaan dengan MVC

Clean Architecture berbeda dengan Model-View-Controller (MVC) dalam hal struktur dan filosofi. MVC lebih fokus pada pemisahan tugas antara model, tampilan, dan kontrol, sedangkan Clean Architecture lebih fokus pada pemisahan konsep bisnis dari detail teknis implementasinya.

## Kendala Sebelum Menerapkan Clean Architecture

1. Pemisahan Konsep Bisnis dan Detail Implementasi:

- MVC: MVC membagi aplikasi menjadi tiga komponen utama: Model (representasi data), View (tampilan dari data), dan Controller (pengontrol alur aplikasi). Namun, MVC tidak memiliki aturan yang ketat tentang bagaimana konsep bisnis harus dipisahkan dari detail implementasinya.

- Clean Architecture: Clean Architecture menempatkan penekanan yang kuat pada pemisahan konsep bisnis dari detail implementasi. Dalam Clean Architecture, konsep bisnis diisolasi dalam lapisan yang independen dari detail teknis seperti database, antarmuka pengguna, atau bahkan framework yang digunakan.

2. Fleksibilitas dan Modifikasi:

- MVC: MVC dapat menjadi kurang fleksibel ketika aplikasi tumbuh karena tumpang tindih antara tugas-tugas yang dijalankan oleh Model, View, dan Controller.

- Clean Architecture: Clean Architecture membuat aplikasi lebih mudah untuk dimodifikasi karena lapisan-lapisannya saling terisolasi. Perubahan pada konsep bisnis tidak harus merubah lapisan lainnya, sehingga meminimalkan dampak perubahan.

3. Pengujian:

- MVC: Pengujian dalam MVC dapat menjadi sulit karena tugas-tugas yang berbeda terkadang bercampur aduk dalam komponen-komponen MVC.

- Clean Architecture: Clean Architecture mendorong pengujian yang lebih baik karena logika bisnis dapat diuji tanpa harus bergantung pada detail implementasinya.

4. Ketergantungan pada Framework:

- MVC: MVC sering kali tergantung pada framework tertentu, karena framework tersebut menentukan bagaimana aplikasi dibangun dan dijalankan.

- Clean Architecture: Clean Architecture berusaha untuk menjadi independen dari framework, sehingga memungkinkan untuk mengganti framework tanpa harus merubah keseluruhan aplikasi

## Keuntungan Clean Architecture

Beberapa keuntungan dari Clean Architecture antara lain:

- Kode Lebih Bersih: Dengan pemisahan yang jelas antara inti aplikasi dan detail implementasi, kode menjadi lebih bersih dan mudah dipahami.

- Mudah Dikelola: Karena struktur yang jelas,
  aplikasi yang menggunakan Clean Architecture menjadi lebih mudah untuk dikelola dan dikembangkan.

- Pengujian yang Lebih Baik: Kode yang terpisah dari detail implementasi memungkinkan pengujian yang lebih baik dan lebih mudah dilakukan.

## Clean Architecture Layer

- Clean Architecture terdiri dari beberapa layer yang masing-masing memiliki tanggung jawabnya sendiri dalam menjaga keterpisahan antara inti aplikasi dengan detail implementasi. Berikut adalah penjelasan lebih detail tentang setiap layer:

- Entities: Layer ini berisi definisi dari objek-objek yang digunakan dalam aplikasi. Entities merepresentasikan konsep bisnis dalam aplikasi dan tidak boleh tergantung pada detail implementasi seperti database atau user interface.

- Usecase: Layer ini berisi logika bisnis dari aplikasi. Usecase adalah tempat di mana aturan bisnis diterapkan dan operasi-operasi yang kompleks dilakukan. Usecase menggunakan Entities untuk mendapatkan dan menyimpan data, tetapi tidak boleh bergantung pada detail implementasi dari layer Repository.

- Repository: Layer ini berfungsi sebagai antarmuka untuk berkomunikasi dengan penyimpanan data, seperti database atau file sistem. Repository menyediakan metode untuk menyimpan dan mengambil data tanpa perlu mengetahui detail implementasi dari penyimpanan data tersebut.

- Driver: Layer ini berisi implementasi detail dari Repository. Driver bertanggung jawab untuk menghubungkan aplikasi dengan penyimpanan data yang sebenarnya, seperti database SQL atau API eksternal. Driver harus dibuat sedemikian rupa sehingga dapat diganti dengan driver lain tanpa memengaruhi layer Usecase dan Entities.

- Controller: Layer ini berisi implementasi detail dari user interface aplikasi. Controller bertugas untuk menerima input dari pengguna, memprosesnya menggunakan Usecase, dan menampilkan output kepada pengguna. Controller tidak boleh berisi logika bisnis atau akses langsung ke penyimpanan data.

## Komunikasi Antar Domain

Komunikasi antar domain dalam Clean Architecture dilakukan melalui penggunaan antarmuka. Antarmuka ini berfungsi sebagai semacam jembatan yang menghubungkan satu domain dengan domain lainnya. Dengan menggunakan antarmuka, setiap domain dapat mendefinisikan cara interaksi dengan domain lain tanpa perlu mengetahui detail implementasi dari domain tersebut. Hal ini memungkinkan setiap domain untuk tetap independen dan fokus pada tugasnya masing-masing, sehingga mempermudah perubahan dan perawatan aplikasi secara keseluruhan.

# Thank You :star2:

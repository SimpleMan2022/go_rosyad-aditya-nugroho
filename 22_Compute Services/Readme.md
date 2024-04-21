# Compute Services :rocket:

## System & Software Deployment

Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen.

## Strategy Deployment

1. Big Bang/Replace Deployment Strategy:

- Deskripsi: Versi baru perangkat lunak dideploy sekaligus, menggantikan versi lama.
- Keuntungan: Sederhana dan cepat, karena deployment dilakukan sekaligus.
- Kerugian: Risiko tinggi, karena masalah dengan versi baru akan mempengaruhi seluruh sistem sekaligus. Rollback bisa sulit dilakukan.

2. Rollout Deployment Strategy:

- Deskripsi: Versi baru dideploy secara bertahap ke subset pengguna atau server sebelum dideploy sepenuhnya.
- Keuntungan: Memungkinkan pengujian dan validasi secara bertahap, mengurangi risiko masalah yang luas. Rollback lebih mudah jika ditemukan masalah.
- Kerugian: Deployment bisa memakan waktu lebih lama, dan mengkoordinasikan beberapa versi bisa kompleks.

3. Blue/Green Deployment Strategy:

- Deskripsi: Dua lingkungan produksi yang identik (biru dan hijau) dipertahankan. Satu lingkungan (misalnya, biru) melayani lalu lintas langsung sementara lingkungan lainnya (hijau) diperbarui dengan versi baru. Setelah pembaruan berhasil, lalu lintas dialihkan ke lingkungan hijau.
- Keuntungan: Downtime dan risiko minimal, karena versi lama masih tersedia sampai versi baru benar-benar diuji.
- Kerugian: Membutuhkan sumber daya tambahan untuk mempertahankan dua lingkungan yang identik.

4. Canary Deployment Strategy:

- Deskripsi: Versi baru dideploy ke subset kecil pengguna atau server terlebih dahulu (kelompok canary) untuk memvalidasi kinerjanya dan stabilitasnya. Jika versi baru berperforma baik, maka akan secara bertahap diperluas ke lebih banyak pengguna atau server.
- Keuntungan: Memungkinkan umpan balik dan pengujian awal dalam skala kecil, mengurangi risiko masalah yang luas.
- Kerugian: Membutuhkan pemantauan dan otomatisasi yang hati-hati untuk mengelola proses rollout.

## AWS EC2 & RDS

AWS EC2 (Elastic Compute Cloud) dan RDS (Relational Database Service) adalah dua layanan yang sangat populer dari Amazon Web Services (AWS) yang digunakan untuk menyediakan infrastruktur komputasi dan database dalam cloud. Berikut adalah penjelasan singkat tentang keduanya:

**AWS ECS2** adalah layanan komputasi berbasis cloud yang memungkinkan pengguna untuk menyewa server virtual (instances) dan menjalankan aplikasi mereka di atasnya.

Fitur utama:

- Skalabilitas: Pengguna dapat dengan mudah menambah atau mengurangi kapasitas server sesuai kebutuhan.
- Fleksibilitas: Tersedia berbagai tipe instance dengan spesifikasi yang berbeda-beda, seperti CPU, RAM, dan penyimpanan.
- Keamanan: EC2 menyediakan berbagai fitur keamanan, termasuk kontrol akses, enkripsi data, dan pemantauan keamanan.
- Manajemen: AWS menyediakan berbagai alat dan layanan untuk mengelola instance EC2, seperti Auto Scaling untuk otomatisasi skalabilitas, dan Elastic Load Balancing untuk mendistribusikan lalu lintas aplikasi.

**AWS RDS(Relational Database Service)** adalah layanan manajemen database berbasis cloud yang menyediakan akses mudah untuk mengoperasikan dan menskalakan database relasional.

Fitur utama:

- Mendukung Berbagai Database: RDS mendukung beberapa jenis database relasional populer, termasuk MySQL, PostgreSQL, Oracle, dan SQL Server.
- Manajemen Otomatis: AWS mengelola tugas-tugas administrasi database seperti pencadangan, pembaruan, dan pemantauan kinerja secara otomatis.
- Skalabilitas: RDS memungkinkan pengguna untuk dengan mudah menyesuaikan kapasitas database sesuai dengan kebutuhan, tanpa mengganggu ketersediaan.
- Keamanan: RDS menyediakan fitur keamanan seperti enkripsi data, firewall, dan pengelolaan akses yang terkontrol.

## Deployment Golang ke ECS2

Untuk mendeploy aplikasi Golang menggunakan AWS EC2, Anda dapat mengikuti langkah-langkah berikut:

**1. Buat Instance EC2:**

- Masuk ke AWS Management Console dan buka layanan EC2.
- Klik tombol "Launch Instance" untuk membuat instance baru.
- Pilih AMI (Amazon Machine Image) yang sesuai dengan kebutuhan Anda (biasanya AMI Linux atau Ubuntu).
- Pilih tipe instance yang diinginkan (misalnya, t2.micro untuk penggunaan percobaan).
- Konfigurasi detail instance seperti jaringan, storage, dan keamanan.
  Buat atau gunakan kunci SSH untuk mengakses instance.

**2. Persiapkan Instance:**

- Akses instance EC2 menggunakan koneksi SSH.
- Instal Golang di instance. Anda bisa mengunduh dan menginstal Golang dari situs resmi atau menggunakan manajer paket yang tersedia di distribusi Linux yang Anda gunakan.
- Setel GOPATH (lokasi di mana paket Go akan diinstal) dan PATH (menambahkan PATH Go binary ke PATH sistem).

**3. Upload dan Compile Aplikasi:**

- Unggah kode sumber aplikasi Golang Anda ke instance EC2 menggunakan SCP atau SFTP.
- Kompilasi aplikasi menggunakan perintah go build atau go install.

**4. Jalankan Aplikasi:**

- Jalankan aplikasi dengan mengeksekusi file biner hasil kompilasi.
- Pastikan aplikasi berjalan dengan benar dan bisa diakses dari luar jika diperlukan.

**5. Konfigurasi Firewall dan Keamanan:**

- Pastikan firewall instance EC2 Anda mengizinkan koneksi ke port yang digunakan oleh aplikasi Anda.
- Gunakan grup keamanan EC2 untuk mengontrol akses ke instance Anda.

**6. Monitoring dan Manajemen:**

- Gunakan layanan AWS CloudWatch untuk memantau kesehatan dan kinerja instance EC2 Anda.
- Gunakan alat manajemen seperti AWS Systems Manager atau SSH untuk mengelola instance EC2 Anda.

# Thank You :star2:

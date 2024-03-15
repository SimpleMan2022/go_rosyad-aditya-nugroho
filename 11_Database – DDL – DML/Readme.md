# Database – DDL – DML :rocket:

Database adalah kumpulan data yang terorganisir dan saling terkait. Data dalam database disimpan dalam tabel yang terdiri dari baris dan kolom. Setiap baris mewakili satu record data, dan setiap kolom mewakili satu atribut data.

## Database Relationship?

> Relationship dalam database mengacu pada hubungan antara entitas atau tabel dalam database. Ada beberapa jenis relasi dalam database, termasuk one-to-one, one-to-many, dan many-to-many.

1. one to one (1:1): Satu record di tabel pertama hanya dapat terkait dengan satu record di tabel kedua.
2. one to many (1:N): Satu record di tabel pertama dapat terkait dengan banyak record di tabel kedua.
3. many to many (N:N): Banyak record di tabel pertama dapat terkait dengan banyak record di tabel kedua.

## RDBMS?

> RDBMS adalah singkatan dari Relational Database Management System. RDBMS adalah sistem manajemen database yang menggunakan model database relasional. RDBMS populer termasuk MySQL, PostgreSQL, dan Oracle.

## Jenis Jenis Perintah SQL

1. DDL (Data Definition Language) digunakan untuk mendefinisikan struktur database. Contoh perintah DDL adalah CREATE, ALTER, DROP.
2. DML (Data Manipulation Language) digunakan untuk mengelola data dalam database. Contoh perintah DML adalah INSERT, UPDATE, DELETE.

3. DCL (Data Control Language) digunakan untuk mengatur hak akses pengguna terhadap database. Contoh perintah DCL adalah GRANT dan REVOKE.

## Tipe Data MySQL

> MySQL memiliki berbagai tipe data yang dapat digunakan untuk menyimpan nilai berbeda sesuai dengan jenis data yang akan disimpan. Berikut adalah beberapa tipe data yang umum digunakan dalam MySQL beserta penjelasannya:

- INTEGER: Tipe data ini digunakan untuk menyimpan bilangan bulat. Ada beberapa sub-tipe INTEGER seperti TINYINT, SMALLINT, MEDIUMINT, INT, dan BIGINT, yang masing-masing memiliki rentang nilai yang berbeda.

- FLOAT dan DOUBLE: Tipe data ini digunakan untuk menyimpan nilai pecahan (bilangan desimal). FLOAT menyimpan nilai dengan presisi tunggal, sedangkan DOUBLE menyimpan nilai dengan presisi ganda.

- DECIMAL: Tipe data ini juga digunakan untuk menyimpan nilai pecahan, namun dengan presisi yang lebih baik daripada FLOAT atau DOUBLE. DECIMAL cocok untuk kebutuhan keuangan atau perhitungan yang membutuhkan presisi yang tinggi.

- CHAR dan VARCHAR: Tipe data ini digunakan untuk menyimpan string teks. CHAR digunakan untuk string dengan panjang tetap, sedangkan VARCHAR digunakan untuk string dengan panjang variabel.

- DATE, TIME, dan DATETIME: Tipe data ini digunakan untuk menyimpan nilai tanggal dan waktu. DATE digunakan untuk tanggal, TIME digunakan untuk waktu, dan DATETIME digunakan untuk kombinasi tanggal dan waktu.

- ENUM: Tipe data ini digunakan untuk membuat kolom yang nilainya terbatas pada daftar nilai yang telah ditentukan sebelumnya. ENUM memungkinkan Anda untuk memilih satu dari beberapa nilai yang telah ditentukan.

- SET: Tipe data ini juga digunakan untuk membuat kolom yang nilainya terbatas pada daftar nilai yang telah ditentukan sebelumnya. Namun, SET memungkinkan Anda untuk memilih beberapa nilai dari daftar nilai yang telah ditentukan.

- BLOB dan TEXT: Tipe data ini digunakan untuk menyimpan data biner (BLOB) atau teks panjang (TEXT). BLOB cocok untuk menyimpan gambar, file, atau data biner lainnya, sedangkan TEXT cocok untuk menyimpan teks panjang.

## Select, Insert, Update, dan Delete

**Select** digunakan untuk mengambil/menampilkan data dari database.

```
SELECT * FROM customers;
```

**Insert** digunakan untuk menambahkan data baru ke dalam database.

```
INSERT INTO customers (name, email) VALUES ('John Doe', 'john.doe@example.com');
```

**Update** digunakan untuk merubah data yang sebelumnya sudah dibuat dari database.

```
UPDATE customers SET email = 'new.email@example.com' WHERE id = 1;
```

**Delete** digunakan untuk menghapus data dari database.

```
DELETE FROM customers WHERE id = 1;
```

## Like, Between, And, Or, Order By, dan Limit.

**Like** digunakan untuk mencari data yang mirip dengan pola tertentu.

```
SELECT * FROM customers WHERE name LIKE "%Doe%";
```

**Between** digunakan untuk mencari data yang berada di antara dua nilai.

```
SELECT * FROM customers WHERE age BETWEEN 18 AND 35;
```

**And dan Or** digunakan untuk menggabungkan dua kondisi.

```
SELECT * FROM customers WHERE name LIKE "%Doe%" AND age BETWEEN 18 AND 35;
```

**Order By** digunakan untuk mengurutkan data.

```
SELECT * FROM customers ORDER BY name ASC;
```

**Limit** digunakan untuk membatasi jumlah data yang diambil.

```
SELECT * FROM customers LIMIT 10;

```

# Thank You :star2:

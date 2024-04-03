# Middleware :rocket:

## Apa itu Middleware

Middleware adalah perangkat lunak yang berperan sebagai perantara antara aplikasi web dan server. Middleware digunakan untuk memodifikasi atau menambah fungsionalitas pada request dan response HTTP sebelum atau sesudah request diproses oleh handler. Middleware dapat digunakan untuk berbagai tujuan seperti logging, authentication, authorization, caching, dan lain-lain.

## Alur Implementasi Middleware

Implementasi middleware dalam framework Echo dilakukan dengan menambahkannya pada router menggunakan method Pre() atau Use(). Method Pre() digunakan untuk middleware yang akan dijalankan sebelum route handler, sedangkan Use() untuk middleware yang akan dijalankan sebelum dan/atau sesudah route handler.

Contoh penggunaan middleware dalam Echo:

```
e := echo.New()

// Middleware menggunakan Pre()
e.Pre(middleware.Logger())
e.Pre(middleware.Recover())

// Middleware menggunakan Use()
e.Use(middleware.JWT([]byte("secret")))
e.Use(middleware.Gzip())
```

## Contoh Third-Party Middleware

**Negroni**
Negroni adalah salah satu middleware yang populer digunakan dengan framework Go. Negroni menyediakan cara yang mudah untuk menambahkan middleware pada aplikasi Go.

**Interpose**
Interpose adalah middleware chaining library yang memungkinkan pengguna untuk menambahkan middleware dalam urutan tertentu.

**Alice**
Alice adalah middleware chaining library yang memungkinkan pengguna untuk menambahkan middleware dengan urutan tertentu.

## Contoh Middleware di Echo

**Echo.Pre()** middleware:

- httpsredirect: Mengarahkan request HTTP ke HTTPS.
- addtralingslash: Menambahkan trailing slash pada URL jika tidak ada.
- methodOverride: Mengizinkan penggunaan method HTTP seperti PUT atau DELETE melalui query parameter atau form value.
- rewrite: Menangani rewrite URL.

**Echo.Use()** middleware:

- bodylimit: Mengatur batas ukuran body request.
- logger: Logging request.
- gzip: Mengompres response menggunakan gzip.
- recover: Menangani panic dan recover.
- jwtauth: Middleware untuk authentication menggunakan JWT.
- cors: Middleware untuk mengatur Cross-Origin Resource Sharing (CORS).

## Apa itu CORS dan bagaimana Konfigurasinya

CORS (Cross-Origin Resource Sharing) adalah mekanisme yang memungkinkan server untuk memberikan izin kepada aplikasi web di domain yang berbeda untuk mengakses resource yang terdapat pada server tersebut. Secara default, browser melarang request lintas domain (cross-origin) untuk mencegah potensi ancaman keamanan seperti serangan CSRF (Cross-Site Request Forgery).

Untuk mengaktifkan CORS, server harus merespons dengan header HTTP yang tepat, seperti Access-Control-Allow-Origin untuk menentukan domain yang diizinkan, Access-Control-Allow-Methods untuk menentukan metode HTTP yang diizinkan, dan Access-Control-Allow-Headers untuk menentukan header tambahan yang diizinkan.

**Middleware CORS**

```
e := echo.New()
e.Use(middleware.CORS())
```

**Middleware CORS dengan Konfigurasi**

```
e := echo.New()
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"}, // Izinkan semua domain
    AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE}, // Izinkan metode HTTP tertentu
}))

```

## Rate Limiter

Rate limiter adalah mekanisme untuk mengontrol jumlah request yang diterima oleh server dalam periode waktu tertentu dari satu pengguna atau IP address. Hal ini berguna untuk mencegah serangan DDoS (Distributed Denial of Service) dan memastikan ketersediaan server.

Implementasi rate limiter dalam Echo dapat dilakukan dengan menggunakan middleware kustom atau library seperti github.com/juju/ratelimit. Dengan middleware ini, Anda dapat menentukan jumlah maksimum request yang diperbolehkan dalam satu periode waktu dan menangani request yang melebihi batas tersebut dengan memberikan respons yang sesuai, seperti HTTP status code 429 (Too Many Requests).

**Rate Limiter Echo**

```
e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
```

Dengan menggunakan rate limiter, Anda dapat mengendalikan lalu lintas request yang masuk ke server Anda, menjaga ketersediaan server, dan meningkatkan keamanan aplikasi web Anda.

## Log Middleware

Middleware logging adalah mekanisme untuk mencatat informasi penting seperti request method, path, status code, dan waktu respons dari server. Middleware logging sangat berguna untuk debugging, analisis performa, dan pemantauan aplikasi.

Dalam Echo, logging dapat diimplementasikan dengan menggunakan middleware Logger() yang disediakan oleh Echo. Middleware ini secara otomatis mencatat informasi request dan response ke output standar atau file log.

Contoh penggunaan logging middleware dalam Echo:

```
e.Use(middleware.Logger())
```

## JWT Middeware

JWT (JSON Web Token) middleware digunakan untuk melakukan authentication menggunakan token JWT. JWT adalah standar terbuka (RFC 7519) yang mendefinisikan format token yang aman untuk ditransmisikan antar pihak.

Middleware JWT pada dasarnya memeriksa apakah token yang dikirim oleh pengguna valid dan belum kadaluarsa. Jika token valid, pengguna dianggap telah terautentikasi dan dapat diizinkan untuk mengakses resource yang dilindungi.

Dengan menggunakan JWT middleware, Anda dapat dengan mudah menambahkan lapisan keamanan tambahan pada aplikasi web Anda, memastikan bahwa hanya pengguna yang memiliki token yang valid yang dapat mengakses resource yang dilindungi.

Jadi, auth middleware dan JWT middleware adalah dua konsep penting dalam mengamankan aplikasi web Anda dan memastikan bahwa hanya pengguna yang sah yang dapat mengakses resource yang dilindungi.

# Thank You :star2:

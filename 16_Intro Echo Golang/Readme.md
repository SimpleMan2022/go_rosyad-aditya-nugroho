# Intro Echo Golang :rocket:

## Apa itu Echo Golang

Melalui web resminya, Proyek Echo adalah kerangka kerja web yang kuat dan serbaguna untuk membangun aplikasi web yang dapat diskalakan dan berkinerja tinggi dalam bahasa pemrograman Go. Proyek ini mengikuti prinsip-prinsip kesederhanaan, fleksibilitas, dan kinerja untuk menyediakan perangkat yang efisien bagi para pengembang untuk membangun aplikasi web yang kuat.

## Keuntungan Menggunakan Echo

1. **Cepat dan Ringan**:Echo dirancang untuk kecepatan dan efisiensi, memastikan overhead minimal dan kinerja tinggi untuk menangani permintaan dan respons HTTP.
2. **Routing**: Framework ini menawarkan sistem perutean yang fleksibel dan intuitif yang memungkinkan pengembang untuk menentukan rute dengan parameter, string kueri, dan penangan khusus.
3. **Dukungan Middleware**: Echo menyediakan dukungan middleware yang luas, memungkinkan pengembang untuk dengan mudah mengimplementasikan masalah lintas sektoral seperti pencatatan, otentikasi, penanganan kesalahan, dan banyak lagi.
4. **Penanganan Permintaan Berbasis Konteks**: Dengan penanganan permintaan berbasis konteks, Echo menawarkan akses mudah ke data dan parameter khusus permintaan, menyederhanakan pengembangan aplikasi web.
5. **Perenderan Templat yang Kuat**: Echo menyertakan mesin rendering template yang kuat yang mendukung berbagai bahasa template, sehingga pengembang dapat membuat konten HTML dinamis dengan mudah.
6. **Validasi dan Pengikatan**: Framework ini menyediakan kemampuan validasi dan pengikatan data yang kuat, membuatnya mudah untuk memvalidasi data permintaan yang masuk dan mengikatnya ke struktur Go.
7. **Ekstensibilitas**: Echo sangat dapat diperluas, dengan dukungan untuk middleware khusus, mesin template, dan komponen lainnya, memungkinkan pengembang untuk menyesuaikan kerangka kerja dengan kebutuhan spesifik mereka.
8. **Komunitas dan Ekosistem**: Proyek Echo mendapat manfaat dari komunitas yang dinamis dan aktif yang menyumbangkan pustaka, plugin, dan ekstensi, menumbuhkan ekosistem komponen yang dapat digunakan kembali.

## REST API

REST API (Representational State Transfer API) adalah jenis API yang mengikuti seperangkat aturan yang membuatnya mudah digunakan dan dipahami. REST API menggunakan metode HTTP standar (GET, POST, PUT, DELETE) untuk mengakses data.

**Karakteristik REST API:**

- Sumber daya: Data yang diakses melalui REST API disebut sumber daya. Setiap sumber daya memiliki alamat URL yang unik.
- Metode HTTP: REST API menggunakan metode HTTP standar (GET, POST, PUT, DELETE) untuk mengoperasikan sumber daya.
  Format data: REST API biasanya menggunakan format data JSON atau XML untuk menukar data.

**Keuntungan REST API:**

- Mudah digunakan: REST API mudah dipahami dan digunakan oleh pengembang.
- Skalabilitas: REST API dapat dengan mudah diskalakan untuk mendukung banyak pengguna dan aplikasi.
- Ketersediaan: REST API tersedia di berbagai platform dan bahasa pemrograman.

## Basic Routing dan Controller pada Echo

Routing pada Echo sangatlah mudah. Berikut adalah contoh penggunaan basic routing dan controller pada Echo:

```
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Basic routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Controller
	e.GET("/users/:id", getUserByID)

	e.Start(":8080")
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "User ID: "+id)
}
```

## Data Rendering pada Echo

Echo menyediakan berbagai metode untuk merender data, seperti JSON, HTML, XML, dan sebagainya. Berikut adalah contoh penggunaan rendering JSON pada Echo:

```
func main() {
	e := echo.New()

	e.GET("/users/:id", getUserJSONByID)

	e.Start(":8080")
}

func getUserJSONByID(c echo.Context) error {
	id := c.Param("id")
	user := User{ID: id, Name: "John Doe"}

  // JSON rendering
	return c.JSON(http.StatusOK, user)
}

```

## HTTP Status Code

Server API mengembalikan kode status HTTP untuk menunjukkan status permintaan. Berikut adalah beberapa kode status HTTP yang umum digunakan:

- 200 OK: Permintaan berhasil.
- 400 Bad Request: Permintaan tidak valid.
- 401 Unauthorized: Akses tidak diizinkan.
- 404 Not Found: Data yang diminta tidak ditemukan.
- 500 Internal Server Error: Terjadi kesalahan pada server.

## Retrieve Data pada Echo (URL Param, Query Param, Form Value)

Untuk mendapatkan data dari URL parameter, query parameter, atau form value, kita dapat menggunakan method-method yang disediakan oleh Context. Berikut adalah contoh penggunaannya:

```
func main() {
	e := echo.New()

	// Retrieve data
	e.GET("/users/:id", getUserData)

	e.Start(":8080")
}

func getUserData(c echo.Context) error {
	// Retrieve URL parameter
	id := c.Param("id")

	// Retrieve query parameter
	queryParam := c.QueryParam("query")

	// Retrieve form value
	formValue := c.FormValue("form")

	return c.String(http.StatusOK, "User ID: "+id+", Query Param: "+queryParam+", Form Value: "+formValue)
}
```

## Binding Data JSON

Untuk mengikat data JSON yang dikirimkan oleh client, kita dapat menggunakan struct di Go. Berikut adalah contoh penggunaannya:

```
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()


	e.POST("/users", createUser)

	e.Start(":8080")
}

func createUser(c echo.Context) error {
	user := new(User)
  // Binding JSON
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}


```

# Thank You :star2:

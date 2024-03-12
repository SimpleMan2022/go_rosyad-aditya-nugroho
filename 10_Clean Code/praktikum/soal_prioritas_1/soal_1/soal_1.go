package main

import "fmt"

// kesalahan penamaan variabel dimana sebaiknya penamaan user dimulai dengan huruf besar.
type user struct {
	Email    string
	Password string
}

// kesalahan penamaan variabel dimana sebaiknya penamaan userRepo dimulai dengan huruf besar.
type userRepo struct {
	DB []user
}

// gunakan nama yang lebih spesifik untuk parameter user, ganti u menjadi user atau usr
// penggunaan pointer untuk penunjuk ke UserRepo
func (r userRepo) Register(u user) {

	// pemisahan fungsi untuk melakukan validasi
	if u.Email == "" || u.Password == "" {
		// lakukan handling ketika error daripada menampilkan pesan ke konsol saja
		fmt.Println("register failed")
	}

	r.DB = append(r.DB, u)
}

// gunakan nama yang lebih spesifik untuk parameter user, ganti u menjadi user atau usr
// penggunaan pointer untuk penunjuk ke UserRepo
func (r userRepo) Login(u user) string {

	// pemisahan fungsi untuk melakukan validasi
	if u.Email == "" || u.Password == "" {
		// lakukan handling ketika error daripada menampilkan pesan ke konsol saja
		fmt.Println("login failed")
	}

	for _, us := range r.DB {
		if us.Email == u.Email && us.Password == u.Password {
			return "auth token"
		}
	}

	return ""
}

// dibawah ini adalah contoh yang benar

type User struct {
	Email    string
	Password string
}

type UserRepo struct {
	DB []User
}

func (r *UserRepo) Register(user *User) error {
	if validUser := isValidUser(user); !validUser {
		return fmt.Errorf("Register Failed")
	}
	r.DB = append(r.DB, *user)
	return nil
}

func (r *UserRepo) Login(user *User) (string, error) {
	if validUser := isValidUser(user); !validUser {
		return "", fmt.Errorf("Login Failed")
	}

	for _, us := range r.DB {
		if us.Email == user.Email && us.Password == user.Password {
			return "auth token", nil
		}
	}

	return "", fmt.Errorf("Login failed: wrong email or password")
}

func isValidUser(user *User) bool {
	return user.Email != "" && user.Password != ""
}

//Ada beberapa kekurangan dalam penulisan kode tersebut:
//
//1. Penamaan variabel yang tidak mengikuti konvensi, dimana sebaiknya dimulai dengan huruf besar
//(contoh: user menjadi User, userRepo menjadi UserRepo).

//2. Penanganan kesalahan yang tidak ideal, hanya mencetak pesan ke konsol tanpa tindakan lebih lanjut.
//3. Penggunaan parameter u yang kurang deskriptif, sebaiknya diganti menjadi user atau usr.
//4. Tidak menggunakan pointer untuk parameter dalam method Register dan Login.

//Kekurangan tersebut terjadi pada bagian:
//
//1. Penamaan variabel user dan userRepo.
//2. Penggunaan parameter u dalam method Register dan Login.
//3. Penanganan kesalahan dan validasi.

//Alasan dari setiap kekurangan tersebut:
//
//1. Penamaan variabel yang tidak mengikuti konvensi membuat kode sulit dibaca dan memahami maksudnya.
//2. Penggunaan parameter u yang kurang deskriptif membuat sulit untuk memahami tujuan dari parameter
//	 tersebut.
//3. Penanganan kesalahan yang hanya mencetak pesan ke konsol tanpa tindakan lebih lanjut tidak memberikan
//	 informasi yang cukup kepada pengguna atau pengembang untuk mengetahui apa yang salah dan bagaimana
//	 cara memperbaikinya.
//4. Tidak menggunakan pointer untuk parameter dalam method Register dan Login dapat menyebabkan perubahan
//	 tidak terlihat pada objek UserRepo yang digunakan dalam method tersebut. Menggunakan pointer akan
//	 memastikan bahwa perubahan yang dilakukan pada objek akan disimpan.

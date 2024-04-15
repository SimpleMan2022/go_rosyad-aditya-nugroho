# Docker :rocket:

## Apa itu Virtual Machine

VM adalah sebuah emulasi komputer yang berjalan di atas sistem operasi lain. VM memiliki sumber dayanya sendiri, seperti CPU, RAM, dan storage, yang terisolasi dari sistem operasi host. Hal ini memungkinkan pengguna untuk menjalankan sistem operasi dan aplikasi yang berbeda pada satu mesin fisik.

## Apa itu Container

Container adalah sebuah unit perangkat lunak yang mengemas aplikasi dan semua dependensinya, termasuk kode, runtime, sistem tools, dan library, menjadi satu paket yang terisolasi. Container tidak memiliki kernelnya sendiri, sehingga mereka berbagi kernel dengan sistem operasi host. Hal ini membuat container lebih ringan dan lebih cepat daripada VM.

## Apa itu Docker

Docker adalah sebuah platform open source untuk membangun, menjalankan, dan mengelola container. Docker menyediakan tools dan API untuk membuat, berbagi, dan menjalankan container di berbagai lingkungan, termasuk laptop, server, dan cloud.

## Bagaimana Docker Bekerja

Docker menggunakan teknologi bernama containerization untuk mengemas aplikasi dan dependensinya ke dalam container. Container terdiri dari beberapa layer, di mana setiap layer berisi perubahan pada image sebelumnya. Hal ini membuat container lebih ringan dan lebih mudah dibagikan.

## Docker Volume

Docker volume adalah direktori persisten yang dapat digunakan untuk menyimpan data aplikasi. Volume tidak terikat pada container dan dapat digunakan kembali oleh container yang berbeda.

## Docker Network

Docker network adalah jaringan virtual yang memungkinkan container untuk berkomunikasi satu sama lain. Docker network dapat digunakan untuk mengisolasi traffic antar container dan untuk membuat jaringan yang kompleks.

## Docker Compose

Docker Compose adalah sebuah tool yang memungkinkan pengguna untuk mendefinisikan dan menjalankan multi-container application dengan mudah. Docker Compose menggunakan file YAML untuk mendefinisikan layanan yang akan dijalankan dan konfigurasinya.

## Basic Sintaks Docker

**docker run:** digunakan untuk menjalankan image Docker dan membuat container baru.

- -d: menjalankan container di background
- -i: menempelkan terminal standar input (stdin) ke container
- -t: menempelkan terminal standar output (stdout) dan standar error (stderr) ke container
- --name: memberi nama pada container
- -p: memetakan port host ke port container (format: <host_port>:<container_port>)
- -v: me-mount volume ke direktori di dalam container (format: <volume_name>:<container_dir>)

- <image_name>:<tag>: nama image dan tagnya (jika ada)

**docker ps:** menampilkan daftar container yang sedang berjalan.

- -a: menampilkan semua container (termasuk yang berhenti)

**docker stop:** menghentikan container yang sedang berjalan.

- <container_name>: nama container
- docker start: memulai container yang telah dihentikan.

- <container_name>: nama container

**docker exec:** menjalankan perintah di dalam container yang sedang berjalan.

- -it: menempelkan terminal standar input (stdin) ke container
- <container_name>: nama container
- <command>: perintah yang akan dijalankan di dalam container

**docker logs:** menampilkan log dari container yang sedang berjalan.

- <container_name>: nama container
- -f: mengikuti log secara real-time
  docker rm: menghapus container yang telah dihentikan.

- <container_name>: nama container

**docker build:** membangun image Docker dari Dockerfile.

- . : membangun image dari Dockerfile di direktori kerja saat ini

**docker images:** menampilkan daftar image Docker yang tersedia di sistem.

**docker search:** mencari image Docker di Docker Hub.

- <search_term>: istilah pencarian
  docker pull: menarik image Docker dari registry (seperti Docker Hub)

- <image_name>:<tag>: nama image dan tagnya (jika ada)

**docker push:** mendorong image Docker ke registry (seperti Docker Hub)

- <image_name>:<tag>: nama image dan tagnya (jika ada)

# Thank You :star2:

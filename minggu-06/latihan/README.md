Di dalam folder ini terdapat empat file program Go, yakni:  
- `go-msyql.go`, adalah program Go untuk konek ke databse mySQL.
- `go-mongodb.go`, adalah program Go untuk konek ke database mongoDB.
- `gin-mysql.go`, adalah RESTful API dengan framework gin dan database mySQL.
- `gin-mongodb.go`, adalah RESTful API dengan framework gin dan database mongoDB.
  
Terdapat juga file lain yakni:  
- `database.sql`, yakni database yang sudah diekspor.
- `persons_tabel.sql`, yakni data tabel *persons* yang sudah diekspor.
- `tekn_cloud.presons.json`, yakni *collection persons* yang sudah diekspor.

## Cara Menggunakan File
Lakukan *clone* dengan perintah berikut:  
```
$ git clone --depth 1 --branch minggu-06 --no-checkout https://github.com/dhanyg/tekn-cloud-computing.git
$ cd tekn-cloud-computing
$ git sparse-checkout set minggu-06/latihan
$ git checkout minggu-06
$ cd minggu
$ cd latihan
```

### File go-mysql
Pastikan Anda sudah menginstal mysql server, kemudian import file `database.sql`.  

Untuk menjalankan file `go-mysql.go`, ubah pada bagian kode berikut:  
```
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/tekn_cloud")
```
dan ganti dengan user dan password yang Anda gunakan.  
  
Jalankan file dengan perintah:  
```
go get .
go run go-mysql.go
``` 

## File go-mongodb
Pastikan Anda sudah menginstal mongodb server dan sudah berjalan. Pada latihan ini saya menggunakan MongoDB Atlas sebagai servernya.  
 
Kemudian buat database baru dengan nama `tekn_cloud`, lalu buat *collection* dengan nama `persons`. Kemudian impor file `tekn_cloud.persons.json`.  

Untuk menjalankan file `go-mongodb.go`, ubah pada bagian kode berikut:  
```
uri := "mongodb+srv://username:password@cluster0.vrbzjyd.mongodb.net/?retryWrites=true&w=majority"
```
ganti *__uri__* sesuai dengan *uri* yang Anda gunakan. Di sini saya menggunakan *uri* dari MongoDB Atlas.    
  
Jalankan file dengan perintah:  
```
go get .
go run go-mongodb.go
``` 

## File gin-mysql
Ubah bagian kode berikut:  
```
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/tekn_cloud")
```
dan ganti dengan user dan password yang Anda gunakan.  
  
Jalankan file dengan perintah:  
```
go get .
go run gin-mysql.go
``` 
  
Lakukan pengetesan dengan menggunakan perintah:  
```
curl http://localhost:8080/persons
```
  
## File gin-mongodb
ubah pada bagian kode berikut:  
```
uri := "mongodb+srv://username:password@cluster0.vrbzjyd.mongodb.net/?retryWrites=true&w=majority"
```
ganti *__uri__* sesuai dengan *uri* yang Anda gunakan. Di sini saya menggunakan *uri* dari MongoDB Atlas.    
  
Jalankan file dengan perintah:  
```
go get .
go run gin-mongodb.go
```
  
Lakukan pengetesan dengan menggunakan perintah:  
```
curl http://localhost:8080/persons
```
  
Selesai.
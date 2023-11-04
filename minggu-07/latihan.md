# Latihan Docker
## 1. Menginstal Docker
Jalankan file instaler Docker yang sudah diunduh. Centang opsi *Add shortcut to desktop*.  
![proses instal docker](./01-add-shortcut.png)  
  
Proses *unpacking files* berlangsung. Tunggu hingga selesai.   
![proses unpacking files docker](./02-unpacking-files.png)  
  
Instalasi telah selesai.  
![instalasi docker selesai](./03-installation-succeeded.png)  
  
Jalankan Docker Desktop kemudian lanjutkan dengan *Sign in* apabila sudah memiliki akun docker atau pilih *Continue without signing in* jika belum.  
![welcome to docker desktop](./04-welcome-to-docker.png)  
  
## 2. Get Started
### 2.1 Containerize an application
Cloning repositori *getting-started-app*.  
```
$ git clone https://github.com/docker/getting-started-app.git$ git clone 
```
  
Masuk ke dalam direktori `getting-started-app` kemudian buat file `Dockerfile`.  
```
$ cd getting-started-app
$ vim Dockerfile
```
![isi file Dockerfile](./05-dockerfile.png)  
  
Build image dengan perintah:  
```
$ docker build -t getting-started .
```
![proses build image dengan docker](./06-docker-build.png)  
  
Jalankan container dengan perintah:  
```
$ docker run -dp 127.0.0.1:3000:3000 getting-started
```
  
Buka browser kemudian akses ke halaman 127.0.0.1:3000.  
![tampilan aplikasi todo app](./07-localhost-3000.png)  
  
Cek dengan perintah `docker ps` untuk melihat daftar container yang berjalan.  
![perintah docker ps](./08-docker-ps.png)  
  
### 2.2 Update the application
Lakukan modifikasi pada file `src/static/js/app.js`.  
```
- <p className="text-center">No items yet! Add one above!</p>
+ <p className="text-center">You have no todo items yet! Add one above!</p>
```
  
Hentikan container yang sedang berjalan.  
```
$ docker stop <ID container atau nama container>
```
  
Kemudian build ulang image dengan perintah:  
```
$ docker build -t getting-started .
```
  
Jalankan container dengan perintah:  
```
$ docker run -dp 127.0.0.1:3000:3000 getting-started
```
![menjalankan perintah docker run](./09-docker-run.png)  
  
Refresh halaman `127.0.0.1:3000`.  
![tampilan todo app setelah di-refresh](./10-localhost-3000.png)  
  
### 2.3 Share the application
Login ke [Docker Hub](https://hub.docker.com/), kemudian buat repositori baru dengan menekan tombol *__Create repository__*.  
  
Pada bagian *Repository Name* diisi dengan `getting-started` dan pastikan *Visibility*-nya `Public`, kemudian klik tombol *__Create__*.  
![membuat repositori baru di docker hub](./11-create-repository.png)  
  
Login ke docker hub melalui terminal dengan perintah berikut:  
```
$ docker login -u <username>
```
![login docker hub melalui terminal](./12-docker-login.png)  
  
Buat *tag* untuk image `getting-started` dengan perintah:  
```
$ docker tag getting-started <username>/getting-started
```
![membuat tag image dengan perintah docker tag](./13-docker-tag.png)  
  
Push image yang sudah diberi tag tadi ke docker hub dengan perintah:  
```
$ docker push <username>/getting-started
```
![push image ke docker hub melalui terminal](./14-docker-push.png)  
![tampilan menu tag pada repositori di docker hub](./15-docker-hub-tag.png)  
  
Akses situs [Play with Docker](https://labs.play-with-docker.com/), kemudian login dengan akun docker yang digunakan.  
  
Klik tombol *__ADD NEW INSTANCE__*, kemudian jalankan container dari image yang sudah di-*push* ke docker hub tadi.  
```
$ docker run -dp 0.0.0.0:3000:3000 <username>/getting-started
```
![menjalankan container di platform Play with Docker](./16-new-instance.png)  
  
Klik nomor port *__3000__* untuk mengakses halaman web.  
![tampilan todo app yang diakses melalui platform Play with Docker](./17-todo-app.png)  
  
### 2.4 Persist the DB
Buat sebuah *volume* dengan perintah:  
```
$ docker volume create <nama-volume>
```
![membuat volume dengan perintah docker volume create](./18-docker-volume-create.png)  
  
Hapus container sebelumnya dengan perintah `docker rm -f <id>`  
![menghapus container dengan perintah docker rm](./19-docker-rm.png)  
  
Jalankan kembali container dari image `getting-started` dengan menambahkan parameter `--mount`:  
```
$ docker run -dp 127.0.0.1:3000:3000 --mount type=volume,src=todo-db,target=/etc/todos getting-started
```
![menjalankan container dengan parameter --mount](./20-docker-run.png)  
  
Akses Todo App melalui browser kemudian tambahkan beberapa data.  
![menambahkan beberapa data pada aplikasi todo app](./21-todo-app.png)  
  
Hentikan dan hapus container yang sedang berjalan.  
![perintah menghentikan dan menghapus container melalui terminal](./22-docker-stop-rm.png)  
  
Jalankan container baru dengan menambahkan parameter `--mount` seperti pada langkah-langkah sebelumnya.  
![menjalankan container baru menggunakan parameter --mount](./23-docker-run.png)  
  
Cek browser kembali dan pastikan data yang sebelumnya sudah pernah diinputkan masih ada.  
![tampilan todo app pada container baru dengan volume yang sama](./24-todo-app.png)  
  
Untuk mengetahui lokasi dimana volume berada, gunakan perintah:  
```
$ docker volume inspect <nama-volume>
```
![hasil inspect volume todo-db dengan perintah docker volume inspect](./25-docker-volume-inspect.png)  
  
### 2.5 Use bind mounts
Jalankan container baru dengan perintah:  
```
$ docker run -it --mount type=bind,src="$(pwd)",target=/src ubuntu bash
```
Perintah di atas akan melakukan *mount* direktori `src` yang ada di dalam container ke direktori yang saat ini sedang aktif pada host.  
![menjalankan container baru dengan mount bind](./26-docker-mount-bind.png)  
  
Dari dalam container, masuk ke direktori `src` kemudian cek isi direktorinya. Setelah itu buat sebuah file baru dengan nama `myfile.txt`.  
![membuat file baru di dalam direktori src](./27-create-new-file.png)  
  
Keluar dari container dengan perintah `exit`, kemudian cek direktori yang sedang aktif saat ini apakah terdapat file yang sudah dibuat di dalam container tadi.  
![mengecek isi direktori saat ini](./28-ls-pwd.png)  
  
### 2.6 Multi-container apps
Pada bagian ini akan dijalankan dua buah container dimana container pertama adalah container untuk database server dan container kedua adalah todo app. Agar kedua container tersebut dapat saling terhubung, keduanya harus berada di dalam network yang sama. Docker memiliki kemampuan untuk membuat sebuah network yang nantinya akan digunakan oleh container.  
  
Buat sebuah network dengan perintah:  
```
$ docker network create <nama-network>
```
![membuat network dengan perintah docker network create](./29-docker-network-create.png)  
  
Jalankan sebuah container MySQL dan hubungkan container tersebut ke dalam network yang telah dibuat tadi. Berikut perintah lengkapnya:  
```
$ docker run -d \
    --network todo-app --network-alias mysql \
    -v todo-mysql-data:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=secret \
    -e MYSQL_DATABASE=todos \
    mysql:8.0
```
![menjalankan container mysql dengan beberapa parameter tambahan](./30-docker-run-mysql.png)  
  
Akses container MySQL menggunakan perintah `docker exec -it <ID Container> mysql -u root -p` dan gunakan password `secure` untuk login mysql. Kemudian cek isi database menggunakan perintah `SHOW DATABASES;`.  
![mengakses container mysql dan melihat isi databasenya](./31-docker-exec-mysql.png)  
  
Selanjutnya jalankan container untuk todo app. Pastikan sudah berada di dalam direktori `getting-started-app`, lalu ketikkan perintah berikut:  
```
$ docker run -dp 127.0.0.1:3000:3000 \
  -w /app -v "$(pwd):/app" \
  --network todo-app \
  -e MYSQL_HOST=mysql \
  -e MYSQL_USER=root \
  -e MYSQL_PASSWORD=secret \
  -e MYSQL_DB=todos \
  node:18-alpine \
  sh -c "yarn install && yarn run dev"
```
![menjalankan container untuk todo app dengan beberapa parameter tambahan](./32-docker-run-todo.png)  
  
Kita bisa melihat log dari container todo app menggunakan perintah:  
```
$ docker logs -f <ID Container>
```
![melihat log container todo app](./33-docker-logs.png)  
  
Buka todo app melalui browser kemudian tambahkan beberapa data.  
![menambahkan data pada todo app](./34-todo-app.png)  
  
Cek pada container MySQL untuk memastikan data yang diinput pada todo app sudah tersimpan ke database.  
```
$ docker exec -it <mysql-container-id> mysql -p todos
```
![mengakses container MySQL dan melihat isi tabel](./35-docker-exec-mysql.png)  
  
### 2.7 Use Docker Compose
Buat file `compose.yaml` di dalam direktori `getting-started-app`. Kemudian tambahkan konfigurasi berikut:  
```
services:
  app:
    image: node:18-alpine
    command: sh -c "yarn install && yarn run dev"
    ports:
      - 127.0.0.1:3000:3000
    working_dir: /app
    volumes:
      - ./:/app
    environment:
      MYSQL_HOST: mysql
      MYSQL_USER: root
      MYSQL_PASSWORD: secret
      MYSQL_DB: todos

  mysql:
    image: mysql:8.0
    volumes:
      - todo-mysql-data:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: todos

volumes:
  todo-mysql-data:
```
`app` dan `mysql` adalah nama container yang akan dibuat.  
![isi file compose.yaml](./36-compose-yaml.png)  
  
Hapus container yang sebelumnya pernah dibuat.  
![menghapus container yang pernah dibuat](./37-docker-rm.png)  
  
Jalankan *Docker Compose* dengan perintah:  
```
$ docker compose up -d
```
![menjalankan docker compose di terminal](./38-docker-compose-up.png)  
  
Buka kembali todo app pada browser (localhost:3000) untuk memastikan jika container sudah berjalan dan saling terhubung.  
  
Selesai.
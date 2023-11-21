# Praktikum Teknologi Cloud Minggu ke-9
## Docker For Beginners - Linux
### Task 0: Prerequisites
Clone repositori yang akan digunakan untuk latihan dengan perintah:  
```
$ git clone https://github.com/dockersamples/linux_tweet_app
```
![perintah git clone pada terminal](./01-clone-repository.png)  
  
### Task 1: Run some simple Docker containers
Terdapat beberapa cara untuk menjalankan container.  
Cara pertama yakni menjalankan *single task* pada container. Jalankan perintah di bawah ini:  
```
$ docker container run alpine hostname
```
Perintah di atas akan menjalankan sebuah container dari image yang bernama `alpine `dan kemudian memerintahkan container tersebut untuk mengeksekusi perintah `hostname`.  
![perintah 'docker run container alpine hostname' pada terminal](./02-docker-run-single-task.png)  
Pada gambar di atas, `e373611bf0e5` merupakan hostname dari container yang juga merupakan ID container.  
  
Cara kedua adalah menjalankan container secara interaktif. Jalankan perintah di bawah ini:  
```
$ docker container run --interactive --tty --rm ubuntu bash
```
Perintah di atas akan menjalankan container dan mengeksekusi perintah *bash* yang kemudian akan menampilkan shell terminal.  
![perintah 'docker container run --interactive --tty --rm ubuntu bash' pada terminal](./03-docker-run-interactive.png)  
  
Di dalam shell kita dapat mengeksekusi perintah-perintah linux seperti `ls`, `ps`, dan `cat`.  
![menjalankan perintah ls /, ps aux, dan cat /etc/issue pada terminal](./04-execute-command.png)  
  
Cara ketiga adalah menjalankan container pada background. Container akan tetap berjalan selama kita belum melakukan perintah untuk stop atau menghapus containernya. Jalankan perintah di bawah ini:  
```
$ docker container run \
 --detach \
 --name mydb \
 -e MYSQL_ROOT_PASSWORD=my-secret-pw \
 mysql:latest
```
Parameter `--detach` pada kode di atas memerintahkan container agar tetap berjalan pada background setelah container selesai dibuat.  
![perintah docker container run dengan parameter --detach](./05-docker-container-detach.png)  
  
Untuk menjalankan perintah di dalam container yang sedang berjalan, menggunakan perintah `docker exec`. Jalankan perintah di bawah ini:  
```
$ docker exec -it mydb \
 mysql --user=root --password=$MYSQL_ROOT_PASSWORD --version
```
![perintah docker exec](./06-docker-exec.png)  
  
### Task 2: Package and run a custom app using Docker
Masuk ke dalam direktori `linux_tweet_app`. Kemudian cek isi file `Dockerfile` dengan perintah *cat*.  
![perintah cd dan cat pada terminal](./07-enter-directory.png)  
  
Buat variabel `DOCKERID` di terminal dengan menggunakan perintah *export*.  
![perintah export DOCKERID di terminal](./08-export-dockerid.png)  
  
Jalankan perintah `docker image build` untuk membuat image dari direktori aplikasi saat ini (linux_tweet_app). Gunakan parameter `--tag` untuk memberikan nama image, dimana penamaan image menggunakan ID docker yang sudah disimpan sebelumnya dalam variabel `$DOCKERID`.  
```
$ docker image build --tag $DOCKERID/linux_tweet_app:1.0 .
```
Cek dengan perintah `docker image ls`.  
![perintah docker image ls di terminal](./09-docker-image-ls.png)  
  
Buat container baru menggunakan image yang sudah dibuat sebelumnya.  
```
$ docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:1.0
```
![perintah docker container run](./10-docker-container-run.png)  
  
Berikut adalah tampilan dari aplikasi yang dibuka pada browser.  
![tampilan halaman web dengan latar belakang biru dan gambar pinguin linux](./11-linux-tweet-app.png)  
  
Hapus container dengan perintah:  
```
$ docker container rm --force linux_tweet_app
```
  
### Task 3: Modify a running website
Jalankan kembali container dengan menambahkan parameter `--mount`.  
```
$ docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 --mount type=bind,source="$(pwd)",target=/usr/share/nginx/html \
 $DOCKERID/linux_tweet_app:1.0
```

Kemudian copy file `index-new.html` dengan nama `index.html`.  
```
$ cp index-new.html index.html
```

Cek halaman web, maka tampilan akan berubah menjadi seperti berikut:  
![tampilan web dengan latar belakang orange dan gambar pinguin linux](./12-linux-tweet-app-modified.png)  
  
Hapus container dengan perintah:  
```
$ docker rm --force linux_tweet_app
```
  
Lalu jalankan container baru dengan perintah yang sama seperti sebelumnya.  
```
$ docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:1.0
```

Cek halaman web, ternyata tampilan halaman web kembali ke tampilan awal sebelum dimodifikasi.  
![tampilan halaman web dengan latar belakang biru dan gambar pinguin linux](./11-linux-tweet-app.png)  
  
Hapus kembali container dengan perintah:  
```
$ docker rm --force linux_tweet_app
```
  
Lakukan update image dengan cara build ulang image menggunakan perintah:  
```
$ docker image build --tag $DOCKERID/linux_tweet_app:2.0 .
```
  
Cek dengan perintah `docker image ls`  
![perintah docker image ls](./13-docker-build-image.png)  
  
Jalankan container menggunakan image yang telah diupdate.  
```
$ docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:2.0
```

Jalankan juga container menggunakan image yang belum diupdate.  
```
$ docker container run \
 --detach \
 --publish 8080:80 \
 --name old_linux_tweet_app \
 $DOCKERID/linux_tweet_app:1.0
```

Cek pada web browser.  
![tampilan web dengan latar belakang orange dan gambar pinguin](./12-linux-tweet-app-modified.png)  
![tampilan web dengan latar belakang biru dan gambar pinguin linux](./11-linux-tweet-app.png)  

Selanjutnya adalah melakukan push image ke Docker Hub. Login dahulu dengan perintah `docker login`.  
![login docker di terminal](./14-docker-login.png)  
  
Push image versi pertama dengan perintah:  
```
$ docker image push $DOCKERID/linux_tweet_app:1.0
```

Kemudian push image versi kedua dengan perintah:  
```
$ docker image push $DOCKERID/linux_tweet_app:2.0
```
  
Cek image pada repository docker hub.  
![tampilan repository linux_tweet_app di docker hub](./15-image-docker-hub.png)  
  
Selesai.
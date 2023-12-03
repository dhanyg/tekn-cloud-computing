# Praktikum Teknologi Cloud Minggu ke-11
## Application Containerization and Microservice Orchestration
Pada praktikum kali ini akan dipelajari tentang *containerization* menggunakan Docker dan menjalankan berbagai aplikasi secara *microservices*. Beberapa service yang berjalan dalam microservice ini antara lain: 
- Aplikasi web yang ditulis dalam bahasa PHP dan menggunakan server Apache yang akan mengambil URL sebagai masukan dan mengekstrak tautan yang ada di dalamya.
- Server API yang ditulis menggunakan Python (dan Ruby) yang akan memproses ekstraksi tautan dan mengembalikan respons JSON.
- Cache Redis yang digunakan oleh server API untuk menghindari pengambilan berulang dan ekstraksi tautan untuk halaman yang sudah di-*scrap*.

### Stage Setup
Cloning repository.
```
$ git clone https://github.com/ibnesayeed/linkextractor.git
$ cd linkextractor
$ git checkout demo
```
![cloning repository dan checkout branch demo](./01-git-clone.png)  
  
### Step 0: Basic Link Extractor Script
```
$ git checkout step0
$ tree
```
![checkout ke branch step0 dan melihat daftar isi direktori](./02-git-checkout.png)  
  
Buat sebuah `Dockerfile` kemudian isi dengan script berikut:  
```
FROM       python:3
LABEL      maintainer="Sawood Alam <@ibnesayeed>"

RUN        pip install beautifulsoup4
RUN        pip install requests

WORKDIR    /app
COPY       linkextractor.py /app/
RUN        chmod a+x linkextractor.py

ENTRYPOINT ["./linkextractor.py"]
```
![membuat Dockerfile](./03-create-dockerfile.png)  
  
Build image menggunakan Dockerfile yang sudah dibuat sebelumnya.  
```
$ docker image build -t linkextractor:step1 .
```
![buil image dengan perintah docker build](./04-docker-build.png)  
  
Jalankan sebuah container dari image yang sudah dibuat sebelumnya dan tambahkan perintah `http://example.com` pada saat menjalankan containernya.  
```
$ docker container run -it --rm linkextractor:step1 http://example.com/
```
![menjalankan container dengan perintah docker container run](./05-docker-container-run.png)  
  
Jalankan kembali sebuah container dengan perintah yang sama namun dengan mengubah perintah pada bagian akhir menjadi `https://training.play-with-docker.com/`.  
```
$ docker container run -it --rm linkextractor:step1 https://training.play-with-docker.com/
```
![menjalankan container dengan perintah docker container run](./06-docker-container-run.png)  
  
### Step 2: Link Extractor Module with Full URI and Anchor Text
Checkout ke branch `step2`.  
```
$ git checkout step2
$ tree
```
![checkout ke branch step2](./07-git-checkout.png)  
  
Terdapat modifikasi pada file `linkextractor.py`. Untuk melihat isinya gunakan perintah `cat`.  
![isi file linkextractor.py](./08-py-file.png)  
  
Build image baru menggunakan Dockerfile pada direktori saat ini.  
```
$ docker image build -t linkextractor:step2 .
```
![membuat image dengan perintah docker build](./09-docker-build.png)  
  
Jalankan sebuah container baru menggunakan image terbaru.  
```
$ docker container run -it --rm linkextractor:step2 https://training.play-with-docker.com/
```
![menjalankan container dari image terbaru dengan perintah docker container run](./10-docker-container-run.png)  
  
Jalankan juga sebuah container baru dengan menggunakan image yang lama.  
```
$ docker container run -it --rm linkextractor:step1 https://training.play-with-docker.com/
```
![menjalankan container dari image lama menggunakan perintah docker container run](./11-docker-container-run.png)  
  
### Step 3: Link Extractor API Service
Checkout ke branch `step3`.  
```
$ git checkout step3
$ tree
```
![pindah ke branch step3 dan melihat isi direktorinya](./12-git-checkout.png)  
  
Terdapat penambahan file baru yakni `main.py` dan `requirements.txt`.  
Terdapat perubahan pada file `Dockerfile`. Untuk melihat isinya gunakan perintah `cat`.  
![isi file Dockerfile](./13-docker-file.png)  
  
Gunakan perintah yang sama untuk melihat isi file `main.py`.  
```
$ cat main.py
```
![isi file main.py](./14-main-py.png)  
  
Build image baru menggunakan Dockerfile saat ini.  
```
$ docker image build -t linkextractor:step3 .
```
![membuat image baru dengan perintah docker build](./15-docker-build.png)  
  
Jalankan sebuah container baru dari image terbaru secara *detach*, menggunakan port `5000:5000`, dan beri nama container dengan `linkextractor`.  
```
$ docker container run -d -p 5000:5000 --name=linkextractor linkextractor:step3
```
![menjalankan container baru dari image terbaru menggunakan perintah docker container run](./16-docker-container-run.png)  
  
Akses service pada container `linkextractor` dengan melakukan *curl* ke endpoint http://localhost:5000/api/http://example.com/.  
```
$ curl -i http://localhost:5000/api/http://example.com/
```
![melakukan curl ke endpoint http://localhost:5000/api/http://example.com/](./17-curl.png)  
  
### Step 4: Link Extractor API and Web Front End Services
Checkout ke branch `step4`.  
```
$ git checkout step4
$ tree
```
![berpindah ke branch step4 dan melihat isi direktori](./18-git-checkout.png)  
Terdapat perubaha struktur direktori dimana saaat ini terdapat dua buah direktori yakni `api` dan `www`. Direktori `api` digunakan untuk menyimpan konfigurasi dan kode untuk service backend, sementara `www` digunakan untuk menyimpan kode file web (user interface). Terdapat juga file `docker-compose.yml` yang berisi konfigurasi untuk membuat service api dan web.  

Lihat isi dari file `docker-compose.yml` menggunakan perintah `cat`.  
![isi file docker-compose.yml](./19-docker-compose.png)  
  
Lihat isi dari file `index.php` dengan menggunanakan perintah `cat`.  
![isi file index.php](./20-index-php.png)  
  
Jalankan container dengan menggunakan perintah `docker-compose up`.  
```
$ docker-compose up -d --build
```
![menjalankan container dengan perintah docker-compose up](./21-docker-compose-up.png)  
  
Lakukan pengecekan dengan mengakses endpoint http://localhost:5000/api/http://example.com/ menggunakan perintah `curl`.  
```
$ curl -i http://localhost:5000/api/http://example.com/
```
![mengakses endpoint menggunakan perintah curl](./22-curl.png)  
  
Lakukan git reset dan `docker compose down`.  
```
$ git reset --hard
$ docker-compose down
```
![melakukan git reset dan menghapus container menggunakan perintah docker-compose down](./23-git-reset.png)  
  
### Step 5: Redis Service for Caching
Checkout ke branch `step5`.  
```
$ git checkout step5
$ tree
```
![berpindah ke branch step5 dan melihat isi direktori](./24-git-checkout.png)  
  
Terdapat beberapa perubahan antara lain penambahan file `Dockerfile` di dalam direktori `www` serta modifikasi pada file `main.py` dan `docker-compose.yml`.  
![isi file www/Dockerfile](./25-cat-dockerfile.png)  
![potongan isi file main.py](./26-cat-main-py.png)  
![isi file docker-compose.yml](./27-cat-docker-compose.png)  
  
Jalankan container menggunakan perintah `docker-compose up`.  
```
$ docker-compose up -d --build
```
![menjalankan container dengan perintah docker-compose up](./28-docker-compose-up.png)  
  
Lakukan pengecekan redis menggunakan perintah berikut:  
```
$ docker-compose exec redis redis-cli monitor
```
Pada saat kita melakukan ekstrak halaman web maka akan muncul entry log dari redis.  
  
Lakukan git reset dan `docker-compose down`.  
```
$ git reset --hard
$ docker-compose down
```
![melakukan git reset dan menghapus container dengan perintah docker-compose down](./29-git-reset.png)  
  
### Step 6: Swap Python API Service with Ruby
Checkout ke branch `step6`.  
```
$ git checkout step6
$ tree
```
![berpindah ke branch step6 dan melihat isi direktori](./30-git-checkout.png)  
Terdapat perubahan dimana pada direktori `api`, server yang digunakan sebagai backend menggunakan Ruby.  

Lihat isi file `linkextractor.rb` menggunakan perintah `cat`.  
![potongan isi file linkextractor.rb](./31-linkextractor-rb.png)  
  
Lihat isi file `Dockerfile` dengan menggunakan perintah `cat`.  
![isi file Dockerfile](./32-dockerfile.png)  
  
Lihat isi file `docker-compose.yml` dengan menggunakan perintah `cat`.  
![isi file docker-compose.yml](./33-cat-docker-compose.png)  
  
Jalankan container menggunakan perintah `docker-compose up`.  
```
$ docker-compose up -d --build
```
![menjalankan container menggunakan perintah docker-compose up](./34-docker-compose-up.png)  
  
Lakukan pengecekan dengan mengakses endpoint http://localhost:4567/api/http://example.com/ menggunakan perintah `curl`.  
![akses endpoint http://localhost:4567/api/http://example.com/ menggunakan perintah curl](./35-curl.png)  
  
Hapus container menggunakan perintah `docker-compose down`.  
![menghapus container menggunakan perintah docker-compose down](./36-docker-compose-down.png)  
  
Selesai.
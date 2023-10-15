# Menginstal Apache Ofbiz
Pada praktikum minggu ke-5 ini mahasiswa diberikan latihan untuk menginstal software *Apache Ofbiz* yakni sebuah *open source* ERP. Latihan ini menggunakan Ubuntu 18.04 sebagai sistem operasinya.  
  
## Langkah-langkah Instalasi
Pastikan di dalam sistem operasi sudah terinstal java. Untuk mengeceknya gunakan perintah `java --version`  
  
Apabila java belum terinstal, silakan instal terlebih dahulu dengan perintah berikut:  
```
$ sudo apt install default-jre -y
```
Perintah di atas akan menginstal `openjdk-11` ke dalam sistem.  
  
Setelah instalasi selesai, cek kembali dengan perintah `java --version`.  
![perintah java --version pada terminal](./01-java-version.png)  
  
Selanjutnya unduh installer Apache Ofbiz pada tautan [ini](https://dlcdn.apache.org/ofbiz/apache-ofbiz-18.12.08.zip)  
```
$ wget https://dlcdn.apache.org/ofbiz/apache-ofbiz-18.12.08.zip
```
  
Lakukan *unzip* dengan perintah berikut:  
```
$ unzip apache-ofbiz-18.12.08.zip -d .
```
  
Maka akan muncul direktori installer Apache Ofbiz. Ubah nama direktori menjadi `ofbiz` agar tidak terlalu panjang.  
```
$ mv apache-ofbiz-18.12.08 ofbiz
```
![perintah ls dan mv pada terminal](./02-rename-directory.png)  
  
Unduh `gradle-wrapper.jar` dan `gradle-wrapper.properties` melalui tautan di bawah ini:  
- gradle-wrapper.jar: https://github.com/gradle/gradle/blob/v5.0.0/gradle/wrapper/gradle-wrapper.jar
- gradle-wrapper.properties: https://github.com/gradle/gradle/blob/v5.0.0/gradle/wrapper/gradle-wrapper.properties  
  
Setelah itu *copy* kedua file tadi ke dalam direktori `ofbiz/gradle/wrapper/`.  
```
cp gradle-wrapper.jar ofbiz/gradle/wrapper/
cp gradle-wrapper.properties ofbiz/gradle/wrapper/
```
![isi direktori wrapper](./03-ls-gradle-wrapper.png)  
  
Selanjutnya instal `gradle` dengan perintah:  
```
$ sudo apt install gradle -y
```
  
Masuk ke dalam direktori `ofbiz` kemudian ketikkan perintah berikut:  
```
./gradlew loadAll ofbiz
```
![menjalankan perintah ./gradlew loadAll ofbiz](./04-gradlew-loadall.png)  
  
Sistem akan melakukan proses unduh. Kemudian sistem akan menjalankan *Gradle daemon* sekaligus service Apache Ofbiz. Proses ini dapat memakan waktu yang lama. 
![proses download dan starting gradle daemon](./05-starting-gradle-daemon.png)  
  
Service Apache Ofbiz telah berjalan ditandai dengan tampilan seperti berikut di terminal:  
![apache ofbiz started di terminal](./06-ofbiz-started.png)  
  
Lakukan pengecekan dengan mengakses halaman https://localhost:8443/accounting.  
![halaman login apache ofbiz accounting](./07-login-ofbiz.png)  
  
Login menggunakan username `admin` dan password `ofbiz`. Setelah berhasil login kita akan diarahkan ke halaman *Accounting* - *Main*.  
![halaman accounting - main](./08-accounting-main.png)  
  
Lakukan pengecekan juga dengan mengakses halaman https://localhost:8443/ecommerce.  
![halaman ecommerce](./09-ofbiz-ecommerce.png)  
  
Selesai.

# SaaS: Perbedaan, Tentang, dan Cara Membangunnya  
  
## Perbedaan IaaS, PaaS, dan SaaS  
Pada *__IaaS__*, layanan yang disediakan berupa infrastruktur komputer. Secara sederhana, IaaS dapat diibaratkan kita menyewa satu atau beberapa server di suatu tempat kemudian server tersebut kita manage sendiri, seperti membuat VM, menginstal system operasi, dan mengkonfigurasi jaringan dan firewall. Contoh layanan *IaaS* adalah Amazon Web Service, Microsoft Azure, dan Google Cloud Platform.  
  
Sedangkan pada *__PaaS__*, layanan yang disediakan berupa *environment* untuk membangun suatu aplikasi. *Development environment* seperti *programming language*, web server, serta database telah tersedia sehingga pelanggan (yang biasanya seorang developer) tinggal menaruh *source code* aplikasinya dan melakukan *testing* dan *deployment*. Contoh layanan *PaaS* adalah Google App Engine dan Heroku.  
  
Sementara itu pada *__SaaS__*, layanan yang disediakan berupa aplikasi yang sudah siap dan tinggal dipakai oleh user. Pelanggan dapat mengakses layanan yang ditawarkan tanpa harus menginstalnya terlebih dahulu di komputer. Aplikasi dapat diakses dari berbagai perangkat dan dari mana saja. Contoh *Saas* adalah Google Drive, Microsoft 365, Spotify.  
  
  
## Lebih Lanjut Tentang SaaS
*SaaS* merupakan model layanan aplikasi dimana aplikasi tersebut diletakkan secara terpusat di sisi penyedia layanan dan dapat diakses oleh pelanggan melalui internet. Pelanggan tidak perlu repot melakukan *download*, instal, atau update aplikasi karena itu semua sudah ditangani oleh penyedia layanan.  

### Mengapa Menggunakan Saas?
Dari sisi pelanggan, *SaaS* merupakan cara paling mudah dan nyaman untuk menggunakan layanan atau produk digital karena pelanggan cukup mengaksesnya melalui web browser. Kemudahan lainnya adalah pelanggan dapat melanjutkan atau berhenti berlangganan kapan saja.  
  
Dari sisi bisnins, bentuk *SaaS* memungkinkan penyedia layanan untuk menawarkan produk atau layanannya secara global dan memiliki control terhadap produknya.  

### Kelebihan Arsitektur SaaS  
SaaS memiliki beberapa keunggulan jika dibandingkan dengan aplikasi tradisional antara lain:  
- Kemudahan dalam penggunaan, dimana aplikasi berbasis *SaaS* cukup diakses melalui web browser.
- Lebih ekonomis karena adanya opsi langganan bulanan atau tahunan.
- Pelanggan tidak perlu memikirkan bagaimana cara mengamankan data mereka karena itu sudah ditangani dan menjadi tanggung jawab penyedia layanan.
- Proses *update* atau perbaikan yang tidak membutuhkan waktu lama dan tidak perlu dilakukan oleh pelanggan karena semuanya sudah dilakukan oleh penyedia layanan.
  
### Kekurangan Arsitektur SaaS
Sementara itu terdapat beberapa kekurangan dari *SaaS* yakni:
- Kurangnya kendali atas aplikasi atau layanan yang digunakan karena kendali penuh berada di sisi penyedia layanan.
- Performa aplikasi yang terkadang masih kurang jika dibandingkan dengan aplikasi yang dipasang on-premise.
- Kekhawatiran akan data dan privasi yang bisa disalahgunakan karena kendali penuh berada pada penyedia layanan.
  
### Komponen Penting Dalam SaaS
Beberapa komponen yang penting dalam pengembangan layanan SaaS antara lain:
- __Keamanan__. Dimana penyedia layanan bertanggung jawab untuk menyimpan dan mengamankan data pelanggannya.
- __Privasi__. Penyedia layanan *SaaS* juga harus mematuhi kebijakan-kebijakan privasi yang berlaku di berbagai wilayah.
- __Kustomisasi dan Konfigurasi__. Penyedia perlu mempertimbangkan opsi untuk kustomisasi dan konfigurasi tambahan yang mungkin saja dibutuhkan oleh pelanggan untuk menyesuaikan dengan model bisnis mereka.  
  
Pertimbangan-pertimbangan lain yang juga penting jika kita ingin menjalankan layanan *SaaS* adalah skalabilitas, *zero downtime* dan *SLA*, serta *multi-tenant architecture*.  

## Tahapan Untuk Membangun Aplikasi SaaS
### 1. Menentukan Bahasa Pemrograman
Jika kita ingin membangun aplikasi berbasis *cloud* tentunya kita membutuhkan bahasa pemrograman yang modern. Salah satunya yang bisa digunakan adalah *Python*.
### 2. Menentukan Database
Terdapat banyak jenis database yang bisa digunakan dalam pengembangan aplikasi. Namun untuk jenis aplikasi berbasis *cloud*, disarankan untuk menggunakan database dengan jenis *document-oriented*. Dimana *document-oriented database* dirasa lebih fleksibel untuk menangani data yang bervariasi dan memiliki dukungan yang baik dengan bahasa pemrograman yang modern. Salah satu contoh yang bisa digunakan adalah *MongoDB*.
### 3. Menggunakan System Queueing
Teknologi *Message Queuein*g memungkinkan aplikasi berbasis web yang kita buat untuk berinteraksi dengan integrasi pihak ketiga, *API*, atau layanan lainnya secara *asynchronous*. Contoh teknologi m*essage queueing* adalah *RabbitMQ*.
### 4. Menyiapkan Infrastruktur
Salah satu infrastruktur yang mendukung untuk pengembangan layanan berbasis cloud adalah Amazon Web Service. Dimana kita dapat menginstal bahasa pemrograman dan database di dalam *Amazon EC2* dan dapat melakukan *scaling* ketika dibutuhkan. Selain itu terdapat *Amason S3* yang digunakan sebagai tempat penyimpanan data.
### 5. Membuat CDN
CDN atau *Content Delivery Network* memungkinkan kita untuk menjalankan layanan *SaaS* dengan performa dan ketersediaan yang tinggi. CDN akan mendistribusikan layanan kita kepada pelanggan sesuai dengan lokasi yang paling dekat dengan pelanggan. Dengan begitu performa dan kecepatan layanan akan menjadi lebih baik.
  
  
*Referensi:*
- https://www.quora.com/What-is-the-difference-between-IaaS-SaaS-and-Paas
- https://hackernoon.com/saas-software-as-a-service-platform-architecture-757a432270f5
- https://www.devteam.space/blog/saas-software-as-a-service-platform-architecture/
- https://usersnap.com/blog/cloud-based-saas-architecture-fundamentals/
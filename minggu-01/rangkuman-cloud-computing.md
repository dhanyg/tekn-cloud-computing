# Cloud Computing

## Pendahuluan
*Cloud computing* merupakan layanan pengaksesan sumber daya komputer secara *on-demand*.

5 karakteristik *cloud computing* menurut **NIST** (*National Institute of Standards anf Technology*):
* *__On-demand self service__*. Pelanggan dapat melakukan aktivitas komputasi sesuai kebutuhan tanpa harus berinteraksi langsung dengan penyedia layanan.
* *__Broad network access__*. Tersedia secara daring dan dapat diakses melalui berbagai perangkat gawai seperti laptop, tablet, atau ponsel.
* *__Resource pooling__*. Sumber daya komputasi digabungkan untuk kemudian dilayankan ke pelanggan dengan model *multi-tenant* dengan alokasi sumberdaya yang dinamis sesuai kebutuhan pelanggan.
* *__Rapid elasticity__*. Kemampuan untuk menambah atau mengurangi jumlah sumber daya komputasi secara elastis dan cepat sesuai dengan permintaan pelanggan.
* *__Measured service__*. Penggunaan sumber daya komputasi dipantau, dikendalikan, dan dilaporkan agar penyedia layanan maupun pelanggan mendapatkan transparansi informasi.

Layanan *cloud computing* diklaim dapat meminimalisir biaya infrastruktur IT pada suatu perusahaan. Beberapa nilai tambah dari layanan *cloud computing* antara lain:
* *__Cost reductions__*. Pelanggan tidak perlu mengeluarkan biaya pembelian perangkat server, namun cukup membayar biaya penggunaan layanannya saja.
* *__Device independence__*. Artinya pengaksesan layanan *cloud* dapat dilakukan dari lokasi mana saja dan melalui berbagai jenis perangkat.
* *__Maintenance__*. Proses maintain dilakukan oleh pihak penyedia layanan sehingga pelanggan tidak perlu memikirkan bagaimana me-maintain layanan yang mereka gunakan.
* *__Multitenancy__*. Memungkinkan *sharing resource* antar pelanggan.
* *__Performance__*. Performa layanan *cloud* dimonitor oleh penyedia layanan.
* *__Productivity__*. Meningkatkan produktivitas dikarenakan user dapat mengakses data yang sama secara bersamaan tanpa harus menunggu satu sama lain.
* *__Availability__*. Layanan *cloud* memiliki *redundant site* yang membuat layanan dapat terus tersedia dan digunakan oleh pelanggan.
* *__Scalability and Elasticity__*. Penggunaan sumber daya komputasi dapat diubah secara dinamis sesuai kebutuhan.
* *__Security__*. Keamanan diatur oleh penyedia layanan, namun pada kasus tertentu, pelanggan dapat melakukan pengamanan sendiri dengan menggunakan layanan *cloud* yang bersifat privat.

## Jenis Layanan
Terdapat beberapa service model pada layanan *cloud computing* antara lain:

### IaaS (*Infrastructure as a Service*)
Model ini merujuk pada layanan daring yang menyediakan *high-level API* untuk mengabstraksi aktivitas-aktivitas seperti *physical computing resources*, *location*, *data partioning*, *scaling*, *security*, *backup*, dan sebagainya.

Pada *__IaaS__*, pelanggan dapat menjalankan perangkat lunak apapun yang mencakup sistem operasi dan aplikasi. Pelanggan tidak dapat mengelola atau mengendalikan infrastruktur cloud yang mendasarinya, namun mereka memiliki kendali atas sistem operasi, penyimpanan, serta aplikasi yang mereka gunakan.

### PaaS (*Platform as a Service*)
Pada *__PaaS__*, penyedia layanan menawarkan *development environment* seperti sistem operasi, *environtment* bahasa pemrograman, *database*, serta web server. Pelanggan, yang biasanya merupakan seorang pengembang (*developer*) dapat melakukan proses pengembangan dan menjalankan aplikasinya pada platform *cloud* tersebut tanpa harus membeli dan mengelola infrastrukturnya.

### Saas (*Software as a Service*)
Pada model ini, layanan yang disediakan berupa *software* yang telah diinstal pada server penyedia layanan *cloud*. Pelanggan benar-benar hanya tinggal menggunakannya saja tanpa perlu melakukan instalasi di komputer mereka, karena akses aplikasi dapat dilakukan dari mana saja dan menggunakan *browser* dari berbagai perangkat secara daring. Contoh dari *__Saas__* adalah *Google Drive* dan *Google Docs*.

### BaaS (*Backend as a Service*)
Model ini merupakan model yang belum lama muncul. Layanannya berupa penyediaan akses agar aplikasi dari pelanggan (*developer*) dapat terhubung ke *cloud storage* atau *cloud service* lainnya. Layanan yang disediakan meliputi manajemen user, *push notifications*, integrasi dengan *social networking service* dan banyak lainnya.

### FaaS (*Function as a Service*)
*__FaaS__* merupakan jenis layanan *cloud* yang memungkinkan pengembang (*developer*) untuk membangun, menjalankan, dan mengelola paket aplikasinya sebagai sebuah *function* tanpa harus memikirkan bagaiamana mengatur dan memelihara infrastrukturnya.

## Jenis Penyebaran (*Deployments*)
Secara umum terdapat tiga jenis *deployment* pada *cloud computing* yakni:
1. *__Private__*, dimana infrastruktur *cloud* diperuntukkan untuk sebuah organisasi tertentu dan dapat dikelola secara internal ataupun oleh pihak ketiga.
2. *__Public__*, berarti layanan *cloud* di akses secara publik melalui jaringan internet. Biasanya layanan yang disediakan digunakan juga oleh pelanggan yang lain (*sharing resource*).
3. *__Hybrid__*, merupakan gabungan antara *private* dan *public*. Misalnya ketika sebuah organisasi yang memiliki *cloud* sendiri membutuhkan interkoneksi dengan sistem lain yang berjalan pada layanan *cloud* publik.

Selain itu terdapat pula beberapa jenis *deployment* lainnya seperti: *community*, *distributed*, *multy*, *poly*, *hybrid*, *big data*, dan *high-performance computing* (HPC).

## Isu Keamanan dan Privasi
Layanan *cloud computing* tak lepas dari isu keamanan dan privasi. Hal ini dikarenakan penyedia layanan *cloud* dapat mengakses informasi yang kita berikan atau data yang kita simpan yang kemudian berpotensi terhadap perubahan data atau bahkan kehilangan data. Untuk meminimalisir hal seperti itu perlu adanya aturan dan perjanjian yang jelas antara penyedia layanan dengan pelanggan.

Pihak penyedia layananpun juga perlu untuk memberikan pengamanan terhadap data pelanggan seperti enkripsi atau penggunaan *identity management system*. *Private cloud* dirasa lebih aman daripada *public cloud* namun bagi sebagian yang lainnya *public cloud* dianggap lebih fleksibel dan membutuhkan lebih sedikit waktu, tenaga, dan juga biaya. Meski begitu penyedia layanan *cloud* juga bertanggung jawab untuk memprioritaskan keamanan data pelanggannya.

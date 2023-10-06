# Deploy Some Code - DIY
Berikut adalah langkah-langkah *deploy* repositori yang ada di Github menggunakan *Cyclic*:  
1. Login atau *Sign-up* menggunakan akun Github ke https://app.cyclic.sh/api/login.  
2. Pada dashboard *Cyclic*, klik pada bagian *Link Your Own*.  
![tampilan dashboard cylic pada bagian Link Your Own](./18-cyclic-dashboard.png)  
3. Ketikkan nama repositori yang akan di-*deploy* kemudian pilih nama respositori tersebut.  
![input memilih nama repositori](./19-choose-repo.png)  
4. Proses selanjutnya adalah menghubungkan *Cyclic* dengan repositori yang telah dipilih. Pada bagian ini pilih opsi *Advanced*.  
![halaman konek cyclic ke repositori](./20-connect-cyclic.png)  
5. Kemudian masuk ke bagian *Build Options*, lalu aktifkan konfigurasi *Static Site*. Konfigurasi ini digunakan apabila di dalam repositori kita hanya terdapat file statik seperti `index.html` atau file-file html lainnya. Jika sudah, klik tombol *Connect Cyclic*.  
![konfigurasi static site](./21-enable-static-site.png)  
6. Masukkan password login Github.  
![halaman konfirmasi login github](./22-confirm-access.png)  
7. Pilih repositori, kemudian klik tombol *Approve and Install*.  
![memilih repositori](./23-approve-and-install.png)  
8. Proses *deploy* akan berlangsung. Setelah selesai akan muncul alamat situs yang bisa diakses.  
![proses deploy cyclic](./24-deploy-process.png)  
9. Kunjungi halaman web dengan mengeklik tautan yang telah diberikan.  
![tampilan halaman web yang telah di-deploy lewat Cyclic](./25-web-view.png)  
  
Selesai.
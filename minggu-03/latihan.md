# Deploy Some Code! - Quick Start
## Deploy to Cyclic 
1. Klik tombol *[deploy to cyclic](https://app.cyclic.sh/api/app/deploy/cyclic-software/express-hello-world)* seperti pada gambar di bawah.   
![tombol Deploy to Cyclic](./01-quick-start.png)  
  
2. Klik tombol *Continue with Github*.  
![halaman continue with github](./02-continue-with-github.png)  
  
3. Klik tombol *Connect Cyclic*.  
![halaman connect cyclic](./03-connect-cyclic.png)  
  
4. Masukkan password untuk login github.  
![halaman konfirmasi akses github](./04-confirm-access.png)  
  
5. Pilih repository yang akan di-*deploy*. Kemudian klik *Approve and Install*.   
![halaman pemilihan repository](./05-aprove-and-install.png)  
  
6. Proses instalasi dan *deploy* akan berlangsung.  
![proses deploy di cyclic](./06-proses-instal.png)  
  
7. Setelah proses *deploy* berhasil, akan muncul alamat situs yang bisa diakses.  
![proses deploy sukses](./07-deploy-success.png)  
  
8. Kunjungi alamat situs yang sudah diberikan tadi.  
![tampilan web express hello world](./08-web-view.png)  
  
## Mengupdate Repositori
1. *Clone* repositori yang sudah dibuatkan pada saat menjalankan *quick start* tadi.  
```
git clone https://github.com/dhanyg/express-hello-world.git
```
  
2. Lakukan pengeditan pada *source code*.  
![kumpulan perintah terminal](./09-edit-code.png)  
  
3. Lakukan *commit* kemudian dilanjutkan dengan *push* ke github.  
![perintah git commit pada terminal](./10-git-commit.png)  
![perintah git push pada terminal](./11-git-push.png)  
  
4. Lakukan *compare & pull request* di Github.  
![melakukan compare & pull request di github](./12-compare-and-pull.png)  
  
5. Pastikan pada bagian *base* dan *compare* sudah sesuai. Kemudian klik tombol *Create pull request*.  
![Proses create pull request github](./13-create-pull.png)  
  
6. Selanjutnya lakukan *merge pull request* dan *confirm merge*.  
![tombol merge pull request github](./14-merge-pull.png)  
![tombol confirm merge github](./15-confirm-merge.png)  
  
7. Setelah repositori terupdate, kemudian cek pada *dashboard* *Cyclic* yakni pada bagian *Deployment History* maka akan muncul aktivitas *log* yang telah dilakukan di Github sebelumnya.  
![tampilan dashboard cyclic bagian deploymeny history](./16-deployment-history.png)  
  
8. Akses kembali alamat web dan sekarang tampilan halaman web telah terupdate.  
![tampilan web express hello world setelah update](./17-view-after-update.png)  
  
Selesai.
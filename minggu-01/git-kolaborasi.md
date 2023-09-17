## Git Kolaborasi (Kontributor)
Login ke github kemudian cari repositori yang akan di-*fork*. Dalam contoh ini adalah repository: https://github.com/batmandvs/Playground

## Fork dan Konfigurasi Remote
Klik pada tombol *Fork*.
![mengeklik tombol fork di github](./39-fork.png)

Beri nama untuk repositori yang akan diubah, atau biarkan sama dengan repositori aslinya. Tambahkan deskripsi jika perlu, kemudian klik *__Create fork__*.
![mengisi nama repo fork](./40-create-new-fork.png)

Lakukan *clone* ke lokal komputer dengan perintah:
```
$ git clone https://github.com/dhanyg/playground.git
```
![eksekusi perintah git clone pada terminal](./41-clone-repo.png)
![eksekusi perintah git status](./42-git-status.png)

Tambahkan *remote* untuk *upstream* dengan perintah `git remote add <url_repo_asli>`. Kemudian cek daftar *remote* dengan perintah `git remote -v`.
![eksekusi perintah git remote add pada terminal](./43-git-remote-add.png)

## Mengirimkan Pull Request
Pertama-tama lakukan sinkronisasi dengan repositori asli dengan perintah `git fetch upstream`.
![eksekusi perintah git fetch](./44-git-fetch-upstream.png)

Buat branch baru dengan nama `add-contributor`.
![eksekusi perintah git branch -b](./45-create-new-branch.png)

Lakukan pengeditan.
![mengedit file README.md](./46-edit-readme.png)

Kemudian lakukan *commit*.
![eksekusi perintah git commit](./47-git-commit.png)

Lakukan *push* ke repositori milik kita.
![eksekusi perintah git push](./48-git-push.png)

Selanjutnya masuk ke repositori milik kita di github dan pilih *Compare & pull request*.
![memilih compare and pull request di repo github kita](./49-compare-pull-request.png)

Pastikan *branch* asal dan tujuannya sudah sesuai. Berikan keterangan dan jika sudah selesai, klik *Create pull request*
![membuat pull request ke repo asli](./50-create-pull-request.png)

Sampai tahap ini proses membuat *Pull Request* telah selesai. Selanjutnya adalah menunggu *author* untuk memeriksa dan memutuskan akan melakukan *merge* atau *close* terhadap *pull request* yang telah dikirim.

Apabila pemilik menerima dan melakukan *merge* maka status dari *pull request* akan menjadi seperti berikut:
![pull request diterima dan di-merged](./51-merged.png)

Selesai.
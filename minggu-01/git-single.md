## Instalasi Git
Berikut cara menginstal git pada sistem operasi berbasis linux Ubuntu:
```
$ sudo apt install git -y
```
![install git](./01-install-git.png)

Setelah proses instalasi selesai, cek dengan perintah `git --version`:
![cek versi git](./02-git-version.png)

## Konfigurasi Git
Pertama, konfigurasi nama user Git:
```
$ git config --global user.name "Ramadhan Yoga"
```
Kedua, konfigurasi email user Git:
```
$ git config --global user.email "ramayoga04@gmail.com"
```
![konfigurasi nama dan email git](./03-git-config.png)

Kemudian cek konfigurasi dengan perintah `git config --list`:
![cek konfigurasi git](./04-git-config-list.png)

## Mengelola Repo
### 1. Membuat Repo
Login ke [Github](https://www.github.com) kemudian buat repositori baru dengan mengeklik tanda `+` di bagian atas kemudian pilih *New repository*.
![membuat repo baru di github](./05-create-new-repo.png)

Isi *__Repository Name__* dengan nama `tekn-cloud-computing` dan berikan deskripsi pada bagian *__Description__* (opsional), lalu klik tombol *__Create repository__*.
![proses pembuatan repo di github](./06-create-repository.png)

### 2. Clone Repo
Salin repositori yang telah dibuat di Github tadi ke lokal komputer dengan perintah:
```
$ git clone https://github.com/<username>/<nama_repositori>
```
![clone repo dari github ke lokal komputer](./07-clone-repo.png)

Jika proses *cloning* berhasil, maka akan muncul sebuah direktori baru dengan nama yang sama seperti nama repositori pada github.
![muncul folder hasil clone](./08-check-folder.png)

### 3. Mengubah Isi - Push Tanpa Branching dan Merging
Buat sebuah file dengan nama README.md.
```
vim README.md
```
![membuat file README.md](./09-create-readme.png)

Cek status *Branch* dengan perintah `git status`.
![output perintah git status](./10-git-status.png)

Tambahkan file yang sudah dibuat tadi ke dalam repo git menggunakan perintah `git add -A`
![menambahkan file dengan perintah git add](./11-git-add.png)

Lakukan *Commit* dengan perintah `git commit -m "Add: README.md"`.
![melakukan commit dengan perintah git commit -m](./12-git-commit.png)

Kemudian lakukan *Push* ke repositori github dengan perintah `git push origin main`.
![melakukan push ke github](./13-git-push-error.png)

Apabila terjadi error seperti pada gambar di atas, maka lakukan setup *Personal Access Token* terlebih dahulu sebelum melakukan *push*.

Login ke [Github](https://github.com) kemudian klik icon profile, kemudian Pilih `Settings > Developer Settings > Personal access tokens > Token (classic)`. Lalu klik *__Generate new token__*.
![membuat personal access token di github](./14-personal-access-token.png)

Konfirmasi dengan memasukkan password akun github.

Isi pada bagian *__Note__*. Set *__Expiration__* menjadi `No expiration`, dan berikan centang pada opsi `repo` dan `workflow`. Klik *__Generate token__*.
![memasukkan password github](./15-new-personal-access-token.png)

Silakan salin token yang sudah dibuat tadi. Kemudian lakukan push kembali. Pada bagian password, masukkan token yang sudah di-*copy* tadi.
![push dengan perintah git push origin main](./16-git-push-success.png)

### 4. Mengubah Isi dengan Branching dan Merging
Buat branch baru dengan nama `edit-readme-1`.
```
$ git checkout -b edit-readme-1
```

Kemudian edit file README.md.
![mengedit file REDME.md](./17-edit-readme.png)

Lakukan *commit* lagi ke repo git.
![commit dengan perintah git commit -am](./18-git-commit.png)

Lakukan *push* ke repositori github dari branch `edit-readme-1`.
```
$ git push origin edit-readme-1
```
![push dengan perintah git push origin main](./19-git-push.png)

Selanjutnya, lakukan *Pull Request* di Github.
![melakukan compare dan pull request di github](./20-compare-pull-request.png)
![memberikan deskripsi pull request](./21-create-pull-request.png)

Setelah itu lakukan *merge*.
![melakukan merge pull request](./22-merge-pull-request.png)

Terakhir, lakukan *merge* pada branch lokal.
```
$ git checkout main
$ git merge edit-readme-1
```
![merge pada branch lokal main](./23-merge-main.png)

### 5. Sinkronisasi
Untuk melakukan sinkronisasi antara repositori di github dengan repo lokal di komputer, gunakan perintah `git pull`.
![melakukan perintah git pull](./24-git-pull.png)

### 6. Membatalkan Perubahan
Buat branch baru dengan nama `edit-readme-2`.
```
$ git checkout -b edit-readme-2
```
![membuat branch baru dengan perintah git checkout -b](./25-create-new-branch.png)

Lakukan modifikasi pada file README.md.
![mengedit README](./26-edit-readme.png)

Kembali ke branch `main` dan cek isi file README.md
![mengecek isi README dari branch main](./27-git-checkout-main.png)

Hapus branch `edit-readme-2` dengan perintah `git branch -D edit-readme-2`. Kemudian cek daftar branch yang masih ada dengan perintah `git branch`.
![menghapus branch dengan perintah git branch -D](./28-delete-branch.png)

Untuk mengembalikan isi file README.md ke *state* awal sebelum diubah dari branch `edit-readme-2`, lakukan reset dengan perintah `git reset --hard`.
![melakukan reset dengan perintah git reset --hard](./29-git-reset.png)

### 7. Undo Commit Terakhir
Lakukan beberapa kali *commit* dan *push* ke repositori github.
![melakukan beberapa kali push](./30-git-push.png)
![](./31-git-log.png)

Kembalikan *state* ke posisi sebelum *commit* terakhir dengan menggunakan perintah `git revert HEAD`.
![melakukan git revert](./32-git-revert.png)
![](./33-cat-readme.png)

Apabila ingin mengembalikan perubahan yang sudah di-*commit* namun belum dilakukan *push* ke github, gunakan perintah `git reset --hard HEAD^`.
![melakukan git reset](./34-git-reset-head.png)
![](./35-cat-readme.png)

Untuk mengembalikan *state* ke perubahan yang sudah lama, gunakan perintah `git revert <posisi>`.
![melakukan git revert ke posisi tertentu](./36-git-revert.png)
![git revert conflict](./37-git-revert-error.png)
Edit kembali file README.md untuk memperbaiki error *conflict*. Setelah itu lanjut proses *revert* dengan perintah `git revert --continue`.
![memperbaiki conflict dan melanjutkan revert](./38-solve-merge-conflict.png)

Selesai.
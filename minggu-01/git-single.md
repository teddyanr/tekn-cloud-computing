# Instalasi Git

## Windows

1. Setelah download GIT, klik install kemudian akan muncul lisesnsi. Klick Next.

![01](img/instalasi-1.png)

2. Kemudian pilih lokasi instalasi, setelah memilih klik Next.

![02](img/instalasi-2.png)

3. Pilih komponen. Tidak perlu diubah-ubah, sesuai dengan default saja. Klik pada Next.

![03](img/instalasi-4.png)

4. Mengisi shotcut di start menu. Setelah selesai klik Net.

![04](img/instalasi-4.png)

5. Pilih editor yang akan digunakan bersama dengan Git. Pada pilihan ini, digunakan Notepad++.

![05](img/instalasi-5.png)

6. Pada saat instalasi, Git menyediakan akses git melalui Bash maupun command prompt. Pilih pilihan kedua supaya bisa menggunakan dari dua antarmuka tersebut. Bash adalah shell di Linux. Dengan memilih pilihan kedua maka kita bisa mengakses GIT dari command promt windows.

![06](img/instalasi-6.png)

7. Memilih eksejusi SSH. Pilih OpenSSH. Kemudian klik Next.

![04](img/instalasi-7.png)

8. Memilih HTTPS transport backend. Pilih Use the OpenSSL library. Klik Next.

![08](img/instalasi-8.png)

9. Memilih line ending, biasanya bisa diliah saat menggunakan text editor. Setelah selesai klik Next.

![09](img/instalasi-9.png)

10. Mengkonfigurasi terminal. Pilih use Windows default. Klik Next.

![10](img/instalasi-10.png)

11. Memilih konfigurasi default untuk git pull. Pilih default kemudian klik Next.

![11](img/instalasi-11.png)

12. Memilih credential manager, kemudian Next.

![12](img/instalasi-12.png)

13. Mengaktifkan system caching. Kemudian Next.

![13](img/instalasi-13.png)

14. Kemudian langsung klik Next.

![14](img/instalasi-14.png)

15. Proses instalasi, tunggu sampai selesai proses instalasi.

![15](img/instalasi-15.png)

16. Klik Next jika proses installasi telah selesai.

![16](img/instalasi-16.png)

17. Klik start menu kemudian klik bash

![17](img/instalasi-17.png)

18. Tampilan git bash command

![18](img/instalasi-18.png)

19. Tampilan git GUI.

![19](img/instalasi-19.png)

20. Cek apakah git sudah terinstall di computer dengan membuka CMD kemudian mengetikkan git –version.

![20](img/instalasi-20.png)


# Konfigurasi Git

Ada 2 hal yang perlu dikonfigurasi yaitu username dan email. Konfig harus disesuaikan dengan nama serta email yang digunakan untuk mendaftar di GitHub. Untuk melihat konfigurasi yang sudah ada. Gunakan perintah berikut:

![01](img/config-git.png)

Langkah ini cukup dilakukan sekali saja, kecuali jika ingin melakukan perubahan nama dan email.

# Mengelola Repo Sendiri di Account Sendiri
## Membuat Repo

1. Buka web [Github] (https://github.com/). kemudian masuk ke akun untuk membuat repository baru.

![01](img/repo-1.png)

2.	Membuat repository baru dengan nama tekn-cloud-computing, jika sudah selesai klik create repository.

![02](img/repo-2.png)

3.	Setelah selesai maka akan muncul repository seperti ini. Setelah langkah-langkah tersebut, repo akan dibuat dan bisa diakses menggunakan pola https://github.com/username/reponame.

![03](img/repo-3.png)

## Clone repository

Proses clone adalah proses untuk menduplikasikan remote repo di GitHub ke komputer lokal. Untuk melakukan proses clone, gunakan perintah berikut:

![01](img/repo-4.png)

Setelah perintah ini, di direktori awesome-project akan disimpan isi repo yang sama dengan di GitHub. Perbedaannya, di komputer lokal terdapat direktori .git yang digunakan secara internal oleh Git.

## Mengelola Repo

### Mengubah Isi - Push Tanpa Branching dan Merging

Perubahan isi bisa terjadi karena satu atau kombinasi beberapa hal berikut:
1.	File dihapus
2.	File diedit
3.	Membuat file / direktori baru
4.	Menghapus direktori

Untuk kasus-kasus tersebut, lakukan perubahan di komputer lokal, setelah itu push ke repo.

![01](img/repo-5.png)

### Mengubah Isi dengan Branching and Merging

1.	Buat branch untuk menampung perubahan-perubahan
2.	Lakukan perubahan-perubahan
3.	Add dan commit perubahan-perubahan tersebut ke branch
4.	Kembali ke repo master
5.	Buat pull request di GitHub
6.	Merge pull request di GitHub
7.	Merge branch untuk menampung perubahan-perubahan tersebut ke master.
8.	Selesai.

![01](img/repo-6.png)
![02](img/repo-7.png)

Membuat pull request agar bisa di-merge

![03](img/repo-8.png)

Setelah itu, Confirm Merge, branch yang kita kirimkan tadi sudah dimasukkan ke branch master. Setelah itu, merge di komputer lokal:
![04](img/repo-9.png)

### Sinkronisasi

![01](img/repo-10.png)

### Membatalkan Perubahan

![01](img/repo-11.png)
![02](img/repo-12.png)

### Undo Commit Terakhir

![01](img/repo-13.png)
![01](img/repo-14.png)

Contoh di atas adalah contoh untuk mengubah README.md dengan beberapa commit. Setelh itu, kita akan mengembalikan ke posisi terakhir sebelum commit terakhir.

![01](img/repo-15.png)

Jika commit sudah dilakukan, tetapi belum di-push ke repo GitHub (masih berada di lokal), cara membatalkannya:

![01](img/repo-16.png)
![02](img/repo-17.png)

Untuk kembali ke perubahan pada saat yang sudah lama, yang perlu dilakukan adalah melakukan git revert <posisi> kemudian mengedit secara manual kemudian push ke repo.

![03](img/repo-18.png)

Edit file tersebut, setelah itu simpan.

![04](img/repo-19.png)

Setelah itu, lanjutkan proses revert. Saat git revert --continue isikan pesan revert.

![05](img/repo-20.png)
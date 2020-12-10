# Application Containerization and Microservice Orchestration

## Stage Setup

1. Melakukan clonning repo dari [linkextractor] (https://github.com/ibnesayeed/linkextractor.git) dan kemudian mengakases derektori serta memeriksa branch dari repo tersebut.

<div align="center"><img src="img/acmo-01.png" width="500px"></div>


## Basic Link Extractor Script

1. Selanjutnya memeriksa branch step0 untuk melihat file didalamnya. Dan kemudian melihat isi dari file python linkextractor.

<div align="center"><img src="img/acmo-02.png" width="500px"></div>

2. Selanjutnnya Menajalankan file linkextractor.py.

<div align="center"><img src="img/acmo-03.png" width="500px"></div>


## Containerized Link Extractor Script

1. Memeriksa branch step1 untuk melihat file didalamnya.

<div align="center"><img src="img/acmo-04.png" width="500px"></div>

2. Kemudian melihat isi dari file Dockerfile.

<div align="center"><img src="img/acmo-05.png" width="500px"></div>

3. Membangun docker image. Dimana dengan menjalankan perintah seperti dibawah ini, dan sekaligus manghasilkan outputnya.

<div align="center"><img src="img/acmo-06_a.png" width="500px"></div>
<div align="center"><img src="img/acmo-06_b.png" width="500px"></div>

4. Kemudian setelah membuat docker image bernama linkextractor: step1 selanjutnya kita mengecek dengan melihat daftar/ist dari image docker yang sudah ada. Kemudian melakukan ekstrak docker image kita dengan mendapat URL.

<div align="center"><img src="img/acmo-07_a.png" width="500px"></div>
<div align="center"><img src="img/acmo-07_b.png" width="500px"></div>

5. Selanjutnya melakukan percobaan pada halaman web untuk melihat lebih banyak tautannya.

<div align="center"><img src="img/acmo-08.png" width="500px"></div>


## Link Extractor Module with Full URI and Anchor Text

1. Kemudian memeriksa branch step2 dan daftar file yang ada di dalamnya.

<div align="center"><img src="img/acmo-09.png" width="500px"></div>

2. Sehingga secara otomatis file linkextractor.py akan diupdate, berikut merupakan hasil update dari file tersebut.

<div align="center"><img src="img/acmo-10.png" width="500px"></div>

3. Kemudian membuat image baru.

<div align="center"><img src="img/acmo-11.png" width="500px"></div>

4. Sehingga image docker yang baru dibuat dengan nama linkextractor:step02, kita cek pada pada list image docker yang sudah ada.

<div align="center"><img src="img/acmo-12.png" width="500px"></div>

5. Kemudian menjalankan image docker tersebut dan menghasilkan keluaran seperti dibawah ini :

<div align="center"><img src="img/acmo-13.png" width="500px"></div>

6. Kemudian menjalankan image docker step1 yang sebelumnya dan menghasilkan keluaran yang masih sama seperti dibawah ini :

<div align="center"><img src="img/acmo-14.png" width="500px"></div>


## Link Extractor API Service

1. Selanjutnya memeriksa branch step3 dan isi file didalamnya.

<div align="center"><img src="img/acmo-15.png" width="500px"></div>

2. Kemudian cek file Dockerfile untuk melihat perubahannya.

<div align="center"><img src="img/acmo-16.png" width="500px"></div>

3. Selanjutnya melihat isi dari file main.py yang baru ditambahkan tersebut.

<div align="center"><img src="img/acmo-17.png" width="500px"></div>

4. Kemudian update image docker step3 ini dengan beberapa langkah perubahan, seperti dibawah ini.

<div align="center"><img src="img/acmo-18_a.png" width="500px"></div>
<div align="center"><img src="img/acmo-18_b.png" width="500px"></div>

5. Selanjutnya menjalankan container dalam mode (-d flag) sehingga terminal dapat tersedia untuk perintah yang lain saat container masih berjalan. Perhatikan juga bahwa disitu terdapat port 5000 dari container dengan host 5000 (menggunakan perintah -p 5000: 5000) agar dapat diakses dari host. Dan juga memberikan nama (--name = linkextractor) ke container untuk lebih mudah melihat log atau menghapus container. Serta melihat list image container yagg baru dibuat tersebut.

<div align="center"><img src="img/acmo-19.png" width="500px"></div>

6. Membuat permintaan HTTP dalam bentuk /api /url untuk mengakses server ini dan mengambil respons berisi link yang diekstrak.

<div align="center"><img src="img/acmo-20.png" width="500px"></div>

7. Karena container berjalan dalam mode terpisah, jadi tidak dapat melihat proses yang terjadi di dalam, tetapi dapat melihat log menggunakan linkextractor yang di tetapkan untuk container. Serta menghapus image container ini.

<div align="center"><img src="img/acmo-21.png" width="500px"></div>


## Link Extractor API and Web Front End Services

1. Selanjutnya memeriksa branch step4 dan isi file didalamnya.

<div align="center"><img src="img/acmo-22.png" width="500px"></div>

2. Kemudian melihat isi dari file docker-compose.yml dan www / index.php.

<div align="center"><img src="img/acmo-23.png" width="500px"></div>

3. Membuat mode terpisah untuk container.

<div align="center"><img src="img/acmo-24_a.png" width="500px"></div>
<div align="center"><img src="img/acmo-24_b.png" width="500px"></div>

4. Memeriksa daftar container yang sedang berjalan serta memastikan bahwa kedua container tersebut benar berjalan. Dan kemudian mengkases layanan API.

<div align="center"><img src="img/acmo-25_a.png" width="500px"></div>

Sehingga ketika megkasesnya pada URL, maka hasilnya seperti pada gambar dibawah ini :

<div align="center"><img src="img/acmo-25_b.png" width="500px"></div>

5. Selanjutnya memodifikasi file www/index.php mengganti semua kemunculan Link Extractor dengan Super Link Extractor. Kemudian mengembalikan perubahan, dan menonaktifkan container ini.

<div align="center"><img src="img/acmo-26.png" width="500px"></div>

6. Kemudian memeriksa branch step5 dan melihat isi file didalamnya.

<div align="center"><img src="img/acmo-27.png" width="500px"></div>

7. Memeriksa file Dockerfile yang baru di dalam direktori www.

<div align="center"><img src="img/acmo-28.png" width="500px"></div>

8. Selanjutnya, melihat isi file api/main.py dengan menggunakan server redis.

<div align="center"><img src="img/acmo-29.png" width="500px"></div>

9. Melihat hasil perubahan pada file docker-compose.yml.

<div align="center"><img src="img/acmo-30.png" width="500px"></div>

10. Selanjutnya melakukan eksekusi container ini, untuk bisa di buka pada browser.

<div align="center"><img src="img/acmo-31_a.png" width="500px"></div>
<div align="center"><img src="img/acmo-31_b.png" width="500px"></div>

Sehingga hasilnya ketika ditampilkan pada browser, seperti pada gambar dibawah ini :

<div align="center"><img src="img/acmo-31_c.png" width="500px"></div>

11. Selanjutnya memeriksa apakah layanan redis dipakai atau tidak.

<div align="center"><img src="img/acmo-32.png" width="500px"></div>

12. Memeriksa ketika folder www tidak tersedia pada container yang sedang berjalan. kemudian melakukan verifikasi bahwa perubahan yang dibuat secara lokal tidak berada dalam layanan yang berjalan dengan memuat ulang halaman web dan kemudian mengembalikan perubahan. Dan menonaktifkan cointainer ini.

<div align="center"><img src="img/acmo-33.png" width="500px"></div>


## Link Extractor API and Web Front End Services

1. Kemudian memeriksa branch step6 dan melihat isi file didalamnya.

<div align="center"><img src="img/acmo-34.png" width="500px"></div>

2. Selanjutnya melihat isi file linkextractor.rb, ini merupakan file ruby untuk mengelola dependency.

<div align="center"><img src="img/acmo-35.png" width="500px"></div>

3. Kemudian melihat isi file Dockerfile.

<div align="center"><img src="img/acmo-36.png" width="500px"></div>

4. Kemudian mengecek perubahan pada file docker-compose.yml.

<div align="center"><img src="img/acmo-37.png" width="500px"></div>

5. Kemudian selanjutnya mengeksekusi container ini.

<div align="center"><img src="img/acmo-38_a.png" width="500px"></div>
<div align="center"><img src="img/acmo-38_b.png" width="500px"></div>

6. Kemudian selanjutnya harus dapat mengakses API (menggunakan nomor port yang diperbarui):

<div align="center"><img src="img/acmo-39_a.png" width="500px"></div>

Sehingga hasil yang ditampilkan pada halaman browser seperti dibawah ini :

<div align="center"><img src="img/acmo-39_b.png" width="500px"></div>

7. Kemudian shut down container ini tetapi lognya akan tetap ada walaupun containernya hilang.

<div align="center"><img src="img/acmo-40.png" width="500px"></div>


<br>
<br>

**Sumber**
<br>
**https://training.play-with-docker.com/microservice-orchestration/**
# Software as a Service


## Perbedaan Iaas, PaaS, SaaS

1.	Infrastructure as a Service (IaaS)

Infrastructure as a Service adalah layanan komputasi awan yang menyediakan infrastruktur IT berupa CPU, RAM, storage, bandwith dan konfigurasi lain. Semua komponen tersebut digunakan untuk membangun komputer virtual. Komputer virtual dapat diinstal sistem operasi dan aplikasi sesuai kebutuhan. Keuntungan layanan IaaS ini adalah tidak perlu membeli komputer fisik sehingga lebih menghemat biaya. Konfigurasi komputer virtual juga bisa diubah sesuai kebutuhan. Misalkan saat storage hampir penuh, storage bisa ditambah dengan segera. Perusahaan yang menyediakan IaaS adalah Amazon EC2, Telkom Cloud dan BizNetCloud.

2.	Platform as a Service (PaaS)

Platform as a Service adalah layanan yang menyediakan computing platform. Biasanya sudah terdapat sistem operasi, database, web server dan framework aplikasi agar dapat menjalankan aplikasi yang telah dibuat. Perusahaan yang menyediakan layanan tersebutlah yang bertanggung jawab dalam pemeliharaan computing platform ini. Keuntungan layanan PaaS ini bagi pengembang adalah mereka bisa fokus pada aplikasi yang mereka buat tanpa memikirkan tentang pemeliharaan dari computing platform. Contoh penyedia layanan PaaS adalah Amazon Web Service dan Windows Azure.

3.	Software as a Service (SaaS)

Software as a Service adalah layanan komputasi awan dimana kita bisa langsung menggunakan aplikasi yang telah disediakan. Penyedia layanan mengelola infrastruktur dan platform yang menjalankan aplikasi tersebut. Contoh layanan aplikasi E-mail yaitu Gmail, Yahoo Mail dan Microsoft Outlook sedangkan contoh aplikasi media sosial adalah Twitter, Facebook dan Google+. Keuntungan dari layanan ini adalah pengguna tidak perlu membeli lisensi untuk mengakses aplikasi tersebut. Pengguna hanya membutuhkan perangkat klien komputasi awan yang terhubung ke internet.


## Arsitektur SaaS

Dengan model ini, satu versi aplikasi, dengan konfigurasi tunggal digunakan untuk semua pelanggan. Aplikasi diinstal pada beberapa mesin untuk mendukung skalabilitas (Horizontal scaling). Dalam beberapa kasus, versi kedua dari aplikasi diatur untuk menawarkan kelompok pelanggan tertentu dengan akses ke versi pra-rilis aplikasi untuk tujuan pengujian. Dalam model tradisional ini, setiap versi aplikasi didasarkan pada kode unik.

Ada 2 variasi SaaS

1.	Vertical SaaS
Software yang melayani apa yang dibutuhkan oleh industry.

2.	Vertical SaaS
Produk yang berfokus pada kategori software (marketing, sales, developer tools, HR) tetapi bersifat agnostik industri.


## SaaS (Software as a Service) Platform itecture.

Perangkat lunak telah didistribusikan ke pelanggan di berbagai saluran selama beberapa dekade terakhir. Saluran distribusi yang lebih baru dalam Perangkat Lunak sebagai Layanan (atau SaaS).

Jadi, mengapa Anda ingin menggunakan produk yang dikirimkan "sebagai layanan"?

**Konsumen**

Dari sudut pandang konsumen, produk SaaS adalah salah satu cara termudah untuk menggunakan layanan atau produk digital. Anda cukup mengaksesnya melalui web, membayar layanan, dan menggunakannya! Dalam beberapa tahun terakhir kami telah melihat munculnya ribuan produk SaaS yang ditargetkan untuk konsumen seperti: (Netflix, Microsoft Office 365, Amazon Prime, Indonesia, Facebook, Google Docs, dll)

**Bisnis**

Dari perspektif bisnis, produk perangkat lunak yang diberikan "sebagai layanan" memungkinkan bisnis menawarkan produk mereka dalam skala besar, secara global, sementara juga memungkinkan mereka untuk mempertahankan kendali keseluruhan atas produk mereka. Beberapa manfaat lain dari penerapan arsitektur SaaS dalam bisnis termasuk, namun tidak terbatas pada:

•	Mengurangi waktu ke pasar

•	Biaya perawatan lebih rendah

•	Peningkatan yang lebih mudah

### Fitur Utama dan Manfaat Platform SaaS

Solusi SaaS memiliki fitur yang berbeda dengan yang ada pada aplikasi tradisional yang diinstal di desktop Anda misalnya, berikut adalah beberapa manfaat lain yang dapat dihasilkan dari penerapan produk berbasis SaaS.


## Bagaimana membangun Aplikasi SaaS berbasis cloud.

Bisnis SaaS adalah industri yang tumbuh sangat cepat dan menarik lebih banyak orang dan perusahaan. Organisasi-organisasi ini lebih dan lebih aplikasi mengambang di awan. Penskalaan di cloud juga memiliki beberapa manfaat dan risiko penting. Untuk membuat layanan SaaS berbasis cloud hal pertama yang dibutuhkan adalah:

1.	Bahas pemrograman apa yang akan digunakan

Membangun produk untuk cloud berarti membangun produk dengan bahasa pemrograman modern. Selain kemampuan dan keterampilan pribadi, pilihan bahasa pemrograman Anda akan dipengaruhi oleh kemungkinan setiap bahasa. Ada berbagai bahasa pemrograman (modern) di luar sana sehingga sulit untuk memilih yang tepat.

2.	Database yang tepat untuk digunakan

Untuk membuat layanan SaaS berbasis Cloud dirokemendasikan menggunakan database dokumen daraipada menggunakan database dengan konsep tradisional. Mengapa memilih database berorientasi dokumen? Database dokumen mendapatkan informasi tipenya dari data itu sendiri. Jadi setiap contoh data dapat berbeda dari yang lain. Ini memungkinkan lebih banyak fleksibilitas, terutama saat menghadapi perubahan. Dan itu sering kali mengurangi ukuran database.

3.	System queuing

Dikenal sebagai teknologi Message Queuing (MSMQ), ini memungkinkan aplikasi web untuk berjalan pada waktu yang berbeda dan untuk berkomunikasi dengan berbagai integrasi pihak ketiga / API / dan layanan lainnya secara asinkron. Sebuah pesan (misalnya kueri yang meminta layanan pihak ketiga melalui API) ditempatkan ke antrean. Itu disimpan di sana sampai penerima mengambilnya. Antrian pesan memiliki batasan terkait ukuran dan jumlah data yang dikirimkan dalam antrian. Hal terbaik tentang sistem antrian modern adalah sistem ini dapat diskalakan dengan mudah.

4.	Content delivery network

Content delivery network (CDN) pada dasarnya adalah sistem server terdistribusi yang memungkinkan Anda menyajikan konten kepada pengguna aplikasi Anda dengan kinerja tinggi dan ketersediaan tinggi.

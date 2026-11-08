# Rangkuman Mata Kuliah: Rekayasa Perangkat Lunak (RPL)

## Konsep Dasar & Urgensi RPL

* **Definisi:** RPL adalah disiplin ilmu yang menerapkan pendekatan sistematis, terukur, dan terorganisasi untuk merencanakan, membangun, menguji, menerapkan, serta memelihara perangkat lunak.  
* **Inti Pemahaman:** RPL bukan hanya sekadar membuat program, tetapi proses menghasilkan solusi perangkat lunak yang tepat guna dan berkualitas.  
* **RPL vs Programming:**  
  * *Programming* hanya berfokus menulis kode untuk membuat fungsi tertentu, seperti membuat fungsi hitung total belanja.  
  * *RPL* mengelola seluruh proses pengembangan dari analisis hingga pemeliharaan untuk memastikan solusi sesuai kebutuhan, contohnya membangun sistem penjualan utuh hingga siap digunakan.  
* **Pentingnya RPL:** Tanpa RPL, pembuatan aplikasi berisiko salah kebutuhan, banyak *bug*, sulit digunakan penggunanya, data tidak aman, dan sistem sulit dikembangkan di kemudian hari.

## Tahapan SDLC (Software Development Life Cycle)

Proses sistematis pengembangan software terbagi menjadi 7 tahap:

1. **Planning:** Merencanakan tujuan, ruang lingkup, waktu, dan anggaran.  
2. **Analysis:** Mengidentifikasi masalah dan menggali kebutuhan pengguna.  
3. **Design:** Merancang arsitektur sistem, struktur database, alur, dan antarmuka.  
4. **Implementation:** Menulis kode program dan menerjemahkan rancangan menjadi software.  
5. **Testing:** Menguji fungsi aplikasi untuk menemukan *bug*.  
6. **Deployment:** Memasang aplikasi ke lingkungan kerja agar bisa digunakan.  
7. **Maintenance:** Memperbaiki kesalahan, memperbarui keamanan, dan menambah fitur.

## 6 Karakteristik Software Berkualitas

Kualitas software tidak sekadar diasumsikan, tetapi harus diuji dan memenuhi:

* **Functionality:** Fungsi berjalan sesuai kebutuhan.  
* **Reliability:** Aplikasi stabil dan dapat diandalkan.  
* **Usability:** Antarmuka mudah dipahami dan digunakan.  
* **Performance:** Respons cepat dan penggunaan sumber daya baik.  
* **Security:** Data dan akses pengguna terlindungi.  
* **Maintainability:** Sistem mudah diperbaiki dan dikembangkan.

## Model / Metode Pengembangan

Tidak ada model yang paling sempurna, pemilihannya menyesuaikan dengan karakter proyek:

* **Waterfall:** Tahapan dilakukan berurutan, sangat cocok jika kebutuhan aplikasi sudah stabil dan jelas.  
* **Prototype:** Membuat model awal untuk mendapatkan masukan, cocok jika kebutuhan pengguna belum sepenuhnya jelas.  
* **Agile:** Menggunakan metode iterasi (*sprint*) yang fleksibel, kolaboratif, dan mudah beradaptasi dengan perubahan.  
* **Spiral:** Siklus pengembangan yang sangat memperhatikan analisis risiko proyek.

---

## Pembahasan Latihan Soal & Jawaban

**1\. Pengertian RPL dan alasan pentingnya** RPL adalah disiplin ilmu dengan pendekatan sistematis, terukur, dan terorganisasi untuk membangun dan memelihara software. Hal ini penting diterapkan karena pengembangan aplikasi tidak sekadar menulis baris program, melainkan memastikan solusi tersebut tepat sasaran, berkualitas, serta meminimalisir masalah di kemudian hari.

**2\. Perbedaan Programming dan RPL beserta contoh** Programming adalah aktivitas menulis kode untuk membangun fungsi yang spesifik, misalnya membuat fitur filter pencarian data. Sedangkan RPL mencakup pengelolaan seluruh tahapan secara menyeluruh (dari analisis hingga pemeliharaan), contohnya membangun aplikasi pendataan gudang secara penuh hingga siap beroperasi.

**3\. Alasan analisis kebutuhan dilakukan sebelum coding** Agar pengembang mengetahui secara akurat siapa pengguna akhirnya, masalah apa yang sebenarnya ingin diselesaikan, serta batasan sistem yang harus dipenuhi. Langkah ini mencegah terjadinya aplikasi yang sudah jadi tapi fitur-fiturnya tidak bisa menyelesaikan masalah pengguna.

**4\. Perbandingan Waterfall, Prototype, dan Agile**

* **Waterfall** berjalan linear dan cocok untuk sistem berskala pasti seperti aplikasi perbankan.  
* **Prototype** berfokus membuat draf awal untuk dievaluasi, sangat cocok jika kebutuhan pengguna belum pasti, contohnya website *e-commerce* baru untuk UMKM.  
* **Agile** menggunakan metode *sprint* yang fleksibel, tepat digunakan untuk proyek seperti *game* yang harus terus mengikuti perkembangan tren dan teknologi.

**5\. Studi Kasus: Sistem Informasi Perpustakaan**

* **a. Masalah utama:** Sulit melacak jumlah buku fisik yang tersedia, repot mencatat keterlambatan pengembalian buku secara manual, dan kesulitan merekap laporan bulanan.  
* **b. Kebutuhan pengguna:** Sistem harus bisa mengelola sirkulasi buku (peminjam/pengembali), menampilkan stok buku terkini, mencatat denda/keterlambatan, serta mencetak laporan aktivitas secara otomatis.  
* **c. 5 Fitur utama yang diusulkan:** Fitur *Login*, Kelola Data Buku (Tambah/Ubah/Hapus), Pencatatan Transaksi Peminjaman & Pengembalian, Pengecekan Ketersediaan & Keterlambatan Buku, dan fitur Cetak Laporan Bulanan (PDF/Excel).
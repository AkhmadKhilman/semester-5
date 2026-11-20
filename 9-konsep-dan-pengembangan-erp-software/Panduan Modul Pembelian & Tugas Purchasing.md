# Panduan Lengkap Modul Pembelian (Buying) & Alur Kerja Purchasing

Dokumen ini berisi panduan mengenai antarmuka modul pembelian pada ERPNext dan rincian langkah kerja operasional bagi tim Purchasing (Pengadaan).

## A. Penjelasan Modul Buying (Pembelian)

Berdasarkan tampilan pada dasbor, Modul Buying adalah pusat kendali untuk mengelola arus belanja perusahaan, memilih vendor, dan memantau pesanan barang. Modul ini terbagi menjadi beberapa bagian utama:

*   **Dashboard & Analytics:** Bagian paling atas yang menampilkan grafik *Purchase Order Analysis*. Grafik donat ini membandingkan nominal pesanan pembelian yang belum ditagihkan (*Amount to Bill*) dengan pesanan yang tagihannya sudah masuk (*Billed Amount*). Terdapat juga metrik *Purchase Orders to Receive* sebagai pengingat jumlah pesanan yang barangnya belum tiba di gudang.
*   **Supplier (Pemasok/Vendor):** Modul ini berisi basis data seluruh pihak luar yang menjadi mitra perusahaan untuk berbelanja. Di dalamnya menyimpan profil lengkap pemasok, mulai dari nama perusahaan, alamat, kontak person, hingga pengelompokan jenis pemasok (*Supplier Group*).
*   **Material Request (Surat Permintaan Barang):** Modul ini menampung daftar keranjang belanja internal. Dokumen di sini biasanya merupakan kiriman/permintaan dari departemen Gudang yang melaporkan bahwa stok barang tertentu sudah menipis dan perlu segera dibeli.
*   **Request for Quotation & Supplier Quotation:** Modul ini digunakan jika perusahaan ingin melakukan "tender" atau meminta penawaran harga dari beberapa *Supplier* sekaligus untuk membandingkan mana yang memberikan harga termurah.
*   **Purchase Order (Pesanan Pembelian):** Modul ini berisi dokumen kontrak resmi dari perusahaan yang dikirimkan kepada pihak pemasok. Dokumen ini mengunci daftar barang yang dipesan, kuantitas, estimasi tanggal kedatangan, dan harga beli (modal) yang sudah disepakati.
*   **Purchase Receipt & Purchase Invoice:** Meskipun secara operasional dieksekusi oleh Gudang dan Finance, menu pintasan dokumen ini tetap ada di modul Buying agar tim Purchasing bisa memantau apakah pesanan mereka sudah diterima fisik barangnya dan apakah tagihannya sudah diproses.
*   **Reports & Setup:** Menu pendukung untuk mengatur grup vendor serta mencetak laporan krusial seperti *Purchase Analytics* (analisis belanja) dan *Supplier Quotation Comparison* (perbandingan harga antar vendor).

---

## B. Langkah Kerja Tim Purchasing dalam Simulasi ERP

Tim Purchasing bertugas merespons kebutuhan stok internal dan mencari harga modal terbaik dari pihak luar. Berikut adalah rincian tugas operasional mereka dari awal hingga akhir:

*   **1. Membuat Master Data Supplier (Data Induk Pemasok)**
    *   **Cara Pembuatan:** Staf pengadaan masuk ke dalam menu *Supplier* yang ada di panel kiri, kemudian membuat dokumen baru. Mereka harus melengkapi identitas pihak vendor, seperti nama perusahaan pemasok, alamat lengkap, narahubung, serta mengklasifikasikannya ke dalam *Supplier Group* yang relevan (misalnya grup pemasok bahan baku atau grup alat tulis kantor).
    *   **Fungsi Dokumen:** Dokumen ini berfungsi untuk mendaftarkan identitas mitra bisnis ke dalam basis data perusahaan. Dengan data yang sudah dipersiapkan (setup) secara valid di awal, pembuatan surat pesanan (*Purchase Order*) di kemudian hari akan berjalan jauh lebih cepat karena sistem otomatis menarik informasi alamat dan kontak yang tepat, mencegah risiko pesanan terkirim ke alamat vendor yang salah.

*   **2. Menindaklanjuti Material Request (Permintaan Barang)**
    *   **Cara Pembuatan:** Staf pengadaan membuka menu *Material Request* dan mencari dokumen berstatus "Pending" yang baru saja diajukan oleh departemen Gudang. Mereka memvalidasi jenis dan jumlah barang yang diminta. Jika disetujui, staf pengadaan mengklik tombol *Create* pada dokumen tersebut dan memilih opsi *Purchase Order* untuk mengubah status permintaan menjadi pesanan riil.
    *   **Fungsi Dokumen:** Proses ini berfungsi sebagai jembatan antara kebutuhan internal dan eksekusi eksternal. Dengan menarik data langsung dari *Material Request*, sistem ERP memastikan bahwa setiap sen uang yang dibelanjakan oleh bagian Pengadaan memiliki dasar kebutuhan yang sah (riil) dari pihak Gudang, mencegah terjadinya pembelian barang fiktif.

*   **3. Menerbitkan Purchase Order (Pesanan Pembelian)**
    *   **Cara Pembuatan:** Setelah *Material Request* ditarik menjadi dokumen *Purchase Order* baru, staf pengadaan diwajibkan untuk memilih nama *Supplier* mana yang akan diberikan proyek pesanan ini. Pada tahap inilah staf pengadaan memasukkan Harga Beli (*Rate*) ke dalam sistem berdasarkan hasil negosiasi terbaru dengan pemasok. Setelah harga dan jadwal pengiriman dipastikan benar, dokumen disahkan (*Submit*).
    *   **Fungsi Dokumen:** Dokumen ini berfungsi sebagai surat perintah kerja dan kontrak belanja yang sah secara hukum dengan pihak pemasok luar. Di dalam alur kerja ERP, dokumen ini sangat esensial karena ia mengunci Harga Pokok/Modal barang. Pengesahan dokumen ini juga akan memberikan "lampu hijau" kepada bagian Gudang untuk bersiap melakukan pencocokan saat kurir pemasok datang membawa barang tersebut (*Purchase Receipt*).
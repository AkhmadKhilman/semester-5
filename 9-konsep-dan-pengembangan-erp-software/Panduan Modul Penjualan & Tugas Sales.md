# Panduan Lengkap Modul Penjualan (Selling) & Alur Kerja Sales

Dokumen ini berisi panduan mengenai antarmuka modul penjualan pada ERPNext dan rincian langkah kerja operasional bagi tim Sales.

## A. Penjelasan Modul Selling (Penjualan)

Berdasarkan tampilan pada dasbor, Modul Selling merupakan pusat administrasi penjualan yang menaungi beberapa menu dan dokumen penting, antara lain:

*   **Dashboard & Analytics:** Bagian paling depan yang menampilkan grafik *Sales Order Analysis*. Grafik ini membandingkan nominal pesanan yang belum ditagih dengan yang sudah lunas, serta menampilkan metrik pengingat otomatis seperti jumlah pesanan yang harus segera dikirim (*Sales Orders to Deliver*) atau ditagih (*Sales Orders to Bill*).
*   **Customer (Pelanggan):** Modul ini berisi basis data seluruh pihak yang membeli barang dari perusahaan Anda. Di dalamnya menyimpan rincian informasi profil pelanggan secara lengkap, mulai dari nama entitas (perusahaan atau perorangan), alamat penagihan dan pengiriman, klasifikasi wilayah pemasaran (*Territory*), hingga pengelompokan segmen pasar (*Customer Group*).
*   **Quotation (Penawaran Harga):** Modul ini adalah tempat menyimpan seluruh riwayat proposal penawaran harga. Isi dari dokumen di dalam modul ini mencakup data pelanggan yang dituju, rincian barang (*Item*) beserta kuantitasnya, harga dasar yang ditarik otomatis dari *Price List*, serta rincian diskon atau kupon promo yang ditawarkan oleh tenaga penjual.
*   **Sales Order (Pesanan Penjualan):** Modul ini memuat daftar pesanan yang sudah disepakati (deal). Isi dokumennya merupakan salinan pasti dari *Quotation* yang telah disetujui, namun statusnya berubah menjadi kontrak mengikat yang mencantumkan jadwal pengiriman barang dan syarat pembayaran. 
*   **Delivery Note (Nota Pengiriman):** Modul ini berisi dokumen jalan yang digunakan sebagai bukti fisik pengiriman barang ke alamat pelanggan. Meskipun biasanya dieksekusi oleh bagian Gudang, dokumen ini tetap masuk di menu *Selling* karena berkaitan langsung dengan pemenuhan *Sales Order*.
*   **Items & Pricing:** Menu pendukung untuk mengelola daftar barang dagangan (*Item*), master harga dasar (*Price List*), dan aturan diskon (*Pricing Rule*).
*   **Setup & Reports:** Menu untuk melakukan pengaturan data awal seperti target *Sales Person*, serta mencetak laporan analitik riwayat penjualan perusahaan.

---

## B. Langkah Kerja Tim Sales dalam Simulasi ERP

Berikut adalah rincian tugas dan langkah kerja yang harus dieksekusi oleh tim Sales dari hulu ke hilir:

*   **1. Membuat Master Data Customer (Data Induk Pelanggan)**
    *   **Cara Pembuatan:** Tenaga penjual masuk ke dalam menu *Customer* yang ada di sisi kiri layar, kemudian membuat dokumen baru. Selanjutnya, tenaga penjual harus mengisi identitas pelanggan secara lengkap, seperti nama perusahaan pembeli, menentukan alamat penagihan yang valid, serta mengelompokkan pelanggan tersebut ke dalam kategori *Territory* (wilayah) dan *Customer Group* (segmen pasar) yang sesuai.
    *   **Fungsi Dokumen:** Dokumen ini berfungsi untuk mendaftarkan identitas pembeli ke dalam basis data pusat perusahaan. Penyiapan data yang valid sejak awal sangat krusial agar pada saat pembuatan dokumen transaksi di tahap selanjutnya, seluruh informasi pelanggan dapat ditarik secara otomatis oleh sistem. Hal ini sangat berguna untuk mencegah kesalahan pengetikan (*human error*) dan mempercepat proses administrasi.

*   **2. Membuat Quotation (Penawaran Harga)**
    *   **Cara Pembuatan:** Tenaga penjual membuka menu *Quotation* dan membuat dokumen baru. Langkah pertamanya adalah memilih nama *Customer* yang sebelumnya sudah didaftarkan, lalu memasukkan barang (*Item*) apa saja yang diminati oleh pembeli. Sistem kemudian akan secara otomatis memunculkan harga barang dari *Price List*. Jika diperlukan untuk merayu pelanggan, tenaga penjual dapat memasukkan kode kupon diskon di kolom yang tersedia sebelum dokumen disimpan.
    *   **Fungsi Dokumen:** Dokumen ini berfungsi sebagai proposal resmi tertulis yang diberikan kepada calon pembeli sebelum adanya kesepakatan final. Dalam konteks sistem ERP, tahap penawaran ini juga berfungsi sebagai kontrol keamanan internal yang ketat. Sistem mengunci kolom Harga Dasar menjadi hanya-baca (*Read-Only*) agar tenaga penjual tidak bisa sembarangan memanipulasi atau menurunkan harga modal di luar wewenang diskon yang diberikan manajemen.

*   **3. Melanjutkan Quotation menjadi Sales Order (Pesanan Penjualan)**
    *   **Cara Pembuatan:** Setelah pelanggan menyatakan setuju dengan harga di penawaran, tenaga penjual membuka kembali dokumen *Quotation* tersebut. Dari dalam dokumen itu, tenaga penjual cukup mengklik tombol *Create* (Buat) di pojok kanan atas dan memilih opsi *Sales Order*. Seluruh data barang, harga, dan diskon akan tersalin secara otomatis menjadi dokumen baru yang siap untuk di-*Submit* (disahkan) tanpa perlu pengetikan ulang sedikit pun.
    *   **Fungsi Dokumen:** Dokumen ini berfungsi sebagai kontrak kesepakatan (*deal*) yang sah dan mengikat antara perusahaan dengan pelanggan. Di dalam alur kerja ERP, dokumen *Sales Order* adalah motor penggerak utama bagi departemen lain. Begitu disahkan, dokumen ini langsung mengirimkan sinyal ke bagian Gudang untuk segera menyiapkan barang dan pengiriman (*Delivery Note*), sekaligus memberi tahu bagian Finance bahwa ada tagihan piutang yang siap dicetak (*Sales Invoice*).
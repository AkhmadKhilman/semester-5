# Rangkuman Mata Kuliah: Konsep dan Pengembangan ERP Software

## Pertemuan 1: Pengantar dan Evolusi ERP
* **Masalah Sistem Lama:** Sebelum ada ERP, sistem di tiap departemen terpisah-pisah, sehingga departemen menyimpan versi data sendiri dan laporan sulit disusun secara real-time. Akibatnya, sering terjadi perbedaan data, misalnya stok di gudang berbeda dengan catatan penjualan.
* **Solusi ERP:** ERP hadir untuk mengintegrasikan proses inti perusahaan ke dalam satu platform dengan basis data bersama (master data), sehingga satu transaksi akan otomatis memperbarui fungsi terkait lainnya.
* **Evolusi Sistem:** Berawal dari IC (hanya untuk persediaan), berkembang menjadi MRP (untuk menghitung kebutuhan material), lalu MRP II (mencakup kapasitas dan sumber daya), hingga menjadi ERP modern yang kini juga memanfaatkan cloud, mobile, dan AI.
* **Penerapan (BPR):** Sebelum memakai ERP, perusahaan biasanya melakukan Business Process Reengineering (BPR) untuk merancang proses bisnis yang lebih baik agar sistem baru tidak sekadar meniru proses lama yang buruk.

## Pertemuan 2: Arsitektur Teknis dan Teknologi Pendukung ERP
* **Three-tier Architecture:** Sistem ERP dibagi menjadi tiga lapisan: Presentation layer (tampilan antarmuka untuk pengguna), Application layer (untuk menjalankan aturan dan proses bisnis), dan Data layer (untuk menyimpan transaksi dan data).
* **Database Terpusat:** Data dikelola sebagai Single Source of Truth (SSOT) agar menjadi rujukan resmi dan seragam bagi seluruh modul.
* **Integrasi Aplikasi:** Pertukaran data dan fungsi dengan aplikasi lain dilakukan menggunakan Middleware, API, atau pendekatan Service Oriented Architecture (SOA).
* **Infrastruktur dan Keamanan:** Sistem bisa ditempatkan on-premise (server sendiri), cloud (server penyedia), atau hybrid. Untuk keamanan, ERP menggunakan Role-Based Access Control (RBAC) yang membatasi hak akses pengguna berdasarkan peran kerjanya.

## Pertemuan 3: Manajemen Rantai Pasok (Supply Chain Management)
* **Konsep SCM:** SCM adalah pengelolaan aliran barang, aliran informasi, dan aliran keuangan mulai dari pemasok sampai ke tangan pelanggan. ERP menyatukan data dari ketiga aliran ini.
* **Perencanaan Kebutuhan:** Perusahaan menghitung kebutuhan material menggunakan metode MRP yang didasarkan pada Bill of Materials (BOM), yaitu daftar komponen pembentuk suatu produk.
* **Proses Pembelian:** Diawali dengan unit kerja membuat Purchase Requisition (PR) sebagai permintaan internal. Setelah disetujui, perusahaan akan menerbitkan Purchase Order (PO) yang merupakan dokumen pesanan resmi kepada pemasok.
* **Manajemen Persediaan dan Gudang:** ERP mengatur safety stock (stok pengaman) dan batas stok agar barang tidak kurang atau berlebih. Aktivitas fisik di dalam gudang sendiri diatur secara rinci oleh Warehouse Management System (WMS).
* **Prediksi dan Retur:** Sistem ERP juga digunakan untuk demand forecasting (meramal permintaan pelanggan ke depan) serta reverse logistics yang mengelola retur barang dari pelanggan yang rusak atau tidak sesuai.

## Tambahan: Alur Kerja ERP Mini (Siklus Pesanan hingga Pembayaran)
* **Sales (Penjualan):** Menerima pesanan dari pelanggan dan menginputnya ke dalam sistem sebagai *Sales Order*. Informasi pesanan ini otomatis masuk dan bisa dilihat oleh bagian gudang secara real-time.
* **Warehouse (Gudang - Pengecekan):** Mengecek ketersediaan barang di sistem. Jika ternyata stok habis, bagian gudang akan membuat permintaan pembelian (*Purchase Requisition*) yang diteruskan ke bagian pembelian.
* **Purchase (Pembelian):** Menerima permintaan dari gudang, lalu mencarikan barang dan menerbitkan pesanan resmi (*Purchase Order*) kepada pemasok (supplier).
* **Warehouse (Penerimaan & Pengiriman):** Saat barang dari pemasok tiba, gudang mencatat penerimaan barang di sistem sehingga stok otomatis bertambah. Setelah itu, gudang langsung menyiapkan dan mengirimkan barang tersebut kepada pelanggan sesuai pesanan awal dari tim Sales.
* **Finance (Keuangan):** Menyelesaikan urusan perputaran uang. Pertama, Finance menerima tagihan dari pemasok dan melakukan pembayaran (*Account Payable*). Kedua, Finance menerbitkan faktur tagihan (*Invoice*) kepada pelanggan dan mencatat uang masuk dari pelanggan (*Account Receivable*).
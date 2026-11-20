# Panduan Lengkap Modul Gudang (Stock) & Alur Kerja Warehouse

Dokumen ini berisi panduan mengenai antarmuka modul gudang (Stock) pada ERPNext dan rincian langkah kerja operasional bagi tim Warehouse.

## A. Penjelasan Modul Stock (Gudang)

Berdasarkan tampilan dasbor, Modul Stock adalah pusat pengelolaan inventaris dan lalu lintas barang perusahaan, yang terdiri dari beberapa bagian utama:

*   **Dashboard & Analytics:** Menampilkan metrik krusial secara langsung, seperti *Total Stock Value* (total nilai aset barang di gudang), *Total Warehouses* (jumlah fisik gudang yang dikelola), dan *Total Active Items* (jumlah varian barang yang aktif diperjualbelikan).
*   **Item (Barang):** Modul ini adalah *database* induk yang menyimpan profil seluruh barang perusahaan. Di sini tersimpan nama barang, kode *barcode*, satuan ukur (*Unit of Measure*/UOM), hingga pengaturan batas minimum stok agar sistem bisa memberi peringatan saat barang mau habis.
*   **Warehouse (Gudang Fisik):** Modul ini digunakan untuk mendaftarkan dan memetakan struktur gudang perusahaan. Gudang bisa dibuat bertingkat (hierarki), misalnya memisahkan antara gudang bahan baku, ruang produksi, hingga gudang toko siap jual.
*   **Material Request (Surat Permintaan Barang):** Modul ini berisi dokumen internal yang digunakan oleh staf gudang untuk meminta pengadaan barang baru ketika stok menipis. Dokumen ini nantinya akan dilempar ke departemen Purchasing untuk dieksekusi menjadi pesanan pembelian.
*   **Stock Entry (Pergerakan Internal):** Modul ini mengelola dokumen untuk pergerakan barang yang *tidak* melibatkan pihak luar. Fungsinya mencakup pemindahan barang antar-gudang (Material Transfer), pengeluaran barang untuk pemakaian sendiri (Material Issue), atau penambahan stok awal.
*   **Stock Reconciliation (Penyesuaian Stok):** Modul ini digunakan khusus saat melakukan *stock opname* (penghitungan fisik berkala). Jika ada perbedaan antara jumlah barang di komputer dengan jumlah fisik di lapangan, dokumen inilah yang dipakai untuk menyamakan angkanya.
*   **Purchase Receipt (Penerimaan Barang):** Modul ini berisi dokumen sah yang membuktikan bahwa barang dari pihak *Supplier* telah tiba secara fisik dan resmi masuk menambah jumlah stok di sistem.
*   **Delivery Note (Nota Pengiriman):** Modul ini berisi dokumen jalan yang digunakan saat barang dikeluarkan dari gudang untuk dikirim kepada pelanggan pembeli. Mengesahkan dokumen ini akan otomatis mengurangi jumlah stok di sistem.
*   **Reports (Laporan):** Menyediakan laporan komprehensif seperti *Stock Ledger* (buku besar mutasi barang), *Stock Balance* (sisa stok saat ini), dan *Stock Ageing* (laporan barang usang yang menumpuk terlalu lama).

---

## B. Langkah Kerja Tim Warehouse dalam Simulasi ERP

Tim Warehouse memegang peranan vital dalam menjaga keakuratan fisik barang. Berikut adalah rincian tugas operasional mereka dari awal hingga akhir:

*   **1. Membuat Master Data Item dan Warehouse (Data Induk Barang dan Gudang)**
    *   **Cara Pembuatan:** Staf gudang masuk ke menu *Warehouse* untuk mendaftarkan lokasi-lokasi penyimpanan yang ada. Selanjutnya, mereka masuk ke menu *Item* untuk mendaftarkan satu per satu barang yang akan dikelola. Mereka harus mengisi kode barang, nama, deskripsi, dan satuan pengukurannya (misal: Pcs, Kg, Box).
    *   **Fungsi Dokumen:** Dokumen-dokumen ini membentuk fondasi dasar sistem inventaris. Dengan mendaftarkan *Item* dan lokasi *Warehouse* sejak awal, seluruh departemen lain (Sales, Purchasing, Finance) dapat memilih barang yang seragam dari satu *database* yang sama, memastikan tidak ada barang fiktif atau lokasi gudang yang tidak valid dalam transaksi harian.

*   **2. Mengajukan Material Request (Permintaan Pengadaan Barang)**
    *   **Cara Pembuatan:** Ketika staf gudang memantau laporan *Stock Balance* dan melihat ada barang yang mau habis, mereka membuka menu *Material Request*. Mereka membuat dokumen baru dengan tipe "Purchase", memasukkan daftar barang yang menipis beserta jumlah yang dibutuhkan, lalu mengesahkan (*Submit*) dokumen tersebut.
    *   **Fungsi Dokumen:** Dokumen ini berfungsi sebagai surat permohonan resmi dari pihak operasional (Gudang) kepada pihak pengadaan (Purchasing). Melalui dokumen ini, sistem ERP memastikan bahwa bagian Purchasing tidak sembarangan membeli barang tanpa ada kebutuhan riil dan persetujuan dari pihak yang menjaga stok.

*   **3. Mengeksekusi Purchase Receipt (Menerima Barang dari Supplier)**
    *   **Cara Pembuatan:** Saat kurir dari *Supplier* datang membawa barang hasil pesanan Purchasing, staf gudang membuka dokumen *Purchase Order* yang bersangkutan. Mereka lalu mengklik tombol *Create* dan memilih *Purchase Receipt*. Staf wajib menghitung fisik barang yang datang; jika jumlahnya sesuai, mereka langsung menyimpan dan mengesahkan (*Submit*) dokumen tersebut.
    *   **Fungsi Dokumen:** Pengesahan dokumen ini adalah momen kritis di mana jumlah barang di dalam sistem ERP secara otomatis bertambah (Stok Masuk). Dokumen ini juga bertindak sebagai bukti terima (tanda tangan elektronik) yang akan digunakan oleh departemen Finance sebagai dasar bahwa tagihan dari *Supplier* sudah boleh dibayar karena barangnya memang benar-benar sudah diterima.

*   **4. Mengeksekusi Delivery Note (Mengirim Barang ke Customer)**
    *   **Cara Pembuatan:** Setelah tim Sales berhasil *deal* dan menghasilkan *Sales Order*, staf gudang akan membuka dokumen pesanan tersebut. Untuk memproses pengiriman, mereka mengklik tombol *Create* dan memilih *Delivery Note*. Staf gudang kemudian melakukan *packing* barang, mencetak dokumen jalan ini, lalu mengesahkannya di sistem saat barang berangkat bersama kurir.
    *   **Fungsi Dokumen:** Pengesahan dokumen ini adalah momen di mana jumlah barang di dalam sistem ERP secara otomatis berkurang (Stok Keluar). Dokumen bukti kirim ini sangat penting karena akan dijadikan landasan oleh departemen Finance untuk menerbitkan faktur tagihan (*Sales Invoice*) kepada pelanggan, memastikan pelanggan hanya ditagih untuk barang yang sudah benar-benar dikirim dari gudang.
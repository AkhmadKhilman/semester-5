# Panduan Operasional & Siklus ERPNext (Simulasi Ritel)

Dokumen ini adalah panduan lengkap fungsionalitas modul dan alur kerja (workflow) pada sistem ERPNext, disusun khusus untuk kelancaran presentasi simulasi perusahaan ritel.

---

## 1. Modul SELLING (Dikelola oleh Tim Sales)
Modul ini adalah garis depan operasional perusahaan yang berinteraksi langsung dengan pelanggan dan mendatangkan pendapatan. 

*   **Customer (Pelanggan):** Menu untuk mendaftarkan dan menyimpan master data pelanggan. Di sini Sales mencatat informasi dasar, kontak, dan alamat pelanggan sebelum transaksi dimulai.
*   **Quotation (Penawaran Harga):** Dokumen draf resmi yang diberikan kepada pelanggan berisi rincian barang, harga, dan estimasi pajak. Di dokumen inilah Sales bisa melakukan promosi dengan memasukkan **Coupon Code** yang tertaut pada *Pricing Rule* untuk memberikan diskon (memotong nilai *Net Total*).
*   **Sales Order (Pesanan Penjualan):** Dokumen yang dibuat (ditarik dari Quotation) saat pelanggan sudah "Deal" atau setuju dengan penawaran. Dokumen ini mengunci kesepakatan harga dan menjadi lampu hijau bagi departemen lain (Gudang dan Keuangan) untuk mulai bekerja melayani pesanan tersebut.

## 2. Modul BUYING (Dikelola oleh Tim Purchasing)
Modul ini bertanggung jawab menjaga ketersediaan barang untuk dijual dengan mencari harga beli (modal) yang paling efisien dari berbagai pihak luar.

*   **Supplier (Pemasok):** Menu untuk mendaftarkan mitra bisnis tempat perusahaan berbelanja.
*   **Material Request (Permintaan Barang):** Dokumen keranjang belanja internal. Purchasing menerima dokumen ini dari pihak Gudang yang melaporkan bahwa stok barang tertentu sudah menipis dan perlu segera dibeli.
*   **Purchase Order (Pesanan Pembelian):** Dokumen pemesanan resmi yang dikirim ke Supplier. Di sini, Purchasing memiliki wewenang penuh untuk memasukkan dan mengubah harga beli (*Rate*) berdasarkan hasil negosiasi terbaru dengan pihak supplier. Dokumen ini juga bisa dipecah untuk beberapa supplier yang berbeda dari satu Material Request yang sama.

## 3. Modul STOCK (Dikelola oleh Tim Warehouse)
Modul ini adalah jantung operasional logistik. Bertugas mengelola perpindahan fisik barang, melacak ketersediaan stok, dan menjaga nilai aset gudang.

*   **Item & Warehouse:** Menu untuk mengelola master data barang (Item) yang dijual beserta harganya (Item Price), dan lokasi penyimpanan fisiknya (Warehouse).
*   **Material Request (MR):** Orang gudang bertugas memicu pengadaan barang dengan membuat dokumen MR bertipe "Purchase" ketika melihat indikator stok menipis, lalu menyerahkannya ke tim Purchasing.
*   **Purchase Receipt (Penerimaan Barang):** Dokumen wajib yang dibuat saat kurir supplier datang membawa barang. Gudang mencocokkan fisik barang dengan *Purchase Order*. ERPNext akan memproses penambahan stok di sistem secara otomatis (bahkan melalui *background job* jika dari multi-supplier).
*   **Delivery Note (Surat Jalan/Pengiriman):** Dokumen yang dibuat saat barang dipacking dan dikeluarkan dari gudang untuk dikirim ke pelanggan, berdasarkan *Sales Order*. Dokumen ini memotong jumlah stok barang di sistem secara real-time.

## 4. Modul ACCOUNTING (Dikelola oleh Tim Finance)
Modul ini adalah muara dari seluruh aktivitas sistem. Mengatur pencatatan utang, piutang, perpajakan, pencairan dana, dan pelaporan laba rugi.

*   **Purchase Invoice (Faktur Pembelian):** Dokumen penagihan/utang perusahaan kepada supplier. Ditarik dari dokumen *Purchase Receipt*. Di tahap ini Finance memvalidasi besaran PPN 11% masukan sebelum tagihan disahkan.
*   **Sales Invoice (Faktur Penjualan):** Dokumen tagihan/piutang kepada pelanggan. Ditarik dari dokumen *Delivery Note*. Dokumen ini mengunci total nilai yang harus ditransfer oleh pelanggan ke rekening perusahaan.
*   **Payment Entry (Pencatatan Pembayaran):** Menu kasir perusahaan. Digunakan untuk mencatat uang masuk dari *Sales Invoice* maupun uang keluar untuk *Purchase Invoice*. Finance wajib mencatat **Reference No** (nomor struk/bukti transfer bank) sebagai jejak audit (*audit trail*).
*   **Profit and Loss (Laba Rugi):** Laporan akhir (dashboard) yang menyajikan selisih antara seluruh pendapatan (uang masuk dari jualan) dan beban operasional/HPP (uang keluar untuk modal), membuktikan kinerja finansial perusahaan.

---

## SIKLUS ERP PERUSAHAAN (Alur Kerja Menyeluruh)

Untuk memahami bagaimana sistem terintegrasi, mari kita lihat dua siklus utama bisnis:

### A. Siklus Belanja (Procure-to-Pay)
*Fokus: Bagaimana perusahaan mengeluarkan uang untuk restock barang.*
1.  **Gudang** melihat stok kosong -> Membuat **Material Request**.
2.  **Purchasing** memproses permintaan -> Memilih supplier dan menerbitkan **Purchase Order** (beserta kesepakatan harga belinya).
3.  **Supplier** mengirim fisik barang ke perusahaan.
4.  **Gudang** menerima barang tersebut -> Menerbitkan **Purchase Receipt** (Stok di sistem bertambah).
5.  **Finance** menerima nota tagihan dari supplier -> Membuat **Purchase Invoice**.
6.  **Finance** melunasi utang ke supplier via bank -> Mencatat pelunasan di **Payment Entry**. (Siklus selesai).

### B. Siklus Jualan (Order-to-Cash)
*Fokus: Bagaimana perusahaan melayani pelanggan hingga menerima uang masuk.*
1.  **Sales** merayu prospek -> Membuat **Quotation** (lengkap dengan diskon kupon).
2.  **Pelanggan** setuju dengan harga -> Sales mengubahnya menjadi **Sales Order**.
3.  **Gudang** menyiapkan barang sesuai pesanan -> Membuat **Delivery Note** dan mengirimnya ke pelanggan (Stok di sistem berkurang).
4.  **Finance** mencetak tagihan untuk pelanggan -> Membuat **Sales Invoice** (lengkap dengan perhitungan PPN 11% dari Net Total).
5.  **Pelanggan** melakukan transfer pembayaran.
6.  **Finance** mengecek mutasi rekening -> Mencatat penerimaan uang di **Payment Entry** beserta nomor referensi transaksinya. (Siklus selesai).
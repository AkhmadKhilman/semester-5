# Panduan Lengkap Modul Keuangan (Accounting) & Alur Kerja Finance

Dokumen ini berisi panduan mengenai antarmuka modul keuangan pada ERPNext dan rincian langkah kerja operasional bagi tim Finance.

## A. Penjelasan Modul Accounting (Keuangan)

Berdasarkan tampilan dasbor, Modul Accounting adalah pusat pelaporan pajak, pencatatan transaksi uang, dan muara akhir dari seluruh aktivitas perusahaan. Modul ini terbagi menjadi beberapa bagian utama:

*   **Dashboard & Analytics:** Menampilkan grafik balok *Profit and Loss* yang membandingkan *Income* (pendapatan) dengan *Expense* (pengeluaran) untuk menghasilkan *Net Profit/Loss* (laba bersih). Di bawahnya terdapat metrik pantauan kas cepat, yaitu *Total Outgoing Bills* (total tagihan yang harus dibayar ke supplier), *Total Incoming Bills* (total tagihan yang harus ditagih ke pelanggan), dan *Total Incoming Payment* (total uang riil yang sudah diterima).
*   **Chart of Accounts (Bagan Akun):** Modul ini adalah fondasi akuntansi yang menyimpan daftar seluruh rekening pembukuan perusahaan (seperti Kas, Piutang, HPP, Pendapatan, dan Beban). Semua transaksi di ERP akan bermuara pada akun-akun ini.
*   **Purchase Invoice (Faktur Pembelian):** Modul ini memuat daftar tagihan utang sah yang diterbitkan berdasarkan barang yang sudah diterima dari pihak *Supplier*.
*   **Sales Invoice (Faktur Penjualan):** Modul ini memuat daftar faktur piutang sah yang diterbitkan untuk menagih pembayaran kepada pihak *Customer* berdasarkan pesanan yang sudah dikirim.
*   **Payment Entry (Entri Pembayaran):** Modul ini adalah tempat mengeksekusi perpindahan uang riil. Digunakan baik untuk mencatat uang keluar (melunasi *Purchase Invoice*) maupun mencatat uang masuk (menerima pelunasan *Sales Invoice*).
*   **Journal Entry (Jurnal Umum):** Modul ini digunakan untuk mencatat transaksi keuangan manual yang tidak terkait langsung dengan jual-beli barang, seperti mencatat biaya gaji karyawan, biaya listrik, atau penyusutan aset.
*   **Taxes (Pajak):** Modul khusus untuk mengatur templat pajak, seperti Pajak Pertambahan Nilai (PPN 11%), agar perhitungan pajak dapat dimasukkan secara otomatis ke dalam setiap faktur tagihan.
*   **Reports (Laporan):** Menyediakan dokumen rekapitulasi keuangan berstandar akuntansi resmi, seperti *General Ledger* (Buku Besar), *Balance Sheet* (Neraca Keuangan), dan *Cash Flow* (Arus Kas).

---

## B. Langkah Kerja Tim Finance dalam Simulasi ERP

Tim Finance bertugas menuntaskan siklus dokumen dari departemen lain menjadi aliran uang dan laporan keuangan. Berikut adalah rincian tugas operasional mereka:

*   **1. Membuat Setup Pajak (Taxes Template)**
    *   **Cara Pembuatan:** Staf Finance masuk ke menu *Sales Tax Template* atau *Purchase Tax Template* di bawah kategori *Taxes*. Mereka membuat templat baru, misalnya bernama "PPN 11%". Di dalamnya, mereka memilih akun penampung pajak (seperti Hutang PPN atau PPN Masukan) dan memasukkan persentase tarif pajak sebesar 11%.
    *   **Fungsi Dokumen:** Setup templat ini sangat esensial agar perhitungan pajak terstandarisasi. Dengan adanya templat ini, staf tidak perlu menghitung pajak secara manual pakai kalkulator setiap kali menerbitkan faktur. Sistem akan otomatis menambahkan nilai pajak 11% dari total dasar pengenaan pajak pada setiap transaksi yang menggunakan templat tersebut.

*   **2. Menerbitkan Purchase Invoice (Mencatat Utang ke Supplier)**
    *   **Cara Pembuatan:** Setelah tim Gudang mengesahkan penerimaan barang (*Purchase Receipt*), staf Finance membuka dokumen penerimaan tersebut. Mereka mengklik tombol *Create* dan memilih opsi *Purchase Invoice*. Staf Finance kemudian memasukkan nomor faktur pajak dari supplier (jika ada), memeriksa kesesuaian harga dan pajak, lalu mengesahkan (*Submit*) dokumennya.
    *   **Fungsi Dokumen:** Dokumen ini berfungsi sebagai pengakuan utang resmi perusahaan kepada pihak *Supplier*. Pengesahan dokumen ini secara otomatis akan memperbarui pembukuan, yaitu menambah saldo di akun Hutang Usaha (Accounts Payable) dan mencatat Pajak Masukan.

*   **3. Menerbitkan Sales Invoice (Menagih Piutang ke Customer)**
    *   **Cara Pembuatan:** Setelah tim Gudang mengesahkan pengiriman barang (*Delivery Note*), staf Finance membuka dokumen pesanan (*Sales Order*) yang bersangkutan. Mereka mengklik tombol *Create* dan memilih opsi *Sales Invoice*. Staf Finance memastikan templat pajak penjualan (PPN Keluaran) sudah terpasang, mengatur tanggal jatuh tempo pembayaran, lalu mengesahkan (*Submit*) dokumen tersebut untuk dicetak dan dikirim ke pelanggan.
    *   **Fungsi Dokumen:** Dokumen ini berfungsi sebagai faktur penagihan resmi yang sah secara hukum dan pajak. Saat dokumen ini disahkan, sistem ERP otomatis membukukan pendapatan (*Revenue*) dan mencatat saldo pada akun Piutang Usaha (*Accounts Receivable*) atas nama pelanggan tersebut.

*   **4. Mengeksekusi Payment Entry (Pencatatan Uang Masuk/Keluar)**
    *   **Cara Pembuatan:** Saat jatuh tempo tiba dan perusahaan mentransfer uang ke *Supplier* (atau menerima transfer dari *Customer*), staf Finance membuka dokumen *Invoice* yang bersangkutan. Mereka mengklik tombol *Create* lalu memilih *Payment*. Staf wajib memilih akun Bank yang digunakan untuk transaksi, mengisi nomor referensi bukti transfer (misalnya nomor struk M-Banking) di kolom *Reference No*, lalu mengesahkannya.
    *   **Fungsi Dokumen:** Dokumen *Payment Entry* adalah langkah pamungkas yang menutup seluruh siklus transaksi. Fungsinya adalah mengubah status dokumen *Invoice* dari "Unpaid" menjadi "Paid", sekaligus memperbarui saldo rekening bank fisik perusahaan di dalam laporan keuangan. Pengisian Nomor Referensi sangat penting sebagai jejak audit (*audit trail*) untuk pencocokan rekening koran di akhir bulan.
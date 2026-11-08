# Rangkuman Materi: Aplikasi Mobile (Flutter)

## 1\. Konsep Dasar UI & Layouting (Fokus Live Coding)

* **Pemahaman `StatelessWidget`:** Pahami cara menyusun antarmuka yang bersifat statis tanpa menggunakan manajemen *state* atau `setState()`.  
* **Konfigurasi `MaterialApp` & Tema:** Pelajari cara mengatur *root* aplikasi, seperti menyembunyikan *banner debug*, serta mengaktifkan tema Material 3 (termasuk *color scheme* dari *seed*, mode gelap, dan mengikuti tema sistem).  
* **Struktur Scaffold & Tata Letak Dasar:** Pahami urutan pembungkus *widget* agar tampilan tidak *overflow* di layar kecil, yaitu kombinasi `SafeArea` \-\> `SingleChildScrollView` \-\> `Padding` \-\> `Column`.  
* **Elemen Antarmuka:** Pelajari penggunaan `AppBar` (termasuk penempatan `IconButton` di *actions*), penyisipan spasi dengan `SizedBox`, serta pembuatan profil dengan `CircleAvatar`.  
* **Tipografi Terpusat:** Pahami cara memanggil gaya teks bawaan dari *context* (misalnya `Theme.of(context).textTheme.headlineSmall` atau `bodyMedium`) ketimbang mengatur ukuran teks secara manual.  
* **Pembuatan Widget Kustom (*Reusable*):** Pelajari konsep membungkus *widget* (seperti `ListTile`) menjadi komponen mandiri yang dinamis dengan melempar nilai melalui parameter konstruktor (*icon*, *label*, *value*).  
* **Memicu Notifikasi Bawah (*SnackBar*):** Pahami cara menggunakan `ScaffoldMessenger.of(context).showSnackBar()` yang dipanggil lewat aksi klik pada `FloatingActionButton`.

## 2\. Form, Navigasi & State Management (Fokus Proyek Individu)

* **Navigasi (*Routing*) Transaksional:** Pahami logika perpindahan halaman, khususnya penggunaan `Navigator.pushReplacement` setelah proses *login* agar pengguna tidak bisa kembali ke halaman *login* dengan tombol *back*.  
* **Pengiriman & Pengembalian Data:** Pelajari cara mengirim data yang sudah diedit dari satu halaman kembali ke halaman utamanya menggunakan `Navigator.pop(context, dataBaru)`.  
* **Formulir & Validasi Input:** Pahami implementasi `Form` dan `TextFormField`. Ini mencakup cara menetapkan nilai awal menggunakan `TextEditingController` serta merancang fungsi *validator* (misal: mengecek kekosongan, memvalidasi format email, dan membatasi input nomor HP).  
* **Penggunaan `StatefulWidget` & `setState()`:** Pelajari bagaimana memperbarui tampilan layar utama secara instan (*real-time*) setelah menerima data baru dari proses *edit*.  
* **Manajemen Aset Lokal:** Pahami cara meregistrasikan direktori folder (seperti `assets/images/`) di dalam file `pubspec.yaml` agar aplikasi bisa merender gambar lokal.  
* **Menampilkan Daftar Berurut:** Pahami penggunaan `ListView` untuk menyajikan banyak *item* data (seperti daftar riwayat pendidikan atau hobi) agar bisa di-*scroll* secara independen.
# Catatan Lengkap: Penyelesaian Studi Kasus Metode Fuzzy Tsukamoto

## Studi Kasus (Soal Cerita)

Sebuah usaha kuliner rumahan setiap pagi harus menentukan berapa porsi makanan yang perlu diproduksi agar tidak mengalami kekurangan stok saat pesanan ramai maupun kerugian akibat makanan sisa yang tidak terjual. Berdasarkan pengalaman pengelola, keputusan jumlah produksi harian dipengaruhi oleh dua faktor utama, yaitu sisa stok **PERSEDIAAN** bahan/porsi yang masih ada di dapur dan perkiraan jumlah **PERMINTAAN** pelanggan pada hari tersebut.

Pengelola menetapkan batasan operasional harian sebagai berikut:
* **PERSEDIAAN** berada pada rentang `12` hingga `88` porsi.
* **PERMINTAAN** berada pada rentang `14` hingga `92` porsi.
* **JUMLAH PRODUKSI** yang dihasilkan berada pada rentang `120` hingga `640` porsi.

Pada suatu hari, bagian dapur mencatat bahwa sisa **PERSEDIAAN = 46 porsi**, sementara perkiraan **PERMINTAAN = 68 porsi**. Berapa porsi **JUMLAH PRODUKSI** yang harus dibuat hari itu menggunakan sistem inferensi **Metode Fuzzy Tsukamoto**?

---

## 1. Tabel Variabel, Himpunan, dan Batasan

### A. Tabel Variabel Input
| Jenis Variabel | Nama Variabel | Semesta Pembicaraan | Nama Himpunan | Bentuk Kurva | Domain / Batasan Himpunan |
| :--- | :--- | :---: | :--- | :--- | :---: |
| **Input 1** | **PERSEDIAAN** | `12 – 88` | **SEDIKIT** | Turun Linear | `12 – 22 – 54` |
| **Input 1** | **PERSEDIAAN** | `12 – 88` | **SEDANG** | Segitiga | `34 – 51 – 68` |
| **Input 1** | **PERSEDIAAN** | `12 – 88` | **BANYAK** | Naik Linear | `52 – 84 – 88` |
| **Input 2** | **PERMINTAAN** | `14 – 92` | **SEPI** | Turun Linear | `14 – 23 – 47` |
| **Input 2** | **PERMINTAAN** | `14 – 92` | **NORMAL** | Trapesium | `31 – 44 – 63 – 77` |
| **Input 2** | **PERMINTAAN** | `14 – 92` | **RAMAI** | Naik Linear | `59 – 86 – 92` |

### B. Tabel Variabel Output *(Terpisah — Wajib Monoton untuk Metode Tsukamoto)*
| Jenis Variabel | Nama Variabel | Semesta Pembicaraan | Nama Himpunan | Bentuk Kurva | Domain / Batasan Himpunan |
| :--- | :--- | :---: | :--- | :--- | :---: |
| **Output** | **JUMLAH PRODUKSI** | `120 – 640` | **SEDIKIT** | Monoton Turun | `120 – 180 – 540` |
| **Output** | **JUMLAH PRODUKSI** | `120 – 640` | **BANYAK** | Monoton Naik | `220 – 580 – 640` |

---

## 2. Pembuatan Grafik Kurva Tiap Variabel

### A. Kurva Variabel Input 1: PERSEDIAAN (`12 – 88`)

**1. Kurva Himpunan SEDIKIT (Turun Linear: `12 – 22 – 54`)**
```text
μ(x)
 1.0 |-------\
     |        \
 0.5 |         \
     |          \
 0.0 +-----------\-----------------> PERSEDIAAN
    12      22   54               88
```

**2. Kurva Himpunan SEDANG (Segitiga: `34 – 51 – 68`)**
```text
μ(x)
 1.0 |           /\
     |          /  \
 0.5 |         /    \
     |        /      \
 0.0 +-------/--------\------------> PERSEDIAAN
    12      34   51   68          88
```

**3. Kurva Himpunan BANYAK (Naik Linear: `52 – 84 – 88`)**
```text
μ(x)
 1.0 |                    /-------
     |                   /
 0.5 |                  /
     |                 /
 0.0 +----------------/------------> PERSEDIAAN
    12               52  84       88
```

**4. Kurva Gabungan Variabel PERSEDIAAN (3 Himpunan)**
```text
μ(x)   [SEDIKIT]    [SEDANG]      [BANYAK]
 1.0 |-------\         /\           /-------
     |        \       /  \         /
 0.5 |         \     /    \       /
     |          \   /      \     /
 0.0 +-----------\-/--------\---/----------> PERSEDIAAN
    12      22   34 51  52  54 68  84     88
```

---

### B. Kurva Variabel Input 2: PERMINTAAN (`14 – 92`)

**1. Kurva Himpunan SEPI (Turun Linear: `14 – 23 – 47`)**
```text
μ(x)
 1.0 |-------\
     |        \
 0.5 |         \
     |          \
 0.0 +-----------\-----------------> PERMINTAAN
    14      23   47               92
```

**2. Kurva Himpunan NORMAL (Trapesium: `31 – 44 – 63 – 77`)**
```text
μ(x)
 1.0 |          /--------\
     |         /          \
 0.5 |        /            \
     |       /              \
 0.0 +------/----------------\-----> PERMINTAAN
    14     31  44        63  77   92
```

**3. Kurva Himpunan RAMAI (Naik Linear: `59 – 86 – 92`)**
```text
μ(x)
 1.0 |                     /-------
     |                    /
 0.5 |                   /
     |                  /
 0.0 +-----------------/-----------> PERMINTAAN
    14                59  86      92
```

**4. Kurva Gabungan Variabel PERMINTAAN (3 Himpunan)**
```text
μ(x)    [SEPI]       [NORMAL]        [RAMAI]
 1.0 |-------\      /--------\        /-------
     |        \    /          \      /
 0.5 |         \  /            \    /
     |          \/              \  /
 0.0 +----------/\---------------\/----------> PERMINTAAN
    14     23  31 44   47   59  63 77  86   92
```

---

### C. Kurva Variabel Output: JUMLAH PRODUKSI (`120 – 640`)

**1. Kurva Himpunan SEDIKIT (Monoton Turun: `120 – 180 – 540`)**
```text
μ(z)
 1.0 |-------\
     |        \
 0.5 |         \
     |          \
 0.0 +-----------\-----------------> JUMLAH PRODUKSI
   120     180  540              640
```

**2. Kurva Himpunan BANYAK (Monoton Naik: `220 – 580 – 640`)**
```text
μ(z)
 1.0 |               /-------------
     |              /
 0.5 |             /
     |            /
 0.0 +-----------/-----------------> JUMLAH PRODUKSI
   120          220 580          640
```

**3. Kurva Gabungan Variabel JUMLAH PRODUKSI (2 Himpunan Monoton)**
```text
μ(z)   [SEDIKIT]          [BANYAK]
 1.0 |-------\             /-------
     |        \           /
 0.5 |         \         /
     |          \       /
 0.0 +-----------\-----/-----------> JUMLAH PRODUKSI
   120     180  220   540 580    640
```

---

## 3. Tahap 1: Fuzzifikasi (Mencari Derajat Keanggotaan / Nilai Fuzzy)

### A. Penjabaran Rumus 4 Bentuk Kurva

1. **Rumus Kurva Turun Linear** *(batas kiri $a$, batas kanan $b$)*:
   * Jika $\text{Input} \le a \rightarrow$ nilainya **`1`**
   * Jika $\text{Input} \ge b \rightarrow$ nilainya **`0`**
   * Jika $a < \text{Input} < b \rightarrow \dfrac{b - \text{Input}}{b - a}$

2. **Rumus Kurva Naik Linear** *(batas kiri $a$, batas kanan $b$)*:
   * Jika $\text{Input} \le a \rightarrow$ nilainya **`0`**
   * Jika $\text{Input} \ge b \rightarrow$ nilainya **`1`**
   * Jika $a < \text{Input} < b \rightarrow \dfrac{\text{Input} - a}{b - a}$

3. **Rumus Kurva Segitiga** *(batas kiri $a$, puncak $b$, batas kanan $c$)*:
   * Jika $\text{Input} \le a$ atau $\text{Input} \ge c \rightarrow$ nilainya **`0`**
   * Jika $\text{Input} = b \rightarrow$ nilainya **`1`**
   * Jika $a < \text{Input} < b$ *(sisi naik)* $\rightarrow \dfrac{\text{Input} - a}{b - a}$
   * Jika $b < \text{Input} < c$ *(sisi turun)* $\rightarrow \dfrac{c - \text{Input}}{c - b}$

4. **Rumus Kurva Trapesium** *(batas kiri $a$, puncak awal $b$, puncak akhir $c$, batas kanan $d$)*:
   * Jika $\text{Input} \le a$ atau $\text{Input} \ge d \rightarrow$ nilainya **`0`**
   * Jika $b \le \text{Input} \le c$ *(dataran puncak)* $\rightarrow$ nilainya **`1`**
   * Jika $a < \text{Input} < b$ *(sisi naik)* $\rightarrow \dfrac{\text{Input} - a}{b - a}$
   * Jika $c < \text{Input} < d$ *(sisi turun)* $\rightarrow \dfrac{d - \text{Input}}{d - c}$

---

### B. Perhitungan Fuzzifikasi Input: PERSEDIAAN = `46`
*(Mengacu pada Kurva Gabungan PERSEDIAAN)*

1. **Himpunan SEDIKIT (`12 – 22 – 54`):**
   * Karena `46` berada di antara `22` dan `54` (sisi turun), gunakan rumus **Kurva Turun Linear**:
   * $\dfrac{54 - 46}{54 - 22} = \dfrac{8}{32} =$ **`0.25`**
2. **Himpunan SEDANG (`34 – 51 – 68`):**
   * Karena `46` berada di antara `34` dan `51` (sisi naik menuju puncak `51`), gunakan rumus **Kurva Segitiga sisi naik**:
   * $\dfrac{46 - 34}{51 - 34} = \dfrac{12}{17} \approx$ **`0.71`** *(dibulatkan 2 desimal)*
3. **Himpunan BANYAK (`52 – 84 – 88`):**
   * Karena $46 \le 52$ (belum menyentuh titik awal garis naik `52`), nilainya langsung **`0`**.

> **Catatan Penjelasan (PERSEDIAAN = 46):**
> Pada kurva gabungan PERSEDIAAN, angka `46` memotong dua garis sekaligus, yaitu garis turun **SEDIKIT** (dengan nilai `0.25`) dan garis naik **SEDANG** (dengan nilai `0.71`), sementara garis **BANYAK** belum dimulai. Artinya, persediaan `46` porsi lebih dominan masuk ke kategori **SEDANG** (`0.71`), namun masih memiliki sedikit unsur **SEDIKIT** (`0.25`).

---

### C. Perhitungan Fuzzifikasi Input: PERMINTAAN = `68`
*(Mengacu pada Kurva Gabungan PERMINTAAN)*

1. **Himpunan SEPI (`14 – 23 – 47`):**
   * Karena $68 \ge 47$ (sudah melewati batas akhir garis turun `47`), nilainya langsung **`0`**.
2. **Himpunan NORMAL (`31 – 44 – 63 – 77`):**
   * Karena `68` berada di antara `63` dan `77` (sisi turun setelah dataran puncak), gunakan rumus **Kurva Trapesium sisi turun**:
   * $\dfrac{77 - 68}{77 - 63} = \dfrac{9}{14} \approx$ **`0.64`** *(dibulatkan 2 desimal)*
3. **Himpunan RAMAI (`59 – 86 – 92`):**
   * Karena `68` berada di antara `59` dan `86` (sisi naik), gunakan rumus **Kurva Naik Linear**:
   * $\dfrac{68 - 59}{86 - 59} = \dfrac{9}{27} \approx$ **`0.33`** *(dibulatkan 2 desimal)*

> **Catatan Penjelasan (PERMINTAAN = 68):**
> Pada kurva gabungan PERMINTAAN, angka `68` sudah jauh melewati daerah **SEPI** sehingga nilainya `0`. Angka `68` berada di area irisan antara sisi turun **NORMAL** (`0.64`) dan sisi naik **RAMAI** (`0.33`). Ini menunjukkan permintaan `68` porsi masih cenderung **NORMAL**, tetapi sudah mulai bergerak masuk ke kategori **RAMAI**.

---

## 4. Tahap 2: Inisialisasi Aturan (`Rule Base`) & Operasi `IF`

### A. Daftar Lengkap Aturan (11 Rule: 9 `AND` + 2 `OR`)
* **R1:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
* **R2:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI SEDIKIT`
* **R3:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
* **R4:** `IF PERSEDIAAN SEDANG AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
* **R5:** `IF PERSEDIAAN SEDANG AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI BANYAK`
* **R6:** `IF PERSEDIAAN SEDANG AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
* **R7:** `IF PERSEDIAAN BANYAK AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
* **R8:** `IF PERSEDIAAN BANYAK AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI SEDIKIT`
* **R9:** `IF PERSEDIAAN BANYAK AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
* **R10:** `IF PERSEDIAAN SEDIKIT OR PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
* **R11:** `IF PERSEDIAAN BANYAK OR PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`

---

### B. Rumus Sederhana Operasi `IF` (Mencari Nilai Predikat / $\alpha$)
* **Jika aturan memakai operator `AND`:**
  $$\alpha = \min(\text{Syarat 1}, \text{Syarat 2})$$
  *(Ambil nilai paling kecil di antara derajat keanggotaan syarat-syaratnya)*
* **Jika aturan memakai operator `OR`:**
  $$\alpha = \max(\text{Syarat 1}, \text{Syarat 2})$$
  *(Ambil nilai paling besar di antara derajat keanggotaan syarat-syaratnya)*

---

### C. Proses Operasi `IF` Per Aturan

* **R1:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN SEPI`
  * Operasi: $\min(0.25, 0) =$ **`0`** *(Output: SEDIKIT)*
* **R2:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN NORMAL`
  * Operasi: $\min(0.25, 0.64) =$ **`0.25`** *(Output: SEDIKIT)*
* **R3:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN RAMAI`
  * Operasi: $\min(0.25, 0.33) =$ **`0.25`** *(Output: BANYAK)*
* **R4:** `IF PERSEDIAAN SEDANG AND PERMINTAAN SEPI`
  * Operasi: $\min(0.71, 0) =$ **`0`** *(Output: SEDIKIT)*
* **R5:** `IF PERSEDIAAN SEDANG AND PERMINTAAN NORMAL`
  * Operasi: $\min(0.71, 0.64) =$ **`0.64`** *(Output: BANYAK)*
* **R6:** `IF PERSEDIAAN SEDANG AND PERMINTAAN RAMAI`
  * Operasi: $\min(0.71, 0.33) =$ **`0.33`** *(Output: BANYAK)*
* **R7:** `IF PERSEDIAAN BANYAK AND PERMINTAAN SEPI`
  * Operasi: $\min(0, 0) =$ **`0`** *(Output: SEDIKIT)*
* **R8:** `IF PERSEDIAAN BANYAK AND PERMINTAAN NORMAL`
  * Operasi: $\min(0, 0.64) =$ **`0`** *(Output: SEDIKIT)*
* **R9:** `IF PERSEDIAAN BANYAK AND PERMINTAAN RAMAI`
  * Operasi: $\min(0, 0.33) =$ **`0`** *(Output: BANYAK)*
* **R10:** `IF PERSEDIAAN SEDIKIT OR PERMINTAAN RAMAI`
  * Operasi: $\max(0.25, 0.33) =$ **`0.33`** *(Output: BANYAK)*
* **R11:** `IF PERSEDIAAN BANYAK OR PERMINTAAN SEPI`
  * Operasi: $\max(0, 0) =$ **`0`** *(Output: SEDIKIT)*

> **Mengapa `AND` Mengambil Angka Terkecil dan `OR` Mengambil Angka Terbesar?**
> * Operator **`AND`** mengharuskan **kedua syarat terpenuhi bersamaan** (irisan). Kekuatan suatu aturan tidak boleh melebihi syarat yang paling lemah—ibarat rantai yang kekuatannya ditentukan oleh mata rantai terlemah, sehingga diambil nilai **terkecil (`min`)**.
> * Operator **`OR`** hanya membutuhkan **salah satu syarat saja yang terpenuhi** (gabungan). Jika salah satu syarat sudah kuat, aturan tersebut langsung aktif mengikuti syarat yang paling dominan, sehingga diambil nilai **terbesar (`max`)**.

---

## 5. Tahap 3: Proses Hitung `THEN` (Mencari Nilai Output Tegas $z$ Per Aturan)

### A. Patokan Kurva Variabel Output: JUMLAH PRODUKSI (`120 – 640`)
```text
μ(z)   [SEDIKIT]          [BANYAK]
 1.0 |-------\             /-------
     |        \           /
 0.5 |         \         /
     |          \       /
 0.0 +-----------\-----/-----------> JUMLAH PRODUKSI
   120     180  220   540 580    640
```

### B. Rumus Sederhana Hitung `THEN` Metode Tsukamoto
Karena kurva output wajib monoton, nilai tegas $z$ dicari dengan membalik rumus garis lurusnya:

* **Jika Output SEDIKIT (Monoton Turun: batas kiri `180`, batas kanan `540`):**
  $$z = \text{Batas Kanan} - \big(\alpha \times (\text{Batas Kanan} - \text{Batas Kiri})\big)$$
  $$z = 540 - \big(\alpha \times (540 - 180)\big) \implies z = 540 - (\alpha \times 360)$$

* **Jika Output BANYAK (Monoton Naik: batas kiri `220`, batas kanan `580`):**
  $$z = \text{Batas Kiri} + \big(\alpha \times (\text{Batas Kanan} - \text{Batas Kiri})\big)$$
  $$z = 220 + \big(\alpha \times (580 - 220)\big) \implies z = 220 + (\alpha \times 360)$$

---

### C. Perhitungan Nilai $z$ Per Aturan

* **R1:** `THEN JUMLAH PRODUKSI SEDIKIT` *(dengan $\alpha_1 = 0$)*
  * Hitung: $z_1 = 540 - (0 \times 360) =$ **`540`**
* **R2:** `THEN JUMLAH PRODUKSI SEDIKIT` *(dengan $\alpha_2 = 0.25$)*
  * Hitung: $z_2 = 540 - (0.25 \times 360) = 540 - 90 =$ **`450`**
* **R3:** `THEN JUMLAH PRODUKSI BANYAK` *(dengan $\alpha_3 = 0.25$)*
  * Hitung: $z_3 = 220 + (0.25 \times 360) = 220 + 90 =$ **`310`**
* **R4:** `THEN JUMLAH PRODUKSI SEDIKIT` *(dengan $\alpha_4 = 0$)*
  * Hitung: $z_4 = 540 - (0 \times 360) =$ **`540`**
* **R5:** `THEN JUMLAH PRODUKSI BANYAK` *(dengan $\alpha_5 = 0.64$)*
  * Hitung: $z_5 = 220 + (0.64 \times 360) = 220 + 230.4 =$ **`450.4`**
* **R6:** `THEN JUMLAH PRODUKSI BANYAK` *(dengan $\alpha_6 = 0.33$)*
  * Hitung: $z_6 = 220 + (0.33 \times 360) = 220 + 118.8 =$ **`338.8`**
* **R7:** `THEN JUMLAH PRODUKSI SEDIKIT` *(dengan $\alpha_7 = 0$)*
  * Hitung: $z_7 = 540 - (0 \times 360) =$ **`540`**
* **R8:** `THEN JUMLAH PRODUKSI SEDIKIT` *(dengan $\alpha_8 = 0$)*
  * Hitung: $z_8 = 540 - (0 \times 360) =$ **`540`**
* **R9:** `THEN JUMLAH PRODUKSI BANYAK` *(dengan $\alpha_9 = 0$)*
  * Hitung: $z_9 = 220 + (0 \times 360) =$ **`220`**
* **R10:** `THEN JUMLAH PRODUKSI BANYAK` *(dengan $\alpha_{10} = 0.33$)*
  * Hitung: $z_{10} = 220 + (0.33 \times 360) = 220 + 118.8 =$ **`338.8`**
* **R11:** `THEN JUMLAH PRODUKSI SEDIKIT` *(dengan $\alpha_{11} = 0$)*
  * Hitung: $z_{11} = 540 - (0 \times 360) =$ **`540`**

> **Keterangan Asal Angka pada Tahap `THEN`:**
> * **`540`** $\rightarrow$ Batas kanan (titik nol) dari kurva output monoton turun **SEDIKIT**.
> * **`220`** $\rightarrow$ Batas kiri (titik awal naik) dari kurva output monoton naik **BANYAK**.
> * **`360`** $\rightarrow$ Selisih rentang garis miring pada masing-masing kurva output ($540 - 180 = 360$ untuk kurva SEDIKIT, dan $580 - 220 = 360$ untuk kurva BANYAK).
> * **`0`, `0.25`, `0.64`, `0.33`** $\rightarrow$ Nilai kekuatan aturan ($\alpha_1$ sampai $\alpha_{11}$) yang diperoleh dari hasil operasi `IF` pada Tahap 2.

---

## 6. Tahap 4: Defuzzifikasi (Rata-Rata Berbobot)

### A. Rumus Sederhana Defuzzifikasi Tsukamoto
$$\text{Jumlah Produksi Akhir} = \frac{\text{Total dari } (\alpha_i \times z_i)}{\text{Total dari } \alpha_i}$$

---

### B. Proses Perhitungan Akhir

1. **Menghitung Bagian Atas (Pembilang: Total dari $\alpha_i \times z_i$):**
   * R1: $0 \times 540 = 0$
   * R2: $0.25 \times 450 = 112.5$
   * R3: $0.25 \times 310 = 77.5$
   * R4: $0 \times 540 = 0$
   * R5: $0.64 \times 450.4 = 288.256$
   * R6: $0.33 \times 338.8 = 111.804$
   * R7: $0 \times 540 = 0$
   * R8: $0 \times 540 = 0$
   * R9: $0 \times 220 = 0$
   * R10: $0.33 \times 338.8 = 111.804$
   * R11: $0 \times 540 = 0$
   * **Total Bagian Atas** $= 0 + 112.5 + 77.5 + 0 + 288.256 + 111.804 + 0 + 0 + 0 + 111.804 + 0 =$ **`701.864`**

2. **Menghitung Bagian Bawah (Penyebut: Total dari $\alpha_i$):**
   * **Total Bagian Bawah** $= 0 + 0.25 + 0.25 + 0 + 0.64 + 0.33 + 0 + 0 + 0 + 0.33 + 0 =$ **`1.80`**

3. **Membagi Bagian Atas dengan Bagian Bawah:**
   * $\text{Jumlah Produksi} = \dfrac{701.864}{1.80} =$ **`389.92` porsi**
   * Dibulatkan ke satuan porsi terdekat menjadi **`390` porsi**.

> **Keterangan Asal Angka pada Tahap Defuzzifikasi:**
> * **`0, 0.25, 0.25, 0, 0.64, 0.33, 0, 0, 0, 0.33, 0`** $\rightarrow$ Seluruh nilai kekuatan aturan ($\alpha_1$ sampai $\alpha_{11}$) hasil evaluasi `IF` pada Tahap 2.
> * **`540, 450, 310, 540, 450.4, 338.8, 540, 540, 220, 338.8, 540`** $\rightarrow$ Seluruh nilai output tegas ($z_1$ sampai $z_{11}$) hasil perhitungan `THEN` pada Tahap 3.
> * **`701.864`** $\rightarrow$ Hasil penjumlahan seluruh perkalian antara kekuatan tiap aturan dengan nilai output tegasnya ($\sum (\alpha_i \times z_i)$).
> * **`1.80`** $\rightarrow$ Hasil penjumlahan seluruh bobot kekuatan aturan yang aktif ($\sum \alpha_i$).

**Kesimpulan Kasus:**
Untuk sisa **PERSEDIAAN = `46` porsi** dan perkiraan **PERMINTAAN = `68` porsi**, jumlah makanan yang harus diproduksi adalah **`390` porsi**.

---

## 7. Soal Latihan Mandiri untuk Pembaca Catatan

Gunakan tabel batasan, grafik kurva, dan ke-11 aturan (`9 AND + 2 OR`) di atas untuk menyelesaikan dua kasus latihan berikut:

1. **Latihan 1:** Pada hari berikutnya, sisa **PERSEDIAAN = `38` porsi** dan perkiraan **PERMINTAAN = `52` porsi**. Berapa porsi **JUMLAH PRODUKSI** yang disarankan?
   *(Petunjuk: Pada kurva trapesium NORMAL, nilai `52` jatuh di area dataran puncak `44–63`, sehingga derajat keanggotaannya langsung bernilai `1`.)*
2. **Latihan 2:** Saat akhir pekan, sisa **PERSEDIAAN = `61` porsi** dan perkiraan **PERMINTAAN = `39` porsi**. Berapa porsi **JUMLAH PRODUKSI** yang disarankan?
   *(Petunjuk: Pada kurva segitiga SEDANG, nilai `61` sudah melewati puncak `51` sehingga menggunakan rumus sisi turun `51–68`.)*
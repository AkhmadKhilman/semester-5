# Catatan Lengkap: Penyelesaian Studi Kasus Metode Fuzzy Tsukamoto

## Studi Kasus (Soal Cerita)

Sebuah usaha kuliner rumahan setiap pagi harus menentukan berapa porsi makanan yang perlu diproduksi agar tidak mengalami kekurangan stok saat pesanan ramai maupun kerugian akibat makanan sisa yang tidak terjual. Berdasarkan pengalaman pengelola, keputusan jumlah produksi harian dipengaruhi oleh dua faktor utama, yaitu sisa stok **PERSEDIAAN** bahan/porsi yang masih ada di dapur dan perkiraan jumlah **PERMINTAAN** pelanggan pada hari tersebut.

Pengelola menetapkan batasan operasional harian sebagai berikut:

- **PERSEDIAAN** berada pada rentang `12` hingga `88` porsi.
- **PERMINTAAN** berada pada rentang `14` hingga `92` porsi.
- **JUMLAH PRODUKSI** yang dihasilkan berada pada rentang `120` hingga `640` porsi.

Pada suatu hari, bagian dapur mencatat bahwa sisa **PERSEDIAAN = 46 porsi**, sementara perkiraan **PERMINTAAN = 68 porsi**. Berapa porsi **JUMLAH PRODUKSI** yang harus dibuat hari itu menggunakan sistem inferensi **Metode Fuzzy Tsukamoto**?

---

## 1. Tabel Variabel, Himpunan, dan Batasan

### A. Tabel Variabel Input

| Jenis Variabel | Nama Variabel  | Semesta Pembicaraan / Batasan Variabel | Nama Himpunan | Bentuk Kurva | Domain / Batasan Himpunan |
| :------------- | :------------- | :------------------------------------: | :------------ | :----------- | :-----------------------: |
| **Input 1**    | **PERSEDIAAN** |               `12 – 88`                | **SEDIKIT**   | Turun Linear |         `22 – 54`         |
| **Input 1**    | **PERSEDIAAN** |               `12 – 88`                | **SEDANG**    | Segitiga     |      `34 – 51 – 68`       |
| **Input 1**    | **PERSEDIAAN** |               `12 – 88`                | **BANYAK**    | Naik Linear  |         `52 – 84`         |
| **Input 2**    | **PERMINTAAN** |               `14 – 92`                | **SEPI**      | Turun Linear |         `23 – 47`         |
| **Input 2**    | **PERMINTAAN** |               `14 – 92`                | **NORMAL**    | Trapesium    |    `31 – 44 – 63 – 77`    |
| **Input 2**    | **PERMINTAAN** |               `14 – 92`                | **RAMAI**     | Naik Linear  |         `59 – 86`         |

### B. Tabel Variabel Output _(Terpisah — Wajib Monoton untuk Metode Tsukamoto)_

| Jenis Variabel | Nama Variabel       | Semesta Pembicaraan / Batasan Variabel | Nama Himpunan | Bentuk Kurva | Domain / Batasan Himpunan |
| :------------- | :------------------ | :------------------------------------: | :------------ | :----------- | :-----------------------: |
| **Output**     | **JUMLAH PRODUKSI** |              `120 – 640`               | **SEDIKIT**   | Turun Linear |        `180 – 540`        |
| **Output**     | **JUMLAH PRODUKSI** |              `120 – 640`               | **BANYAK**    | Naik Linear  |        `220 – 580`        |

---

## 2. Pembuatan Grafik Kurva Tiap Variabel

### A. Kurva Variabel Input 1: PERSEDIAAN (`12 – 88`)

**1. Kurva Himpunan SEDIKIT (Turun Linear: `22 – 54`)**

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

**3. Kurva Himpunan BANYAK (Naik Linear: `52 – 84`)**

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

**1. Kurva Himpunan SEPI (Turun Linear: `23 – 47`)**

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

**3. Kurva Himpunan RAMAI (Naik Linear: `59 – 86`)**

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

**1. Kurva Himpunan SEDIKIT ( Turun Linear: `180 – 540`)**

```text
μ(z)
 1.0 |-------\
     |        \
 0.5 |         \
     |          \
 0.0 +-----------\-----------------> JUMLAH PRODUKSI
   120     180  540              640
```

**2. Kurva Himpunan BANYAK ( Naik: Linear `220 – 580`)**

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

## 3. Tahap 1: Fuzzifikasi (Mencari Derajat Keanggotaan)

### A. Penjabaran Rumus 4 Bentuk Kurva

1. **Rumus Kurva Turun Linear** _(batas kiri `a`, batas kanan `b`)_:
   - Jika nilai inputan **lebih kecil atau sama dengan `a`**, maka nilainya adalah **`1`**.
   - Jika nilai inputan **lebih besar atau sama dengan `b`**, maka nilainya adalah **`0`**.
   - Jika nilai inputan **berada di antara `a` dan `b`**, maka gunakan rumus: **`(b - Input) dibagi (b - a)`**.

2. **Rumus Kurva Naik Linear** _(batas kiri `a`, batas kanan `b`)_:
   - Jika nilai inputan **lebih kecil atau sama dengan `a`**, maka nilainya adalah **`0`**.
   - Jika nilai inputan **lebih besar atau sama dengan `b`**, maka nilainya adalah **`1`**.
   - Jika nilai inputan **berada di antara `a` dan `b`**, maka gunakan rumus: **`(Input - a) dibagi (b - a)`**.

3. **Rumus Kurva Segitiga** _(batas kiri `a`, puncak `b`, batas kanan `c`)_:
   - Jika nilai inputan **lebih kecil atau sama dengan `a`**, atau **lebih besar atau sama dengan `c`**, maka nilainya adalah **`0`**.
   - Jika nilai inputan **sama dengan titik puncak `b`**, maka nilainya adalah **`1`**.
   - Jika nilai inputan **berada di antara `a` dan `b`** _(sisi garis naik)_, maka gunakan rumus: **`(Input - a) dibagi (b - a)`**.
   - Jika nilai inputan **berada di antara `b` dan `c`** _(sisi garis turun)_, maka gunakan rumus: **`(c - Input) dibagi (c - b)`**.

4. **Rumus Kurva Trapesium** _(batas kiri `a`, puncak awal `b`, puncak akhir `c`, batas kanan `d`)_:
   - Jika nilai inputan **lebih kecil atau sama dengan `a`**, atau **lebih besar atau sama dengan `d`**, maka nilainya adalah **`0`**.
   - Jika nilai inputan **berada di antara atau sama dengan `b` dan `c`** _(area dataran puncak)_, maka nilainya adalah **`1`**.
   - Jika nilai inputan **berada di antara `a` dan `b`** _(sisi garis naik)_, maka gunakan rumus: **`(Input - a) dibagi (b - a)`**.
   - Jika nilai inputan **berada di antara `c` dan `d`** _(sisi garis turun)_, maka gunakan rumus: **`(d - Input) dibagi (d - c)`**.

---

### B. Perhitungan Fuzzifikasi Input: PERSEDIAAN = `46`

```text
μ(x)   [SEDIKIT]    [SEDANG]      [BANYAK]
 1.0 |-------\         /\           /-------
     |        \       /  \         /
 0.5 |         \     /    \       /
     |          \   /      \     /
 0.0 +-----------\-/--------\---/----------> PERSEDIAAN
    12      22   34 51  52  54 68  84     88
```

1. **Himpunan SEDIKIT (`22 – 54` — Kurva Turun Linear):**
   - **Kondisi:** Nilai inputan `46` berada di antara `22` dan `54`.
   - **Gunakan aturan:** Jika nilai inputan **berada di antara `a` dan `b`**, maka gunakan rumus: **`(b - Input) dibagi (b - a)`**.
   - **Keterangan nilai:** `Input = 46`, batas kiri `a = 22`, dan batas kanan `b = 54`.
   - **Perhitungan:**
     `(54 - 46) dibagi (54 - 22)` = `8 dibagi 32` = **`0.25`**

2. **Himpunan SEDANG (`34 – 51 – 68` — Kurva Segitiga):**
   - **Kondisi:** Nilai inputan `46` berada di antara `34` dan `51` (sisi garis naik menuju titik puncak `51`).
   - **Gunakan aturan:** Jika nilai inputan **berada di antara `a` dan `b`** _(sisi garis naik)_, maka gunakan rumus: **`(Input - a) dibagi (b - a)`**.
   - **Keterangan nilai:** `Input = 46`, batas kiri `a = 34`, puncak `b = 51`, dan batas kanan `c = 68`.
   - **Perhitungan:**
     `(46 - 34) dibagi (51 - 34)` = `12 dibagi 17` = **`0.71`** _(dibulatkan 2 desimal)_

3. **Himpunan BANYAK (`52 – 84` — Kurva Naik Linear):**
   - **Kondisi:** Nilai inputan `46` lebih kecil dari `52`.
   - **Gunakan aturan:** Jika nilai inputan **lebih kecil atau sama dengan `a`**, maka nilainya adalah **`0`**.
   - **Keterangan nilai:** `Input = 46`, batas kiri `a = 52`, dan batas kanan `b = 84`.
   - **Perhitungan:**
     Karena `46` lebih kecil dari `a` (`52`), maka nilainya langsung **`0`**.

---

### C. Perhitungan Fuzzifikasi Input: PERMINTAAN = `68`

```text
μ(x)    [SEPI]       [NORMAL]        [RAMAI]
 1.0 |-------\      /--------\        /-------
     |        \    /          \      /
 0.5 |         \  /            \    /
     |          \/              \  /
 0.0 +----------/\---------------\/----------> PERMINTAAN
    14     23  31 44   47   59  63 77  86   92
```

1. **Himpunan SEPI (`23 – 47` — Kurva Turun Linear):**
   - **Kondisi:** Nilai inputan `68` lebih besar dari `47`.
   - **Gunakan aturan:** Jika nilai inputan **lebih besar atau sama dengan `b`**, maka nilainya adalah **`0`**.
   - **Keterangan nilai:** `Input = 68`, batas kiri `a = 23`, dan batas kanan `b = 47`.
   - **Perhitungan:**
     Karena `68` lebih besar dari `b` (`47`), maka nilainya langsung **`0`**.

2. **Himpunan NORMAL (`31 – 44 – 63 – 77` — Kurva Trapesium):**
   - **Kondisi:** Nilai inputan `68` berada di antara `63` dan `77` (sisi garis turun setelah dataran puncak).
   - **Gunakan aturan:** Jika nilai inputan **berada di antara `c` dan `d`** _(sisi garis turun)_, maka gunakan rumus: **`(d - Input) dibagi (d - c)`**.
   - **Keterangan nilai:** `Input = 68`, batas kiri `a = 31`, puncak awal `b = 44`, puncak akhir `c = 63`, dan batas kanan `d = 77`.
   - **Perhitungan:**
     `(77 - 68) dibagi (77 - 63)` = `9 dibagi 14` = **`0.64`** _(dibulatkan 2 desimal)_

3. **Himpunan RAMAI (`59 – 86` — Kurva Naik Linear):**
   - **Kondisi:** Nilai inputan `68` berada di antara `59` dan `86`.
   - **Gunakan aturan:** Jika nilai inputan **berada di antara `a` dan `b`**, maka gunakan rumus: **`(Input - a) dibagi (b - a)`**.
   - **Keterangan nilai:** `Input = 68`, batas kiri `a = 59`, dan batas kanan `b = 86`.
   - **Perhitungan:**
     `(68 - 59) dibagi (86 - 59)` = `9 dibagi 27` = **`0.33`** _(dibulatkan 2 desimal)_

---

## 4. Tahap 2: Inisialisasi Aturan (`Rule Base`) & Operasi `IF`

### A. Daftar Lengkap Aturan (11 Rule: 9 `AND` + 2 `OR`)

- **R1:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
- **R2:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI SEDIKIT`
- **R3:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
- **R4:** `IF PERSEDIAAN SEDANG AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
- **R5:** `IF PERSEDIAAN SEDANG AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI BANYAK`
- **R6:** `IF PERSEDIAAN SEDANG AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
- **R7:** `IF PERSEDIAAN BANYAK AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
- **R8:** `IF PERSEDIAAN BANYAK AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI SEDIKIT`
- **R9:** `IF PERSEDIAAN BANYAK AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
- **R10:** `IF PERSEDIAAN SEDIKIT OR PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
- **R11:** `IF PERSEDIAAN BANYAK OR PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`

---

### B. Aturan Sederhana Operasi `IF`

- Jika rule/aturan menggunakan operator **`AND`**, maka **ambil nilai paling kecil** di antara derajat keanggotaan syarat-syaratnya.
- Jika rule/aturan menggunakan operator **`OR`**, maka **ambil nilai paling besar** di antara derajat keanggotaan syarat-syaratnya.

---

### C. Proses Operasi `IF` Per Aturan

- **R1:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0.25`** berasal dari derajat keanggotaan `PERSEDIAAN SEDIKIT`, dan **`0`** berasal dari derajat keanggotaan `PERMINTAAN SEPI`.
  - Perbandingan: Nilai terkecil antara `0.25` dan `0` adalah **`0`** _(untuk output: SEDIKIT)_.

- **R2:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0.25`** berasal dari derajat keanggotaan `PERSEDIAAN SEDIKIT`, dan **`0.64`** berasal dari derajat keanggotaan `PERMINTAAN NORMAL`.
  - Perbandingan: Nilai terkecil antara `0.25` dan `0.64` adalah **`0.25`** _(untuk output: SEDIKIT)_.

- **R3:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0.25`** berasal dari derajat keanggotaan `PERSEDIAAN SEDIKIT`, dan **`0.33`** berasal dari derajat keanggotaan `PERMINTAAN RAMAI`.
  - Perbandingan: Nilai terkecil antara `0.25` dan `0.33` adalah **`0.25`** _(untuk output: BANYAK)_.

- **R4:** `IF PERSEDIAAN SEDANG AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0.71`** berasal dari derajat keanggotaan `PERSEDIAAN SEDANG`, dan **`0`** berasal dari derajat keanggotaan `PERMINTAAN SEPI`.
  - Perbandingan: Nilai terkecil antara `0.71` dan `0` adalah **`0`** _(untuk output: SEDIKIT)_.

- **R5:** `IF PERSEDIAAN SEDANG AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI BANYAK`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0.71`** berasal dari derajat keanggotaan `PERSEDIAAN SEDANG`, dan **`0.64`** berasal dari derajat keanggotaan `PERMINTAAN NORMAL`.
  - Perbandingan: Nilai terkecil antara `0.71` dan `0.64` adalah **`0.64`** _(untuk output: BANYAK)_.

- **R6:** `IF PERSEDIAAN SEDANG AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0.71`** berasal dari derajat keanggotaan `PERSEDIAAN SEDANG`, dan **`0.33`** berasal dari derajat keanggotaan `PERMINTAAN RAMAI`.
  - Perbandingan: Nilai terkecil antara `0.71` dan `0.33` adalah **`0.33`** _(untuk output: BANYAK)_.

- **R7:** `IF PERSEDIAAN BANYAK AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0`** berasal dari derajat keanggotaan `PERSEDIAAN BANYAK`, dan **`0`** berasal dari derajat keanggotaan `PERMINTAAN SEPI`.
  - Perbandingan: Nilai terkecil antara `0` dan `0` adalah **`0`** _(untuk output: SEDIKIT)_.

- **R8:** `IF PERSEDIAAN BANYAK AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0`** berasal dari derajat keanggotaan `PERSEDIAAN BANYAK`, dan **`0.64`** berasal dari derajat keanggotaan `PERMINTAAN NORMAL`.
  - Perbandingan: Nilai terkecil antara `0` dan `0.64` adalah **`0`** _(untuk output: SEDIKIT)_.

- **R9:** `IF PERSEDIAAN BANYAK AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
  - Karena aturan menggunakan operator **`AND`**, maka ambil nilai terkecil.
  - Keterangan asal angka: **`0`** berasal dari derajat keanggotaan `PERSEDIAAN BANYAK`, dan **`0.33`** berasal dari derajat keanggotaan `PERMINTAAN RAMAI`.
  - Perbandingan: Nilai terkecil antara `0` dan `0.33` adalah **`0`** _(untuk output: BANYAK)_.

- **R10:** `IF PERSEDIAAN SEDIKIT OR PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
  - Karena aturan menggunakan operator **`OR`**, maka ambil nilai terbesar.
  - Keterangan asal angka: **`0.25`** berasal dari derajat keanggotaan `PERSEDIAAN SEDIKIT`, dan **`0.33`** berasal dari derajat keanggotaan `PERMINTAAN RAMAI`.
  - Perbandingan: Nilai terbesar antara `0.25` dan `0.33` adalah **`0.33`** _(untuk output: BANYAK)_.

- **R11:** `IF PERSEDIAAN BANYAK OR PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena aturan menggunakan operator **`OR`**, maka ambil nilai terbesar.
  - Keterangan asal angka: **`0`** berasal dari derajat keanggotaan `PERSEDIAAN BANYAK`, dan **`0`** berasal dari derajat keanggotaan `PERMINTAAN SEPI`.
  - Perbandingan: Nilai terbesar antara `0` dan `0` adalah **`0`** _(untuk output: SEDIKIT)_.

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

Karena kurva variabel output wajib berbentuk monoton (naik terus atau turun terus), cara menghitung nilai tegasnya ditentukan dari akhir kalimat `THEN` pada masing-masing aturan:

1. **Jika aturan berakhiran `THEN JUMLAH PRODUKSI SEDIKIT`:**
   Artinya output yang dituju masuk ke himpunan **SEDIKIT** (berbentuk kurva linear turun), maka gunakan rumus berikut:
   **`Batas Kanan - (X x (Batas Kanan - Batas Kiri))`**
   - **Keterangan nilai pada kurva SEDIKIT:**
     - **`Batas Kiri`** adalah titik awal garis mulai turun, yaitu angka **`180`**.
     - **`Batas Kanan`** adalah titik akhir garis menyentuh nol, yaitu angka **`540`**.
     - **`X`** adalah nilai kekuatan aturan yang diperoleh dari hasil operasi **`IF`** pada tahap sebelumnya.
   - **Rumus siap pakai untuk output SEDIKIT:**
     **`540 - (X x (540 - 180))`** atau disederhanakan menjadi **`540 - (X x 360)`**

2. **Jika aturan berakhiran `THEN JUMLAH PRODUKSI BANYAK`:**
   Artinya output yang dituju masuk ke himpunan **BANYAK** (berbentuk kurva linear naik), maka gunakan rumus berikut:
   **`Batas Kiri + (X x (Batas Kanan - Batas Kiri))`**
   - **Keterangan nilai pada kurva BANYAK:**
     - **`Batas Kiri`** adalah titik awal garis mulai naik, yaitu angka **`220`**.
     - **`Batas Kanan`** adalah titik akhir garis mencapai puncak penuh, yaitu angka **`580`**.
     - **`X`** adalah nilai kekuatan aturan yang diperoleh dari hasil operasi **`IF`** pada tahap sebelumnya.
   - **Rumus siap pakai untuk output BANYAK:**
     **`220 + (X x (580 - 220))`** atau disederhanakan menjadi **`220 + (X x 360)`**

---

### C. Perhitungan Nilai Output Tegas Per Aturan

- **R1:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena berakhiran `THEN JUMLAH PRODUKSI SEDIKIT`, gunakan rumus output **SEDIKIT** dengan nilai `X = 0` _(hasil operasi IF pada R1)_.
  - Perhitungan: `540 - (0 x (540 - 180))` = `540 - (0 x 360)` = **`540`**

- **R2:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena berakhiran `THEN JUMLAH PRODUKSI SEDIKIT`, gunakan rumus output **SEDIKIT** dengan nilai `X = 0.25` _(hasil operasi IF pada R2)_.
  - Perhitungan: `540 - (0.25 x (540 - 180))` = `540 - (0.25 x 360)` = `540 - 90` = **`450`**

- **R3:** `IF PERSEDIAAN SEDIKIT AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
  - Karena berakhiran `THEN JUMLAH PRODUKSI BANYAK`, gunakan rumus output **BANYAK** dengan nilai `X = 0.25` _(hasil operasi IF pada R3)_.
  - Perhitungan: `220 + (0.25 x (580 - 220))` = `220 + (0.25 x 360)` = `220 + 90` = **`310`**

- **R4:** `IF PERSEDIAAN SEDANG AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena berakhiran `THEN JUMLAH PRODUKSI SEDIKIT`, gunakan rumus output **SEDIKIT** dengan nilai `X = 0` _(hasil operasi IF pada R4)_.
  - Perhitungan: `540 - (0 x (540 - 180))` = `540 - (0 x 360)` = **`540`**

- **R5:** `IF PERSEDIAAN SEDANG AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI BANYAK`
  - Karena berakhiran `THEN JUMLAH PRODUKSI BANYAK`, gunakan rumus output **BANYAK** dengan nilai `X = 0.64` _(hasil operasi IF pada R5)_.
  - Perhitungan: `220 + (0.64 x (580 - 220))` = `220 + (0.64 x 360)` = `220 + 230.4` = **`450.4`**

- **R6:** `IF PERSEDIAAN SEDANG AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
  - Karena berakhiran `THEN JUMLAH PRODUKSI BANYAK`, gunakan rumus output **BANYAK** dengan nilai `X = 0.33` _(hasil operasi IF pada R6)_.
  - Perhitungan: `220 + (0.33 x (580 - 220))` = `220 + (0.33 x 360)` = `220 + 118.8` = **`338.8`**

- **R7:** `IF PERSEDIAAN BANYAK AND PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena berakhiran `THEN JUMLAH PRODUKSI SEDIKIT`, gunakan rumus output **SEDIKIT** dengan nilai `X = 0` _(hasil operasi IF pada R7)_.
  - Perhitungan: `540 - (0 x (540 - 180))` = `540 - (0 x 360)` = **`540`**

- **R8:** `IF PERSEDIAAN BANYAK AND PERMINTAAN NORMAL THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena berakhiran `THEN JUMLAH PRODUKSI SEDIKIT`, gunakan rumus output **SEDIKIT** dengan nilai `X = 0` _(hasil operasi IF pada R8)_.
  - Perhitungan: `540 - (0 x (540 - 180))` = `540 - (0 x 360)` = **`540`**

- **R9:** `IF PERSEDIAAN BANYAK AND PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
  - Karena berakhiran `THEN JUMLAH PRODUKSI BANYAK`, gunakan rumus output **BANYAK** dengan nilai `X = 0` _(hasil operasi IF pada R9)_.
  - Perhitungan: `220 + (0 x (580 - 220))` = `220 + (0 x 360)` = **`220`**

- **R10:** `IF PERSEDIAAN SEDIKIT OR PERMINTAAN RAMAI THEN JUMLAH PRODUKSI BANYAK`
  - Karena berakhiran `THEN JUMLAH PRODUKSI BANYAK`, gunakan rumus output **BANYAK** dengan nilai `X = 0.33` _(hasil operasi IF pada R10)_.
  - Perhitungan: `220 + (0.33 x (580 - 220))` = `220 + (0.33 x 360)` = `220 + 118.8` = **`338.8`**

- **R11:** `IF PERSEDIAAN BANYAK OR PERMINTAAN SEPI THEN JUMLAH PRODUKSI SEDIKIT`
  - Karena berakhiran `THEN JUMLAH PRODUKSI SEDIKIT`, gunakan rumus output **SEDIKIT** dengan nilai `X = 0` _(hasil operasi IF pada R11)_.
  - Perhitungan: `540 - (0 x (540 - 180))` = `540 - (0 x 360)` = **`540`**

---

## 6. Tahap 4: Defuzzifikasi (Rata-Rata Berbobot)

### A. Rumus Sederhana Defuzzifikasi Tsukamoto

**`Jumlah Produksi Akhir = Total dari (hasil operasi IF dikali hasil operasi THEN) dibagi Total dari (hasil operasi IF)`**

- **Keterangan nilai pada kasus ini:**
  - **Hasil operasi `IF`** _(dari R1 sampai R11)_: **`0`, `0.25`, `0.25`, `0`, `0.64`, `0.33`, `0`, `0`, `0`, `0.33`, `0`**.
  - **Hasil operasi `THEN`** _(dari R1 sampai R11)_: **`540`, `450`, `310`, `540`, `450.4`, `338.8`, `540`, `540`, `220`, `338.8`, `540`**.

---

### B. Proses Perhitungan Akhir

1. **Menghitung Bagian Atas (Total dari hasil operasi `IF` dikali hasil operasi `THEN`):**
   - R1: `0 x 540` = `0`
   - R2: `0.25 x 450` = `112.5`
   - R3: `0.25 x 310` = `77.5`
   - R4: `0 x 540` = `0`
   - R5: `0.64 x 450.4` = `288.256`
   - R6: `0.33 x 338.8` = `111.804`
   - R7: `0 x 540` = `0`
   - R8: `0 x 540` = `0`
   - R9: `0 x 220` = `0`
   - R10: `0.33 x 338.8` = `111.804`
   - R11: `0 x 540` = `0`
   - **Total Bagian Atas** = `0 + 112.5 + 77.5 + 0 + 288.256 + 111.804 + 0 + 0 + 0 + 111.804 + 0` = **`701.864`**

2. **Menghitung Bagian Bawah (Total dari hasil operasi `IF`):**
   - **Total Bagian Bawah** = `0 + 0.25 + 0.25 + 0 + 0.64 + 0.33 + 0 + 0 + 0 + 0.33 + 0` = **`1.80`**

3. **Membagi Bagian Atas dengan Bagian Bawah:**
   - `Jumlah Produksi Akhir = 701.864 dibagi 1.80` = **`389.92` porsi**
   - Dibulatkan ke satuan porsi terdekat menjadi **`390` porsi**.

**Kesimpulan Kasus:**
Untuk sisa **PERSEDIAAN = `46` porsi** dan perkiraan **PERMINTAAN = `68` porsi**, jumlah makanan yang harus diproduksi adalah **`390` porsi**.

---

## 7. Soal Latihan Mandiri untuk Pembaca Catatan

Gunakan tabel batasan, grafik kurva, dan ke-11 aturan (`9 AND + 2 OR`) di atas untuk menyelesaikan dua kasus latihan berikut:

1. **Latihan 1:** Pada hari berikutnya, sisa **PERSEDIAAN = `38` porsi** dan perkiraan **PERMINTAAN = `52` porsi**. Berapa porsi **JUMLAH PRODUKSI** yang disarankan?
2. **Latihan 2:** Saat akhir pekan, sisa **PERSEDIAAN = `61` porsi** dan perkiraan **PERMINTAAN = `39` porsi**. Berapa porsi **JUMLAH PRODUKSI** yang disarankan?

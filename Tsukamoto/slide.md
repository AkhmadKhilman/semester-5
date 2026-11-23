# Kerangka Susunan Slide Presentasi: Sistem Inferensi Fuzzy Metode Tsukamoto

---

## Slide 1: Halaman Judul

- **Judul Utama:** Sistem Inferensi Fuzzy Metode Tsukamoto: Konsep Dasar dan Penerapan Studi Kasus
- **Subjudul:** Mata Kuliah [Nama Mata Kuliah]
- **Identitas:** Nama Penyaji / Kelompok, NIM, dan Nama Dosen Pengampu

---

## Slide 2: Teori Dasar Logika Fuzzy (Bagian 1 — Konsep & Alur)

- **Judul Slide:** Konsep Dasar Logika Fuzzy
- **Poin Materi:**
  - **Logika Tegas (_Crisp_) vs Logika Fuzzy:** Logika tegas hanya mengenal nilai mutlak (`0` atau `1` / Ya atau Tidak), sedangkan Logika Fuzzy menoleransi nilai samar di antara `0` hingga `1`.
  - **Alasan Penggunaan:** Mampu menerjemahkan bahasa alami manusia (seperti _"sedikit"_, _"sedang"_, _"ramai"_) ke dalam perhitungan matematis komputer.
  - **3 Tahap Utama Sistem Fuzzy:**
    1. **Fuzzifikasi:** Mengubah angka masukan tegas menjadi nilai derajat keanggotaan (`0 – 1`).
    2. **Inferensi (`Rule Base`):** Mengevaluasi aturan pakar `IF – THEN`.
    3. **Defuzzifikasi:** Mengubah kembali hasil fuzzy menjadi satu angka keputusan tegas.

---

## Slide 3: Teori Dasar Logika Fuzzy (Bagian 2 — Bentuk Kurva Keanggotaan)

- **Judul Slide:** 4 Bentuk Kurva Fungsi Keanggotaan
- **Poin Materi (Visual + Penjelasan Singkat):**
  1. **Kurva Turun Linear:** Bernilai penuh (`1`) di batas kiri, lalu turun lurus menuju `0` di batas kanan.
  2. **Kurva Naik Linear:** Bernilai `0` di batas kiri, lalu naik lurus menuju penuh (`1`) di batas kanan.
  3. **Kurva Segitiga:** Naik dari `0` menuju **satu titik puncak (`1`)**, lalu turun kembali ke `0`.
  4. **Kurva Trapesium:** Naik dari `0` menuju **dataran puncak (`1`)**, lalu turun kembali ke `0`.

---

## Slide 4: Teori Metode Fuzzy Tsukamoto (Bagian 1 — Karakteristik Utama)

- **Judul Slide:** Mengenal Metode Fuzzy Tsukamoto
- **Poin Materi:**
  - **Pengertian:** Salah satu metode inferensi fuzzy yang setiap konsekuen (`THEN`) pada aturannya direpresentasikan menggunakan himpunan fuzzy dengan **kurva monoton**.
  - **Aturan Mutlak Tsukamoto:**
    - Kurva pada variabel **Input (`IF`)** bebas (boleh turun, naik, segitiga, maupun trapesium).
    - Kurva pada variabel **Output (`THEN`)** **WAJIB monoton** (hanya naik terus atau turun terus, tidak boleh segitiga/trapesium).
  - **Pembeda Singkat:** Mamdani menghasilkan area kurva di output, Sugeno menghasilkan konstanta/persamaan linear, sedangkan **Tsukamoto menghasilkan nilai tegas per aturan dari kurva monoton**.

---

## Slide 5: Teori Metode Fuzzy Tsukamoto (Bagian 2 — Rumus & Mekanisme Kerja)

- **Judul Slide:** Mekanisme Perhitungan Metode Tsukamoto
- **Poin Materi:**
  1. **Operasi `IF` (Mencari Kekuatan Aturan):**
     - Operator **`AND`** $\rightarrow$ ambil nilai **terkecil (`min`)**.
     - Operator **`OR`** $\rightarrow$ ambil nilai **terbesar (`max`)**.
  2. **Operasi `THEN` (Mencari Nilai Tegas Tiap Aturan):**
     - Output **Monoton Turun:** `Batas Kanan - (Hasil IF x (Batas Kanan - Batas Kiri))`
     - Output **Monoton Naik:** `Batas Kiri + (Hasil IF x (Batas Kanan - Batas Kiri))`
  3. **Defuzzifikasi (Rata-Rata Berbobot):**
     - `Hasil Akhir = Total dari (Hasil IF x Hasil THEN) dibagi Total dari (Hasil IF)`

---

## Slide 6: Studi Kasus 1 — Setup & Kurva (Fokus Kurva Segitiga)

- **Judul Slide:** Studi Kasus 1: Penentuan Skor Bonus Karyawan
- **Soal Cerita Singkat:**
  Perusahaan menentukan **SKOR BONUS** (`10 – 90`) berdasarkan **MASA KERJA** (`10 – 50` bulan) dan **NILAI KINERJA** (`40 – 90`). Kasus uji: **Masa Kerja = `25` bulan** dan **Nilai Kinerja = `70`**.
- **Tabel Variabel & Batasan (Angka Puluhan `≤ 100`):**
  - **Input 1 — MASA KERJA (`10 – 50`):**
    - _Baru_ (Turun Linear): `10 – 20 – 40`
    - _Lama_ (Naik Linear): `20 – 40 – 50`
  - **Input 2 — NILAI KINERJA (`40 – 90`):**
    - _Kurang_ (Turun Linear): `40 – 50 – 65`
    - _Cukup_ (Segitiga): `55 – 65 – 80`
    - _Baik_ (Naik Linear): `65 – 85 – 90`
  - **Output — SKOR BONUS (`10 – 90` — Monoton):**
    - _Kecil_ (Monoton Turun): `10 – 20 – 80`
    - _Besar_ (Monoton Naik): `20 – 80 – 90`
- **Visual Slide:** Tampilkan grafik gabungan **MASA KERJA**, **NILAI KINERJA** (menyorot **Segitiga**), dan **SKOR BONUS**.

---

## Slide 7: Studi Kasus 1 — Fuzzifikasi & Evaluasi Aturan (`IF` – `THEN`)

- **Judul Slide:** Studi Kasus 1: Proses Fuzzifikasi & Inferensi Aturan
- **1. Hasil Fuzzifikasi (`Masa Kerja = 25`, `Nilai Kinerja = 70`):**
  - **MASA KERJA (`25`):** _Baru_ = **`0.75`**, _Lama_ = **`0.25`**
  - **NILAI KINERJA (`70`):** _Kurang_ = **`0`**, _Cukup_ (segitiga sisi turun) = **`0.67`**, _Baik_ = **`0.25`**
- **2. Evaluasi 6 Aturan Utama (`AND` $\rightarrow$ ambil terkecil) & Nilai `THEN`:**
  - **R1:** `IF Baru AND Kurang THEN Bonus Kecil` $\rightarrow$ IF = `0` | THEN = `80`
  - **R2:** `IF Baru AND Cukup THEN Bonus Kecil` $\rightarrow$ IF = **`0.67`** | THEN = `80 - (0.67 x 60)` = **`39.8`**
  - **R3:** `IF Baru AND Baik THEN Bonus Besar` $\rightarrow$ IF = **`0.25`** | THEN = `20 + (0.25 x 60)` = **`35`**
  - **R4:** `IF Lama AND Kurang THEN Bonus Kecil` $\rightarrow$ IF = `0` | THEN = `80`
  - **R5:** `IF Lama AND Cukup THEN Bonus Besar` $\rightarrow$ IF = **`0.25`** | THEN = `20 + (0.25 x 60)` = **`35`**
  - **R6:** `IF Lama AND Baik THEN Bonus Besar` $\rightarrow$ IF = **`0.25`** | THEN = `20 + (0.25 x 60)` = **`35`**

---

## Slide 8: Studi Kasus 1 — Defuzzifikasi & Kesimpulan Akhir

- **Judul Slide:** Studi Kasus 1: Defuzzifikasi & Keputusan Akhir
- **Perhitungan Rata-Rata Berbobot (Fokus pada 4 Aturan Aktif):**
  - **Total Pembilang (`Hasil IF x Hasil THEN`):**
    `(0.67 x 39.8) + (0.25 x 35) + (0.25 x 35) + (0.25 x 35)`
    `= 26.666 + 8.75 + 8.75 + 8.75` = **`52.916`**
  - **Total Penyebut (`Total Hasil IF`):**
    `0.67 + 0.25 + 0.25 + 0.25` = **`1.42`**
  - **Hasil Akhir:** `52.916 / 1.42` = **`37.26`**
- **Kesimpulan Kasus 1:** Karyawan dengan masa kerja `25` bulan dan nilai kinerja `70` memperoleh **Skor Bonus = `37.26`**.

---

## Slide 9: Studi Kasus 2 — Setup & Kurva (Fokus Kurva Trapesium)

- **Judul Slide:** Studi Kasus 2: Pengaturan Kecepatan Kipas AC Otomatis
- **Soal Cerita Singkat:**
  Sistem AC otomatis mengatur **KECEPATAN KIPAS** (`20 – 80` RPM) berdasarkan **SUHU RUANGAN** (`15 – 40` °C) dan **JUMLAH ORANG** (`10 – 60` orang). Kasus uji: **Suhu Ruangan = `30` °C** dan **Jumlah Orang = `45` orang**.
- **Tabel Variabel & Batasan (Angka Puluhan `≤ 100`):**
  - **Input 1 — SUHU RUANGAN (`15 – 40`):**
    - _Sejuk_ (Turun Linear): `15 – 20 – 35`
    - _Panas_ (Naik Linear): `20 – 35 – 40`
  - **Input 2 — JUMLAH ORANG (`10 – 60`):**
    - _Sepi_ (Turun Linear): `10 – 15 – 30`
    - _Sedang_ (Trapesium): `20 – 30 – 40 – 50`
    - _Ramai_ (Naik Linear): `40 – 55 – 60`
  - **Output — KECEPATAN KIPAS (`20 – 80` — Monoton):**
    - _Pelan_ (Monoton Turun): `20 – 30 – 70`
    - _Cepat_ (Monoton Naik): `30 – 70 – 80`
- **Visual Slide:** Tampilkan grafik gabungan **SUHU RUANGAN**, **JUMLAH ORANG** (menyorot dataran puncak **Trapesium**), dan **KECEPATAN KIPAS**.

---

## Slide 10: Studi Kasus 2 — Ringkasan Matriks Perhitungan & Hasil Akhir

- **Judul Slide:** Studi Kasus 2: Ringkasan Hasil & Kesimpulan
- **1. Ringkasan Fuzzifikasi (`Suhu = 30`, `Jumlah Orang = 45`):**
  - **SUHU (`30`):** _Sejuk_ = **`0.33`**, _Panas_ = **`0.67`**
  - **JUMLAH ORANG (`45`):** _Sepi_ = **`0`**, _Sedang_ (trapesium sisi turun) = **`0.50`**, _Ramai_ = **`0.33`**
- **2. Ringkasan Aturan Aktif (`AND` / `OR`) & Nilai `THEN`:**
  - **R2 (`Sejuk AND Sedang -> Pelan`):** IF = `0.33` | THEN = `70 - (0.33 x 40)` = **`56.8`**
  - **R3 (`Sejuk AND Ramai -> Cepat`):** IF = `0.33` | THEN = `30 + (0.33 x 40)` = **`43.2`**
  - **R5 (`Panas AND Sedang -> Cepat`):** IF = `0.50` | THEN = `30 + (0.50 x 40)` = **`50.0`**
  - **R6 (`Panas AND Ramai -> Cepat`):** IF = `0.33` | THEN = `30 + (0.33 x 40)` = **`43.2`**
  - **R7 Tambahan (`Panas OR Ramai -> Cepat`):** IF = `0.67` | THEN = `30 + (0.67 x 40)` = **`56.8`**
- **3. Hasil Defuzzifikasi & Kesimpulan:**
  - `Total Pembilang = 18.744 + 14.256 + 25.0 + 14.256 + 38.056` = **`110.312`**
  - `Total Penyebut = 0.33 + 0.33 + 0.50 + 0.33 + 0.67` = **`2.16`**
  - **Kecepatan Kipas Akhir:** `110.312 / 2.16` = **`51.07` RPM** (dibulatkan **`51` RPM**).

---

## Slide 11: Kesimpulan & Penutup

- **Judul Slide:** Kesimpulan Presentasi
- **Poin Penutup:**
  - Metode Fuzzy Tsukamoto efektif memetakan keputusan berbasis aturan pakar (`AND`/`OR`) selama kurva output dibuat monoton.
  - Dari dua studi kasus (kurva segitiga maupun trapesium), alur perhitungan berjalan konsisten hingga menghasilkan rekomendasi nilai tegas yang terukur.
  - **Sesi Tanya Jawab & Diskusi.**

# Catatan Lengkap Metode Fuzzy Mamdani (Studi Kasus & Panduan Manual)

Catatan ini disusun dengan struktur yang sistematis agar mudah dipahami dan dibandingkan secara langsung dengan metode **Tsukamoto**, lengkap dengan penjabaran kondisi khusus yang sering muncul saat menghitung defuzzifikasi *Centroid* secara manual.

---

## 1. Perbedaan Prinsip Dasar (Mamdani vs. Tsukamoto)

Sebelum masuk ke studi kasus, pahami dulu perbedaan alur pengerjaannya:

| Tahapan | Metode Tsukamoto | Metode Mamdani |
| :--- | :--- | :--- |
| **Bentuk Kurva Output** | Wajib **monoton** (hanya linear naik atau linear turun). | **Bebas** (boleh segitiga, trapesium, maupun bahu/linear naik-turun). |
| **Implikasi Tiap Aturan** | Mencari nilai tegas ($z_i$) langsung pada tiap aturan dari nilai $\alpha_i$. | **Memotong kurva output** secara mendatar setinggi nilai $\alpha_i$ (fungsi `MIN`). |
| **Penggabungan Aturan** | Tidak ada penggabungan kurva (langsung ke rata-rata terbobot). | **Agregasi (`MAX`)**: menggabungkan semua potongan kurva menjadi satu daerah arsiran terluar. |
| **Defuzzifikasi** | Rata-rata terbobot: $Z^* = \frac{\sum (\alpha_i \cdot z_i)}{\sum \alpha_i}$ | **Centroid (Titik Berat)**: $Z^* = \frac{\sum (A_i \cdot x_i)}{\sum A_i}$ (memecah area menjadi bangun datar). |

---

## 2. Studi Kasus: Penentuan Jumlah Produksi Barang

Sebuah perusahaan manufaktur ingin menentukan **Jumlah Produksi** optimal berdasarkan dua variabel input, yaitu **Permintaan** dan **Persediaan**:
* **Permintaan:** turun (`1000` kemasan/hari) hingga naik (`5000` kemasan/hari).
* **Persediaan:** sedikit (`100` kemasan/hari) hingga banyak (`600` kemasan/hari).
* **Produksi (Output):** berkurang (`2000` kemasan/hari), tetap (`4500` kemasan/hari), hingga bertambah (`7000` kemasan/hari). *(Catatan: Pada Mamdani kita tambahkan himpunan segitiga **Tetap** di tengah agar terlihat jelas cara pemotongan dan penggabungan kurvanya).*

**Pertanyaan Kasus:**
Berapa kemasan yang harus diproduksi jika **Permintaan = `4000`** dan **Persediaan = `300`**?

---

## 3. Langkah 1: Fuzzifikasi (Pembentukan Himpunan & Derajat Keanggotaan)

### A. Variabel Input 1: Permintaan (`x = 4000`)
* **Turun `[1000, 5000]` (Linear Turun):**
  $$\mu_{\text{Turun}}(4000) = \frac{5000 - 4000}{5000 - 1000} = \frac{1000}{4000} = \mathbf{0,25}$$
* **Naik `[1000, 5000]` (Linear Naik):**
  $$\mu_{\text{Naik}}(4000) = \frac{4000 - 1000}{5000 - 1000} = \frac{3000}{4000} = \mathbf{0,75}$$

### B. Variabel Input 2: Persediaan (`y = 300`)
* **Sedikit `[100, 600]` (Linear Turun):**
  $$\mu_{\text{Sedikit}}(300) = \frac{600 - 300}{600 - 100} = \frac{300}{500} = \mathbf{0,60}$$
* **Banyak `[100, 600]` (Linear Naik):**
  $$\mu_{\text{Banyak}}(300) = \frac{300 - 100}{600 - 100} = \frac{200}{500} = \mathbf{0,40}$$

### C. Variabel Output: Jumlah Produksi (`[2000 – 7000]`)
Pada metode Mamdani, variabel output dibagi menjadi 3 himpunan fuzzy:
1. **Berkurang (Linear Turun / Bahu Kiri `[2000, 4500]`):**
   * Bernilai `1` jika $z \le 2000$, turun lurus hingga `0` di $z = 4500$ (lebar sisi miring = $4500 - 2000 = \mathbf{2500}$).
2. **Tetap (Segitiga `[3000, 4500, 6000]`):**
   * Naik dari $a = 3000$ ke puncak $b = 4500$ (lebar miring kiri = $\mathbf{1500}$), lalu turun ke $c = 6000$ (lebar miring kanan = $\mathbf{1500}$).
3. **Bertambah (Linear Naik / Bahu Kanan `[4500, 7000]`):**
   * Naik lurus dari `0` di $z = 4500$ hingga `1` di $z = 7000$ (lebar sisi miring = $7000 - 4500 = \mathbf{2500}$).

---

## 4. Langkah 2: Inferensi (Evaluasi Aturan `MIN` & Agregasi `MAX`)

Misalkan terdapat **4 Aturan (Rules)** yang digunakan:

* **[R1]** `IF` Permintaan **Turun** `AND` Persediaan **Banyak** `THEN` Produksi **Berkurang**
  $$\alpha_1 = \min(0,25 \,;\, 0,40) = \mathbf{0,25}$$
* **[R2]** `IF` Permintaan **Turun** `AND` Persediaan **Sedikit** `THEN` Produksi **Tetap**
  $$\alpha_2 = \min(0,25 \,;\, 0,60) = \mathbf{0,25}$$
* **[R3]** `IF` Permintaan **Naik** `AND` Persediaan **Banyak** `THEN` Produksi **Tetap**
  $$\alpha_3 = \min(0,75 \,;\, 0,40) = \mathbf{0,40}$$
* **[R4]** `IF` Permintaan **Naik** `AND` Persediaan **Sedikit** `THEN` Produksi **Bertambah**
  $$\alpha_4 = \min(0,75 \,;\, 0,60) = \mathbf{0,60}$$

### Agregasi (`MAX`) per Himpunan Output
Jika ada lebih dari satu aturan yang memiliki output sama, ambil nilai **terbesarnya (`MAX`)**:
1. **Tinggi Potongan Kurva Berkurang:** $\max(0,25) = \mathbf{0,25}$
2. **Tinggi Potongan Kurva Tetap:** $\max(0,25 \,;\, 0,40) = \mathbf{0,40}$ *(R2 kalah oleh R3 karena 0,40 > 0,25)*
3. **Tinggi Potongan Kurva Bertambah:** $\max(0,60) = \mathbf{0,60}$

---

## 5. Langkah 3: Pemotongan Kurva & Pembagian Bangun Datar

Di sinilah inti metode Mamdani. Kita cari **titik-titik batas alas ($z$)** dari kiri ke kanan tempat garis atap berubah arah (miring ke datar, atau datar ke miring).

### Rumus Pencari Titik Batas Alas ($z$)
* **Garis Naik (dari $a$ ke $b$):** $z = a + (\alpha \times \text{lebar miring})$
* **Garis Turun (dari $b$ ke $c$):** $z = c - (\alpha \times \text{lebar miring})$

### Menentukan Titik Batas Alas dari Kiri (`2000`) ke Kanan (`7000`)
1. **$z_1 = 2000$:** Batas paling kiri semesta output.
2. **$z_2 = 3600$:** Titik pertemuan dataran **Berkurang (`0,25`)** dengan tanjakan kurva **Tetap `[3000, 4500, 6000]`**.
   * *Cara hitung:* Kapan sisi naik *Tetap* (awal `3000`, lebar miring `1500`) mencapai tinggi `0,25`?
   * $z_2 = 3000 + (0,25 \times 1500) = 3000 + 375 = \mathbf{3375}$
3. **$z_3 = 3600$:** Titik saat tanjakan kurva **Tetap** berhenti terpotong di tinggi **`0,40`**.
   * *Cara hitung:* $z_3 = 3000 + (0,40 \times 1500) = 3000 + 600 = \mathbf{3600}$
4. **$z_4 = 5400$:** Titik saat dataran atas **Tetap (`0,40`)** habis dan mulai turun ke kanan (atau bertabrakan dengan tanjakan **Bertambah**).
   * *Cek sisi turun Tetap di tinggi `0,40`:* $6000 - (0,40 \times 1500) = \mathbf{5400}$.
   * *Cek sisi naik Bertambah `[4500, 7000]` di tinggi `0,40`:* $4500 + (0,40 \times 2500) = \mathbf{5500}$.
   * Karena `5400 < 5500`, maka pada rentang `5400 – 5500` terjadi **lekukan "V" dangkal** (perpotongan sisi turun *Tetap* dan sisi naik *Bertambah* di titik $z \approx \mathbf{5462}$ pada ketinggian $\alpha \approx \mathbf{0,385}$).
   * *(Catatan Praktis: Untuk hitungan manual yang cepat, lekukan dangkal `5400–5500` ini sering disederhanakan menjadi dataran rata langsung sampai **`5500`**—selisih akhirnya sangat kecil).*
5. **$z_5 = 6000$:** Titik saat tanjakan kurva **Bertambah `[4500, 7000]`** berhenti terpotong di tinggi **`0,60`**.
   * *Cara hitung:* $z_5 = 4500 + (0,60 \times 2500) = 4500 + 1500 = \mathbf{6000}$
6. **$z_6 = 7000$:** Batas paling kanan semesta output (datar di tinggi `0,60`).

---

## 6. Langkah 4: Defuzzifikasi Centroid (Menghitung Luas & Titik Tengah)

Dengan menggunakan **penyederhanaan manual** (meratakan dataran `0,40` langsung hingga titik `5500`), daerah arsiran terbagi menjadi **6 bangun datar sederhana** (4 Persegi Panjang + 2 Segitiga Siku-Siku):

### Rumus Luas ($A$) dan Titik Tengah ($x$)
* **Persegi Panjang:**
  * $A = (z_{\text{akhir}} - z_{\text{awal}}) \times t$
  * $x = \frac{z_{\text{awal}} + z_{\text{akhir}}}{2}$
* **Segitiga Siku-Siku (Sisi Naik):**
  * $A = \frac{1}{2} \times (z_{\text{akhir}} - z_{\text{awal}}) \times t$
  * $x = z_{\text{awal}} + \frac{2}{3} \times (z_{\text{akhir}} - z_{\text{awal}})$

### Tabel Perhitungan Lengkap

| Bagian | Rentang Alas & Bentuk Bangun | Perhitungan Luas ($A_i$) | Perhitungan Titik Tengah ($x_i$) | Momen ($A_i \times x_i$) |
| :--- | :--- | :--- | :--- | :--- |
| **A1** | `2000 – 3375`<br>Persegi Panjang ($t = 0,25$) | $(3375 - 2000) \times 0,25$<br>$= 1375 \times 0,25 = \mathbf{343,75}$ | $\frac{2000 + 3375}{2} = \mathbf{2687,50}$ | $343,75 \times 2687,50$<br>$= \mathbf{923.828,13}$ |
| **A2** | `3375 – 3600`<br>Persegi Panjang Bawah ($t = 0,25$) | $(3600 - 3375) \times 0,25$<br>$= 225 \times 0,25 = \mathbf{56,25}$ | $\frac{3375 + 3600}{2} = \mathbf{3487,50}$ | $56,25 \times 3487,50$<br>$= \mathbf{196.171,88}$ |
| **A3** | `3375 – 3600`<br>Segitiga Siku-Siku Atas ($t = 0,40 - 0,25 = 0,15$) | $0,5 \times 225 \times 0,15$<br>$= \mathbf{16,88}$ | $3375 + \frac{2}{3}(225)$<br>$= 3375 + 150 = \mathbf{3525,00}$ | $16,88 \times 3525,00$<br>$= \mathbf{59.502,00}$ |
| **A4** | `3600 – 5500`<br>Persegi Panjang ($t = 0,40$) | $(5500 - 3600) \times 0,40$<br>$= 1900 \times 0,40 = \mathbf{760,00}$ | $\frac{3600 + 5500}{2} = \mathbf{4550,00}$ | $760,00 \times 4550,00$<br>$= \mathbf{3.458.000,00}$ |
| **A5** | `5500 – 6000`<br>Persegi Panjang Bawah ($t = 0,40$) | $(6000 - 5500) \times 0,40$<br>$= 500 \times 0,40 = \mathbf{200,00}$ | $\frac{5500 + 6000}{2} = \mathbf{5750,00}$ | $200,00 \times 5750,00$<br>$= \mathbf{1.150.000,00}$ |
| **A6** | `5500 – 6000`<br>Segitiga Siku-Siku Atas ($t = 0,60 - 0,40 = 0,20$) | $0,5 \times 500 \times 0,20$<br>$= \mathbf{50,00}$ | $5500 + \frac{2}{3}(500)$<br>$= 5500 + 333,33 = \mathbf{5833,33}$ | $50,00 \times 5833,33$<br>$= \mathbf{291.666,50}$ |
| **A7** | `6000 – 7000`<br>Persegi Panjang ($t = 0,60$) | $(7000 - 6000) \times 0,60$<br>$= 1000 \times 0,60 = \mathbf{600,00}$ | $\frac{6000 + 7000}{2} = \mathbf{6500,00}$ | $600,00 \times 6500,00$<br>$= \mathbf{3.900.000,00}$ |

### Hasil Akhir ($Z^*$)
* **Total Luas ($\sum A_i$):**
  $$343,75 + 56,25 + 16,88 + 760,00 + 200,00 + 50,00 + 600,00 = \mathbf{2.026,88}$$
* **Total Momen ($\sum A_i \times x_i$):**
  $$923.828,13 + 196.171,88 + 59.502,00 + 3.458.000,00 + 1.150.000,00 + 291.666,50 + 3.900.000,00 = \mathbf{9.979.168,51}$$
* **Keputusan Produksi Optimal:**
  $$Z^* = \frac{9.979.168,51}{2.026,88} \approx \mathbf{4.923,41} \text{ kemasan}$$

---

## 7. Kondisi-Kondisi Khusus yang Sering Muncul di Mamdani

Saat mengerjakan soal Mamdani secara manual, perhatikan **4 kondisi khusus** berikut agar tidak bingung saat membagi bangun datar:

1. **Kondisi 1: Banyak Aturan (misal 6 atau 9 Rules) dengan Himpunan Output Terbatas**
   * Semua aturan yang mengarah ke himpunan output yang sama **tidak dipotong berkali-kali**, melainkan digabung terlebih dahulu dengan operator **`MAX`** sehingga tiap himpunan output hanya memiliki **1 tinggi potongan final**.
2. **Kondisi 2: Ada Himpunan Output yang Bernilai $\alpha = 0$**
   * Jika sebuah himpunan output memiliki hasil agregasi `0` (misalnya di ujung kiri atau kanan), maka rentang kurva tersebut **diabaikan/kosong** (luas = `0`) dan batas perhitungan langsung berhenti pada ujung kurva aktif terakhir.
3. **Kondisi 3: Sisi Miring Terjepit di Antara Dua Dataran (Trapesium Miring)**
   * Bagian yang atapnya miring dan tidak menyentuh sumbu dasar ($y = 0$) selalu dipecah secara mendatar menjadi **2 bangun**: **Persegi Panjang** di bagian bawah (setinggi dataran terendah) dan **Segitiga Siku-Siku** di bagian atas (setinggi selisih tinggi kedua dataran).
4. **Kondisi 4: Munculnya Lekukan Huruf "V" (Perpotongan Dua Sisi Miring)**
   * Terjadi ketika satu kurva sedang turun miring sementara kurva sebelahnya mulai naik miring sebelum mencapai tinggi potongan masing-masing.
   * **Jika mengejar presisi komputer (MATLAB):** Titik tabrakan kedua garis miring harus dicari dan lekukan "V" dipecah mengikuti garis aslinya.
   * **Jika menghitung manual di kertas:** Lekukan "V" yang dangkal boleh **disederhanakan/diratakan** menyatu dengan dataran di sebelahnya agar jumlah bangun datar lebih sedikit dan perhitungan lebih cepat.
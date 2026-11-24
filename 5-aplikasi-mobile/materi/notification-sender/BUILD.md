# Cara Build & Deploy — FCM Sender (PHP · Python · Golang)

Panduan ini menjelaskan cara mengubah kode menjadi aplikasi yang siap dijalankan
tanpa editor/IDE: **PHP** (tanpa kompilasi, tinggal deploy), **Python** (server produksi + `.exe`),
dan **Golang** (satu file binary untuk Windows/Linux/macOS).

> Semua perintah dijalankan dari folder `notification_sender/`.
> Pastikan `service-account.json` sudah ada (lihat `README.md` Langkah 2).

## Ringkasan

| Bahasa | Perlu kompilasi? | Hasil build | Yang harus ikut dibawa saat deploy |
|---|---|---|---|
| PHP | ❌ Tidak (bahasa interpreted) | Folder `php/` apa adanya | `index.php`, `.htaccess`, `service-account.json` |
| Python | ❌ Tidak wajib — opsional dibundel jadi `.exe` | `dist/fcmsender.exe` (± 17 MB) | Versi source: folder `python/` + `pip install`. Versi exe: `fcmsender.exe` + `service-account.json` |
| Golang | ✅ Ya (compiled) | `fcmsender.exe` / `fcmsender-linux` (± 8 MB) | Hanya **1 file binary** + `service-account.json` (HTML sudah tertanam di dalam binary) |

Lokasi `service-account.json` yang dicari program (berurutan):
1. Folder program itu sendiri (`php/`, `python/`, `golang/`, atau folder tempat file `.exe` berada).
2. Folder induknya (`notification_sender/`).

---

## 1. PHP

PHP **tidak perlu di-build**. Kode langsung dijalankan oleh interpreter PHP. "Build" di sini artinya:
cek sintaks → siapkan server → deploy.

### 1.1 Cek sintaks (pengganti "compile")
```bash
php -l php/index.php
# Output: No syntax errors detected in php/index.php
```

### 1.2 Cek ekstensi yang dibutuhkan
```bash
php -m | findstr /I "curl openssl"      # Windows (CMD/PowerShell)
php -m | grep -iE "curl|openssl"        # Linux / macOS / Git Bash
```
Jika tidak muncul, buka `php.ini` lalu hilangkan tanda `;` pada baris:
```ini
extension=curl
extension=openssl
```

### 1.3 Menjalankan
**a) Server bawaan PHP (untuk praktik/development)**
```bash
cd php
php -S localhost:8000
```

**b) Laragon / XAMPP (Apache)**
1. Salin seluruh folder `notification_sender/` ke `C:\laragon\www\` (Laragon) atau `C:\xampp\htdocs\` (XAMPP).
2. Start Apache.
3. Buka `http://localhost/notification_sender/php/`

**c) Hosting / VPS Linux (Apache)**
```bash
# upload folder, contoh ke /var/www/html/notification_sender
sudo apt install php php-curl            # openssl sudah termasuk di paket php
sudo systemctl restart apache2
```

### 1.4 ⚠️ Lindungi `service-account.json`
Jika folder ditaruh di dalam *web root* (`www/`, `htdocs/`, `/var/www/html`), file JSON bisa
**diunduh siapa saja** lewat `http://.../service-account.json`. Sudah disediakan file `.htaccess`
(di `notification_sender/` dan `php/`) yang memblokir semua file `.json` di Apache:

```apache
<FilesMatch "\.(json)$">
    Require all denied
</FilesMatch>
```

Uji: buka `http://localhost/notification_sender/service-account.json` → harus muncul **403 Forbidden**.

- Laragon/XAMPP (Apache) membaca `.htaccess` secara otomatis.
- Jika memakai **Nginx**, `.htaccess` tidak berlaku. Tambahkan di konfigurasi server:
  ```nginx
  location ~* \.json$ { deny all; }
  ```
- Cara paling aman: taruh `service-account.json` **di luar** web root, lalu ubah konstanta
  `SERVICE_ACCOUNT_FILE` di `index.php` ke path tersebut.

---

## 2. Python

Python juga *interpreted*. Ada 2 cara "build":
**(A)** siapkan virtual environment + server produksi, atau **(B)** bundel menjadi **satu file `.exe`** dengan PyInstaller.

### 2.1 Siapkan Virtual Environment (wajib untuk kedua cara)
```bash
cd python
python -m venv venv
```
Aktifkan:
```bash
venv\Scripts\activate          # Windows (CMD/PowerShell)
source venv/Scripts/activate   # Windows (Git Bash)
source venv/bin/activate       # Linux / macOS
```
Install dependency:
```bash
pip install -r requirements.txt
```

### 2.2 Cara A — Jalankan dengan Server Produksi
`python app.py` memakai server development Flask (ada mode debug, **tidak untuk produksi**).
Untuk produksi gunakan:

**Windows → waitress**
```bash
waitress-serve --listen=127.0.0.1:5000 app:app
```

**Linux → gunicorn**
```bash
pip install gunicorn
gunicorn -w 2 -b 127.0.0.1:5000 app:app
```
`app:app` artinya: file `app.py`, variabel Flask bernama `app`.

### 2.3 Cara B — Build menjadi `.exe` (PyInstaller)
```bash
pip install pyinstaller
pyinstaller --onefile --name fcmsender --add-data "templates;templates" app.py
```
- `--onefile` → semua dibundel menjadi 1 file.
- `--add-data "templates;templates"` → ikutkan folder template HTML.
  (Di **Linux/macOS** pemisahnya titik dua: `--add-data "templates:templates"`.)

Hasil: `python/dist/fcmsender.exe` (± 17 MB, Python sudah ikut di dalamnya).

**Deploy / menjalankan:**
```
folder_bebas/
├── fcmsender.exe
└── service-account.json      ← taruh di sebelah .exe
```
Klik dua kali `fcmsender.exe` (atau jalankan dari terminal), lalu buka http://localhost:5000.
Versi `.exe` otomatis memakai server produksi **waitress** (tanpa debug).
Ganti port: `set PORT=5001` (CMD) / `$env:PORT=5001` (PowerShell) sebelum menjalankan.

> Catatan PyInstaller:
> - `.exe` hanya berjalan di OS tempat ia di-build (build di Windows → hanya untuk Windows).
> - Antivirus kadang mencurigai `.exe` PyInstaller (*false positive*). Tambahkan pengecualian bila perlu.
> - Folder `build/`, `dist/`, dan file `fcmsender.spec` adalah hasil build — boleh dihapus.

---

## 3. Golang

Go adalah bahasa *compiled*: hasil build berupa **1 file binary** yang bisa langsung dijalankan
tanpa menginstal Go di komputer tujuan. File `index.html` sudah **tertanam** di dalam binary
(lewat `//go:embed`), jadi tidak perlu ikut disalin.

### 3.1 Build untuk Windows
```bash
cd golang
go build -o fcmsender.exe .
```

Versi lebih kecil (membuang info debug):
```bash
go build -ldflags "-s -w" -o fcmsender.exe .
```

### 3.2 Cross-compile (build untuk OS lain dari Windows)
Go bisa membuat binary untuk OS lain cukup dengan mengatur `GOOS` dan `GOARCH`.

**PowerShell:**
```powershell
$env:GOOS="linux";   $env:GOARCH="amd64"; go build -ldflags "-s -w" -o fcmsender-linux .
$env:GOOS="darwin";  $env:GOARCH="arm64"; go build -ldflags "-s -w" -o fcmsender-mac .
$env:GOOS="windows"; $env:GOARCH="amd64"   # kembalikan ke Windows
```

**Git Bash / Linux / macOS:**
```bash
GOOS=linux  GOARCH=amd64 go build -ldflags "-s -w" -o fcmsender-linux .
GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o fcmsender-mac .
```

| Target | GOOS | GOARCH |
|---|---|---|
| Windows 64-bit | `windows` | `amd64` |
| Linux server / VPS | `linux` | `amd64` |
| Linux ARM (Raspberry Pi 4/5, STB) | `linux` | `arm64` |
| macOS Apple Silicon (M1–M4) | `darwin` | `arm64` |
| macOS Intel | `darwin` | `amd64` |

### 3.3 Deploy / menjalankan
```
folder_bebas/
├── fcmsender.exe   (atau fcmsender-linux)
└── service-account.json
```
**Windows:** klik dua kali `fcmsender.exe` → buka http://localhost:8080

**Linux:**
```bash
chmod +x fcmsender-linux
./fcmsender-linux
```

Ganti alamat/port lewat environment variable `ADDR`:
```bash
ADDR=127.0.0.1:9090 ./fcmsender-linux        # Linux / Git Bash
```
```powershell
$env:ADDR="127.0.0.1:9090"; .\fcmsender.exe  # PowerShell
```

### 3.4 (Opsional) Jalan otomatis sebagai service di Linux (systemd)
`/etc/systemd/system/fcmsender.service`:
```ini
[Unit]
Description=FCM Sender (Golang)
After=network.target

[Service]
WorkingDirectory=/opt/fcmsender
ExecStart=/opt/fcmsender/fcmsender-linux
Restart=always

[Install]
WantedBy=multi-user.target
```
```bash
sudo systemctl daemon-reload
sudo systemctl enable --now fcmsender
```

---

## 4. ⚠️ Keamanan Saat Deploy

Web ini **tidak punya login**. Siapa pun yang bisa membuka halamannya dapat mengirim notifikasi
ke semua pengguna aplikasi. Karena itu:

1. Secara default ketiga versi hanya menerima akses dari komputer sendiri (`127.0.0.1` / `localhost`). **Pertahankan ini** untuk praktik.
2. Jika ingin dibuka dari jaringan (misalnya `ADDR=0.0.0.0:8080`), tambahkan login atau batasi IP lebih dulu.
3. `service-account.json` = kunci rahasia. Jangan di-commit ke Git (sudah ada `.gitignore`), jangan diunggah ke tempat publik, dan jangan dibundel ke dalam `.exe` yang dibagikan ke orang lain.

---

## 5. Perbandingan Hasil Build

| | PHP | Python (`.exe`) | Golang |
|---|---|---|---|
| Ukuran | ± 10 KB (source) | ± 17 MB | ± 8 MB |
| Butuh runtime di komputer tujuan? | Ya (PHP + Apache/Nginx) | Tidak | Tidak |
| Library tambahan | Tidak ada | flask, google-auth, requests, waitress | Tidak ada |
| Cross-platform build | Tidak perlu | Harus build di tiap OS | Bisa dari 1 komputer (`GOOS`/`GOARCH`) |
| Cocok untuk | Shared hosting, Laragon/XAMPP | Desktop Windows tanpa Python | Server/VPS, desktop, Raspberry Pi |

"""
Web Pengirim Notifikasi FCM (HTTP v1) - versi Python (Flask)

Install : pip install -r requirements.txt
Jalankan: python app.py
Buka    : http://localhost:5000

Syarat: file service-account.json (dari Firebase Console) ada di folder ini.
"""
import json
import os
import sys
from pathlib import Path

import requests
from flask import Flask, render_template, request
from google.auth.transport.requests import Request
from google.oauth2 import service_account

# Cari di folder ini dulu, jika tidak ada pakai yang di folder induk (notification_sender/)
# Jika dijalankan sebagai .exe hasil PyInstaller, "folder ini" = folder tempat .exe berada.
FROZEN = getattr(sys, "frozen", False)
_HERE = Path(sys.executable).resolve().parent if FROZEN else Path(__file__).resolve().parent
SERVICE_ACCOUNT_FILE = next(
    (p for p in (_HERE / "service-account.json", _HERE.parent / "service-account.json") if p.exists()),
    _HERE / "service-account.json",
)
FCM_SCOPE = "https://www.googleapis.com/auth/firebase.messaging"
CHANNEL_ID = "high_importance_channel"

app = Flask(__name__)
_credentials = None  # disimpan agar access token dipakai ulang selama masih berlaku


def get_access_token() -> tuple[str, str]:
    """LANGKAH 1: Ambil Access Token OAuth 2.0 dari Service Account.
    Library google-auth yang membuat & menandatangani JWT untuk kita."""
    global _credentials
    if not SERVICE_ACCOUNT_FILE.exists():
        raise RuntimeError("File service-account.json tidak ditemukan di folder python/ maupun notification_sender/.")
    if _credentials is None:
        _credentials = service_account.Credentials.from_service_account_file(
            str(SERVICE_ACCOUNT_FILE), scopes=[FCM_SCOPE]
        )
    if not _credentials.valid:  # belum ada token / sudah kedaluwarsa
        _credentials.refresh(Request())
    return _credentials.token, _credentials.project_id


def parse_data(text: str) -> dict:
    """Ubah teks 'key=value' per baris menjadi dict (nilai selalu string)."""
    data = {}
    for line in text.splitlines():
        if "=" in line:
            key, value = line.split("=", 1)
            if key.strip():
                data[key.strip()] = value.strip()
    return data


def send_notification(form: dict) -> dict:
    """LANGKAH 2: Kirim pesan ke FCM HTTP v1."""
    access_token, project_id = get_access_token()

    message = {
        form["target_type"]: form["target"].strip(),  # "token" atau "topic"
        "notification": {"title": form["title"], "body": form["body"]},
        "android": {
            "priority": "high",
            "notification": {"channel_id": CHANNEL_ID},
        },
    }
    if form["image"].strip():
        message["notification"]["image"] = form["image"].strip()
    data = parse_data(form["data"])
    if data:
        message["data"] = data
    payload = {"message": message}

    resp = requests.post(
        f"https://fcm.googleapis.com/v1/projects/{project_id}/messages:send",
        headers={
            "Authorization": f"Bearer {access_token}",
            "Content-Type": "application/json; charset=utf-8",
        },
        json=payload,
        timeout=30,
    )
    return {
        "status": resp.status_code,
        "payload": json.dumps(payload, indent=2, ensure_ascii=False),
        "response": resp.text,
    }


@app.route("/", methods=["GET", "POST"])
def index():
    form = {
        "target_type": request.form.get("target_type", "token"),
        "target": request.form.get("target", ""),
        "title": request.form.get("title", "Halo Mahasiswa"),
        "body": request.form.get("body", "Notifikasi dari web Python"),
        "image": request.form.get("image", ""),
        "data": request.form.get("data", "kategori=pengumuman\nid_pesan=001"),
    }
    result, error = None, None
    if request.method == "POST":
        if form["target_type"] not in ("token", "topic"):
            form["target_type"] = "token"
        try:
            result = send_notification(form)
        except Exception as e:  # tampilkan error apa pun ke halaman
            error = str(e)
    return render_template("index.html", form=form, result=result, error=error)


if __name__ == "__main__":
    port = int(os.environ.get("PORT", "5000"))
    if FROZEN:
        # Versi .exe: pakai server produksi waitress (tanpa debug)
        from waitress import serve

        print(f"FCM Sender (Python) berjalan di http://localhost:{port}", flush=True)
        serve(app, host="127.0.0.1", port=port)
    else:
        # Mode development
        app.run(host="127.0.0.1", port=port, debug=True)

// Web Pengirim Notifikasi FCM (HTTP v1) - versi Golang
// Hanya memakai standard library (tanpa package tambahan).
//
// Jalankan: go run .
// Buka    : http://localhost:8080
//
// Syarat: file service-account.json (dari Firebase Console) ada di folder ini.
package main

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	serviceAccountFile = "service-account.json"
	fcmScope           = "https://www.googleapis.com/auth/firebase.messaging"
	channelID          = "high_importance_channel"
)

//go:embed index.html
var indexHTML string
var tmpl = template.Must(template.New("index").Parse(indexHTML))

// ServiceAccount berisi field penting dari service-account.json
type ServiceAccount struct {
	ProjectID   string `json:"project_id"`
	ClientEmail string `json:"client_email"`
	PrivateKey  string `json:"private_key"`
	TokenURI    string `json:"token_uri"`
}

// Cache access token agar tidak meminta token baru setiap kirim.
var (
	tokenMu     sync.Mutex
	cachedToken string
	tokenExpiry time.Time
)

func loadServiceAccount() (*ServiceAccount, error) {
	// Urutan pencarian: folder kerja, folder induk, lalu folder tempat file .exe berada
	candidates := []string{serviceAccountFile, filepath.Join("..", serviceAccountFile)}
	if exe, e := os.Executable(); e == nil {
		candidates = append(candidates, filepath.Join(filepath.Dir(exe), serviceAccountFile))
	}
	var raw []byte
	var err error
	for _, path := range candidates {
		if raw, err = os.ReadFile(path); err == nil {
			break
		}
	}
	if err != nil {
		return nil, errors.New("file service-account.json tidak ditemukan di folder golang/ maupun notification_sender/")
	}
	var sa ServiceAccount
	if err := json.Unmarshal(raw, &sa); err != nil {
		return nil, fmt.Errorf("service-account.json tidak valid: %w", err)
	}
	return &sa, nil
}

func b64url(b []byte) string { return base64.RawURLEncoding.EncodeToString(b) }

// getAccessToken - LANGKAH 1: buat JWT, tanda tangani (RS256),
// lalu tukar ke Google untuk mendapatkan access token OAuth 2.0.
func getAccessToken(sa *ServiceAccount) (string, error) {
	tokenMu.Lock()
	defer tokenMu.Unlock()
	if cachedToken != "" && time.Now().Before(tokenExpiry) {
		return cachedToken, nil
	}

	block, _ := pem.Decode([]byte(sa.PrivateKey))
	if block == nil {
		return "", errors.New("private_key tidak bisa dibaca (format PEM)")
	}
	parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("private_key tidak valid: %w", err)
	}
	key, ok := parsed.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("private_key bukan RSA")
	}

	now := time.Now().Unix()
	header, _ := json.Marshal(map[string]string{"alg": "RS256", "typ": "JWT"})
	claims, _ := json.Marshal(map[string]any{
		"iss":   sa.ClientEmail,
		"scope": fcmScope,
		"aud":   sa.TokenURI,
		"iat":   now,
		"exp":   now + 3600,
	})
	unsigned := b64url(header) + "." + b64url(claims)
	hash := sha256.Sum256([]byte(unsigned))
	sig, err := rsa.SignPKCS1v15(nil, key, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	jwt := unsigned + "." + b64url(sig)

	resp, err := http.PostForm(sa.TokenURI, url.Values{
		"grant_type": {"urn:ietf:params:oauth:grant-type:jwt-bearer"},
		"assertion":  {jwt},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var tok struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if resp.StatusCode != 200 || json.Unmarshal(body, &tok) != nil || tok.AccessToken == "" {
		return "", fmt.Errorf("gagal mendapatkan access token (%d): %s", resp.StatusCode, body)
	}
	cachedToken = tok.AccessToken
	tokenExpiry = time.Now().Add(time.Duration(tok.ExpiresIn-60) * time.Second)
	return cachedToken, nil
}

// parseData mengubah teks "key=value" per baris menjadi map (nilai string).
func parseData(text string) map[string]string {
	data := map[string]string{}
	for _, line := range strings.Split(text, "\n") {
		k, v, found := strings.Cut(line, "=")
		if found && strings.TrimSpace(k) != "" {
			data[strings.TrimSpace(k)] = strings.TrimSpace(v)
		}
	}
	return data
}

type Form struct {
	TargetType, Target, Title, Body, Image, Data string
}

type Result struct {
	Status            int
	Payload, Response string
}

// sendNotification - LANGKAH 2: kirim pesan ke FCM HTTP v1.
func sendNotification(f Form) (*Result, error) {
	sa, err := loadServiceAccount()
	if err != nil {
		return nil, err
	}
	accessToken, err := getAccessToken(sa)
	if err != nil {
		return nil, err
	}

	notification := map[string]string{"title": f.Title, "body": f.Body}
	if img := strings.TrimSpace(f.Image); img != "" {
		notification["image"] = img
	}
	message := map[string]any{
		f.TargetType:   strings.TrimSpace(f.Target), // "token" atau "topic"
		"notification": notification,
		"android": map[string]any{
			"priority":     "high",
			"notification": map[string]string{"channel_id": channelID},
		},
	}
	if data := parseData(f.Data); len(data) > 0 {
		message["data"] = data
	}
	payload, _ := json.MarshalIndent(map[string]any{"message": message}, "", "  ")

	req, _ := http.NewRequest(http.MethodPost,
		"https://fcm.googleapis.com/v1/projects/"+sa.ProjectID+"/messages:send",
		bytes.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return &Result{Status: resp.StatusCode, Payload: string(payload), Response: string(body)}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	f := Form{
		TargetType: "token",
		Title:      "Halo Mahasiswa",
		Body:       "Notifikasi dari web Golang",
		Data:       "kategori=pengumuman\nid_pesan=001",
	}
	view := map[string]any{"Form": &f}

	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		f = Form{
			TargetType: r.FormValue("target_type"),
			Target:     r.FormValue("target"),
			Title:      r.FormValue("title"),
			Body:       r.FormValue("body"),
			Image:      r.FormValue("image"),
			Data:       r.FormValue("data"),
		}
		if f.TargetType != "token" && f.TargetType != "topic" {
			f.TargetType = "token"
		}
		if res, err := sendNotification(f); err != nil {
			view["Error"] = err.Error()
		} else {
			view["Result"] = res
		}
	}
	if err := tmpl.Execute(w, view); err != nil {
		log.Println("template error:", err)
	}
}

func main() {
	// Alamat bisa diganti lewat environment variable ADDR, contoh: ADDR=127.0.0.1:9090
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	http.HandleFunc("/", handler)
	log.Println("FCM Sender (Golang) berjalan di http://" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

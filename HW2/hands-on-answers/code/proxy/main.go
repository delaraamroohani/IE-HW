package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func b64(b []byte) string { return base64.RawURLEncoding.EncodeToString(b) }

func mac(msg string) string {
	m := hmac.New(sha256.New, secret)
	m.Write([]byte(msg))
	return b64(m.Sum(nil))
}

func sign(user string) string {
	h := b64([]byte(`{"alg":"HS256","typ":"JWT"}`))
	p, _ := json.Marshal(map[string]any{
		"sub": user,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	msg := h + "." + b64(p)
	return msg + "." + mac(msg)
}

func verify(tok string) bool {
	parts := strings.Split(tok, ".")
	if len(parts) != 3 {
		return false
	}
	hb, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}
	var hdr struct{ Alg string `json:"alg"` }
	if json.Unmarshal(hb, &hdr) != nil || hdr.Alg != "HS256" { // reject "none" etc.
		return false
	}
	if !hmac.Equal([]byte(mac(parts[0]+"."+parts[1])), []byte(parts[2])) {
		return false
	}
	pb, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}
	var c struct{ Exp int64 `json:"exp"` }
	if json.Unmarshal(pb, &c) != nil {
		return false
	}
	return c.Exp > time.Now().Unix()
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var in struct{ Username, Password string }
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	okU := subtle.ConstantTimeCompare([]byte(in.Username), []byte(os.Getenv("AUTH_USER"))) == 1
	okP := subtle.ConstantTimeCompare([]byte(in.Password), []byte(os.Getenv("AUTH_PASS"))) == 1
	if !okU || !okP {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": sign(in.Username)})
}

func main() {
	if len(secret) == 0 {
		log.Fatal("JWT_SECRET is required")
	}
	target, err := url.Parse(os.Getenv("UPSTREAM"))
	if err != nil {
		log.Fatal(err)
	}
	rp := httputil.NewSingleHostReverseProxy(target)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") || !verify(strings.TrimPrefix(h, "Bearer ")) {
			w.Header().Set("WWW-Authenticate", `Bearer realm="proxy"`)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		rp.ServeHTTP(w, r)
	})
	log.Println("proxy listening on :80")
	log.Fatal(http.ListenAndServe(":80", mux))
}
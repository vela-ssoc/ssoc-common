package preadtls

import (
	"crypto/tls"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestHTTP(t *testing.T) {
	pair, err := tls.X509KeyPair([]byte(publicKey), []byte(privateKey))
	if err != nil {
		t.Error(err)
		return
	}

	httpTCP := http.NewServeMux()
	httpTLS := http.NewServeMux()

	httpTCP.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello http"))
	})

	httpTLS.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello https"))
	})

	ln, err := net.Listen("tcp", ":8877")
	if err != nil {
		t.Fatal(err)
	}
	lis := NewListener(ln, time.Minute)

	srvTCP := &http.Server{Handler: httpTCP}
	srvTLS := &http.Server{Handler: httpTLS, TLSConfig: &tls.Config{Certificates: []tls.Certificate{pair}}}

	errs := make(chan error)
	go func() {
		errs <- srvTCP.Serve(lis.TCPListener())
	}()
	go func() {
		errs <- srvTLS.ServeTLS(lis.TLSListener(), "", "")
	}()

	err = <-errs
	t.Log(err)
}

const publicKey = `-----BEGIN CERTIFICATE-----
MIICCTCCAY+gAwIBAgIQAM3Rd0GaNCsCDUjNGfmd3zAKBggqhkjOPQQDAjAWMRQw
EgYDVQQDDAtleGFtcGxlLmNvbTAeFw0yNTEyMDMwNjIwNTJaFw0yNjEyMDMwNjIw
NTJaMBYxFDASBgNVBAMMC2V4YW1wbGUuY29tMHYwEAYHKoZIzj0CAQYFK4EEACID
YgAELz+4JRKBuNkKvsqdvy0v4HSSorcquJmStD/soj6dSwKSSjZrTgZoGONqW/cD
G3LcVh+UWrhoIdAXP4T0uSb9y7Ted+E2Lr7Q6fylx4iK+lViYrMFAeBcWkIuyfXi
cGOco4GhMIGeMB0GA1UdDgQWBBSTBAoW4wWUSP20mxaSBj0C1hbuADAOBgNVHQ8B
Af8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zA7BgNVHSUENDAyBggrBgEFBQcDAgYI
KwYBBQUHAwEGCCsGAQUFBwMDBggrBgEFBQcDBAYIKwYBBQUHAwgwHwYDVR0jBBgw
FoAUkwQKFuMFlEj9tJsWkgY9AtYW7gAwCgYIKoZIzj0EAwIDaAAwZQIwIbsC8NaS
BtfW2NAOkl2ytnMh8RdE39vQmvw+n9lQkEtZKQlt5ErHwF9OONMYfyxwAjEAzrOP
0R/p+aBMvuvFAc6uAAh5CKgXzg172jxlYcRlJvl3i3xPfm0mkG6kkvDVk2+/
-----END CERTIFICATE-----`

const privateKey = `-----BEGIN PRIVATE KEY-----
MIG/AgEAMBAGByqGSM49AgEGBSuBBAAiBIGnMIGkAgEBBDAdRkyPAkjFMg0no4N9
IardNg7t3bGA8JZyDvvYUxvtQeDoiDmJz4m5ROMrXCulwCqgBwYFK4EEACKhZANi
AAQvP7glEoG42Qq+yp2/LS/gdJKityq4mZK0P+yiPp1LApJKNmtOBmgY42pb9wMb
ctxWH5RauGgh0Bc/hPS5Jv3LtN534TYuvtDp/KXHiIr6VWJiswUB4FxaQi7J9eJw
Y5w=
-----END PRIVATE KEY-----`

package http

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const CONN_HOST = "localhost"
const CONN_PORT = "8081"
const ADMIN_USER = "admin"
const ADMIN_PASSWORD = "admin123"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
}

func Authentication(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, pwd, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pwd), []byte(ADMIN_PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", "Basic realm="+realm)
			w.WriteHeader(401)
			w.Write([]byte("authorized failed\n"))
			return
		}
		handler(w, r)
	}
}
func StartServer() {
	http.HandleFunc("/", Authentication(helloWorld, "Please enter username and password"))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server:", err)
		return
	}
}

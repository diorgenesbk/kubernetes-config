package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

var startedAt = time.Now()

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Success!")
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "Hello Success! User: %s, Password: %s.", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte("Tudo belezinha por aqui"))
	//duration := time.Since(startedAt)
	// if duration.Seconds() < 20 {
	// 	w.WriteHeader(500)
	// 	w.Write([]byte(fmt.Sprintf("Ainda não ta pronto: %v", duration.Seconds())))
	// } else {
	// 	if duration.Seconds() > 40 {
	// 		w.WriteHeader(500)
	// 		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	// 	} else {
	// 		w.WriteHeader(200)
	// 		w.Write([]byte("Tudo belezinha por aqui"))
	// 	}
	// }

}

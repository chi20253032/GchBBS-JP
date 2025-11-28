package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GchBBS-JP API Server is Running!")
}

func main() {
	// .envで設定したポートを取得 (デフォルト: 8080)
	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	fmt.Printf("Go application started on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
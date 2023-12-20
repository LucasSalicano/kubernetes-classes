package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kubernetes"))
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// HTTP GET isteğine yanıt vermek için bir işlev tanımlama
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Merhaba, Dünya!")
	})

	// Sunucuyu başlatın ve istekleri belirtilen adres ve port üzerinden dinleme
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}


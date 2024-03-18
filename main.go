package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// Mulai goroutine untuk memperbarui file JSON setiap 15 detik
	go func() {
		for {
			updateJSON()
			time.Sleep(15 * time.Second)
		}
	}()

	// Layani halaman web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Layani file JSON status
	http.HandleFunc("/status.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "status.json")
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func updateJSON() {
	// Buat struktur Status dengan nilai acak untuk air dan angin
	status := Status{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}

	// Marshal status ke format JSON dengan indentasi
	statusJSON, err := json.MarshalIndent(status, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Buat atau buka file status.json untuk ditulis
	file, err := os.Create("status.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Tulis status JSON ke file
	_, err = io.WriteString(file, string(statusJSON))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

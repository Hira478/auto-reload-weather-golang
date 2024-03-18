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
	// Start the goroutine to update the JSON file every 15 seconds
	go func() {
		for {
			updateJSON()
			time.Sleep(15 * time.Second)
		}
	}()

	// Serve the web page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/status.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "status.json")
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func updateJSON() {
	status := Status{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}

	statusJSON, err := json.MarshalIndent(status, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	file, err := os.Create("status.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.WriteString(file, string(statusJSON))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

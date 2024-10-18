package main

import (
	"gomailer/handlers"
	"log"
	"net/http"
)

func main() {
	// Enable CORS from every origin
	http.HandleFunc("/send-email", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}
		handlers.SendEmailHandler(w, r)
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Load users from JSON

}

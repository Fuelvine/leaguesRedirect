package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	// Load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Redirector setup
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rdPath := fmt.Sprintf("https://fuelvine.net/leagues%s", r.URL.Path)
		http.Redirect(w, r, rdPath, http.StatusPermanentRedirect)
	})

	// Activate
	fmt.Println("Redirector activated")
	err := http.ListenAndServe(os.Getenv("PORT"), mux)
	if err != nil {
		panic(err)
	}
}

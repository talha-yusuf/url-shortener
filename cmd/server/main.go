package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/talha-yusuf/url-shortener/internal/handlers"
	"github.com/talha-yusuf/url-shortener/internal/storage"
)

func main() {
	// Initialize storage
	urlStorage := storage.NewMemoryStorage()

	// Initialize handlers
	redirectHandler := handlers.NewRedirectHandler(urlStorage)
	homeHandler := handlers.NewHomeHandler(redirectHandler)
	createHandler := handlers.NewCreateHandler(urlStorage)
	analyticsHandler := handlers.NewAnalyticsHandler(urlStorage)

	// Set up routes
	http.Handle("/", homeHandler)
	http.Handle("/create", createHandler)
	http.Handle("/analytics", analyticsHandler)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Visit http://localhost:8080 in your browser to use the web interface!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Apply the CORS middleware to every request
	router.Use(corsMiddleware)

	router.HandleFunc("/generate_search_space", handleGenerateSearchSpace).Methods("GET")

	// router.HandleFunc("/search_spaces", indexSearchSpaces).Methods("GET")
	// router.HandleFunc("/search_spaces", createSearchSpace).Methods("POST")
	// router.HandleFunc("/search_spaces/{id}", showSearchSpace).Methods("GET")
	// router.HandleFunc("/search_spaces/{id}", deleteSearchSpace).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow specific origin or use "*" to allow all
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Check if the request is for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func sendErrorResponse(writer http.ResponseWriter, message string, code int) {
	response, err := errorResponse(message)
	if err != nil {
		http.Error(writer, "An error occurred", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(code)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func errorResponse(message string) ([]byte, error) {
	errorResponse := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: false,
		Message: message,
	}

	return json.Marshal(errorResponse)
}

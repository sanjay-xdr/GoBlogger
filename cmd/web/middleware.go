package main

import (
	"log"
	"net/http"
	"os"
)

func loggerfunc(next http.Handler) http.Handler {
	file, err := os.OpenFile("logger.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "INFO: ", log.LstdFlags)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr
		path := r.URL.Path
		method := r.Method

		// Log the IP address, method, path, and timestamp
		logger.Printf("IP: %s | Method: %s | Path: %s ", ip, method, path)
		// logger.Print("=====================================================")

		// Call the next handler
		next.ServeHTTP(w, r)
	})

}

func SessionLoad(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}

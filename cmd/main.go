package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "go-hello"
	}
	envName := os.Getenv("ENV_NAME")
	if envName == "" {
		envName = "unknown"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message":   fmt.Sprintf("Hello from %s!", appName),
			"env":       envName,
			"framework": "Go (stdlib)",
		})
	})

	fmt.Printf("%s listening on port %s\n", appName, port)
	http.ListenAndServe(":"+port, nil)
}

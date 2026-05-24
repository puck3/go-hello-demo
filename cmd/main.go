package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

const appVersion = 0

type settings struct {
	Port    string
	AppName string
	EnvName string
}

func main() {
	_ = godotenv.Load()

	cfg := settings{
		Port:    env("PORT", "8080"),
		AppName: env("APP_NAME", "go-hello"),
		EnvName: env("ENV_NAME", "unknown"),
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		noStore(w)
		writeJSON(w, map[string]string{"status": "ok"})
	})

	router.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		noStore(w)
		writeJSON(w, map[string]int{"version": appVersion})
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		noStore(w)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(page(cfg)))
	})

	fmt.Printf("%s listening on port %s\n", cfg.AppName, cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		panic(err)
	}
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func noStore(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Surrogate-Control", "no-store")
	w.Header().Set("X-App-Version", fmt.Sprintf("%d", appVersion))
}

func writeJSON(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func page(cfg settings) string {
	appName := html.EscapeString(cfg.AppName)
	envName := html.EscapeString(cfg.EnvName)
	renderedAt := html.EscapeString(time.Now().UTC().Format("2006-01-02 15:04:05 UTC"))

	return fmt.Sprintf(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>%s</title>
    <style>
      :root {
        color-scheme: light;
        font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
        background: #f4f7fb;
        color: #162033;
      }
      body {
        margin: 0;
        min-height: 100vh;
        display: grid;
        place-items: center;
      }
      main {
        width: min(760px, calc(100vw - 40px));
        padding: 48px;
        border: 1px solid #d9e2ef;
        border-radius: 8px;
        background: #ffffff;
        box-shadow: 0 18px 45px rgba(22, 32, 51, 0.08);
      }
      h1 {
        margin: 0 0 12px;
        font-size: 44px;
        line-height: 1.08;
        letter-spacing: 0;
      }
      p {
        margin: 0;
        color: #506174;
        font-size: 18px;
        line-height: 1.6;
      }
      dl {
        display: grid;
        grid-template-columns: max-content 1fr;
        gap: 10px 18px;
        margin: 32px 0 0;
        padding-top: 24px;
        border-top: 1px solid #edf1f7;
      }
      dt {
        color: #6a7889;
      }
      dd {
        margin: 0;
        font-weight: 700;
      }
    </style>
  </head>
  <body>
    <main>
      <h1>Hello from %s</h1>
      <p>This page is served by a Go application deployed through the automated deployment system.</p>
      <dl>
        <dt>Environment</dt>
        <dd>%s</dd>
        <dt>Framework</dt>
        <dd>Go + chi</dd>
        <dt>Version</dt>
        <dd>%d</dd>
        <dt>Rendered at</dt>
        <dd>%s</dd>
      </dl>
    </main>
  </body>
</html>`, appName, appName, envName, appVersion, renderedAt)
}

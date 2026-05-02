package main

import (
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
		body := fmt.Sprintf(`<!doctype html>
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
      <p>This page is served by a minimal Go application deployed through the automated deployment system.</p>
      <dl>
        <dt>Environment</dt>
        <dd>%s</dd>
        <dt>Framework</dt>
        <dd>Go (stdlib)</dd>
      </dl>
    </main>
  </body>
</html>`, appName, appName, envName)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(body))
	})

	fmt.Printf("%s listening on port %s\n", appName, port)
	http.ListenAndServe(":"+port, nil)
}

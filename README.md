# Go Hello Demo

Минимальное Go HTTP-приложение для проверки Go deployment template.

## Запуск приложения

- Port: `8080`
- HTML page: `/`
- Проверка состояния: `/health`
- Зависимости: нет

## Локальный запуск

```bash
go run ./cmd
```

Необязательные переменные:

```bash
APP_NAME="Go Hello App" ENV_NAME=production PORT=8080 go run ./cmd
```

## Настройки шаблона

Рекомендуемые значения deployment template:

- Base image: `golang:1.23-alpine`
- Root directory: `.`
- Output directory: `.`
- Install command: пусто
- Build command: `go build -o app ./cmd`
- Run command: `./app`
- Application port: `8080`

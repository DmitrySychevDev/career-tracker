# Career Tracker

Backend-сервис для трекинга вакансий, компаний и откликов при поиске работы.

Цель проекта: сделать production-like REST API на Go без лишней сложности, но с практиками, которые часто встречаются в backend-разработке: PostgreSQL, Docker Compose, миграции, JWT-аутентификация, middleware, валидация, тесты и CI.

## Стек

- Go
- PostgreSQL
- Docker Compose
- GORM
- golang-migrate
- gofumpt
- golangci-lint

## Локальный запуск

Скопируй пример env-файла:

```bash
cp .env.example .env
```

На Windows PowerShell:

```powershell
Copy-Item .env.example .env
```

Скачай зависимости приложения:

```bash
go mod download
```

Подними PostgreSQL:

```bash
docker compose up -d
```

Установи CLI для миграций, если он еще не установлен:

```bash
go install -tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Накати миграции:

```bash
migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/career_tracker?sslmode=disable" up
```

Запусти API:

```bash
go run ./cmd/api
```

Проверка:

```bash
curl http://localhost:8080/health
```

## Code Quality

Установка форматера:

```bash
go install mvdan.cc/gofumpt@latest
```

Установка линтера:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Форматирование:

```bash
gofumpt -w .
```

Линтинг:

```bash
golangci-lint run ./...
```

Проверка и автоисправление порядка полей в структурах:

```bash
golangci-lint run --enable fieldalignment --fix ./...
```

## Миграции

Создать новую миграцию:

```bash
migrate create -ext sql -dir migrations migration_name
```

Применить все новые миграции:

```bash
migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/career_tracker?sslmode=disable" up
```

Откатить последнюю миграцию:

```bash
migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/career_tracker?sslmode=disable" down 1
```

Для dev-сброса базы вместе с Docker volume:

```bash
docker compose down -v
docker compose up -d
migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/career_tracker?sslmode=disable" up
```
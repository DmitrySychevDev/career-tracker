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

Подними PostgreSQL:

```bash
docker compose up -d
```

Установи migrate CLI с PostgreSQL-драйвером:

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

## MVP Roadmap

### Epic 1. Аутентификация

- Регистрация пользователя
- Хеширование пароля через bcrypt
- Login
- JWT access token
- Refresh token
- Logout

### Epic 2. Пользователь

- Просмотр профиля
- Изменение имени
- Изменение почты
- Смена пароля

### Epic 3. Компании

- CRUD компаний
- Название, сайт, страна, город, заметка, ссылка на карьерную страницу

### Epic 4. Вакансии

- CRUD вакансий
- Компания, зарплата, валюта, тип работы, ссылка, описание, стек, дата публикации

### Epic 5. Отклики

- Создание отклика
- Статусы: Wishlist, Applied, HR, Test Task, Technical, Final Interview, Offer, Accepted, Rejected, Archived
- Дата отклика, ожидаемая зарплата, заметки, контакты рекрутера

### Epic 6. Таймлайн

- Сохранять историю изменения статусов отклика
- Пример: Applied -> HR Interview -> Technical -> Offer

### Epic 7. Заметки

- Комментарии к откликам
- Например: "Техничка через неделю", "Нужно повторить Docker"

### Epic 8. Dashboard

- Всего вакансий
- Всего откликов
- Офферы
- Отказы
- Конверсия
- Среднее время до ответа
- Средняя зарплата

### Epic 9. Поиск

- По компании
- По названию вакансии
- По стеку

### Epic 10. Фильтрация

- По статусу
- По зарплате
- По компании
- По типу работы
- По стране
- По диапазону дат

### Epic 11. Сортировка

- По зарплате
- По дате
- По названию
- По компании

### Epic 12. Пагинация

- page
- limit
- offset

### Epic 13. Избранное

- Отметка интересной вакансии

### Epic 14. Теги

- Например: Go, Backend, Remote, Dream Job, Relocation

### Epic 15. Документы

- Резюме
- Тестовое задание
- Сопроводительное письмо
- На первом этапе хранить путь к файлу

### Epic 16. Уведомления

- Фоновый воркер на goroutines/ticker
- Проверка откликов в статусе Applied старше 14 дней
- Создание уведомления "Пора написать рекрутеру"

### Epic 17. Админка

- Количество пользователей
- Количество вакансий
- Количество откликов

### Epic 18. REST API

- `/auth`
- `/users`
- `/companies`
- `/vacancies`
- `/applications`
- `/timeline`
- `/dashboard`
- `/tags`

## Нефункциональные требования

- Конфигурация через `.env`
- Docker Compose
- PostgreSQL
- Redis для кэша статистики/dashboard
- Swagger
- Миграции БД
- Structured logging через `slog`
- Graceful shutdown
- Middleware: logger, recovery, auth, request id
- Валидация входящих данных
- Централизованная обработка ошибок
- Unit-тесты
- Integration-тесты
- GitHub Actions: lint, test, build
- golangci-lint

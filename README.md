# KhanProj

Тестовый проект на Go (Golang) для позиции Junior Golang Developer в компании Effective Mobile.

Проект реализует простой CRUD-сервис для работы с данными людей, с обогащением информации через внешние публичные API.

---

## 🚀 Как запустить проект

1. Склонировать репозиторий:
    ```bash
    git clone https://github.com/ATursunbekov/KhanProj.git
    cd KhanProj
    ```

2. Создать и настроить файл `.env`:
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=your_password
    DB_NAME=person_db
    PORT=8080
    ```

3. Запустить базу данных через Docker:
    ```bash
    docker-compose up -d
    ```

4. Выполнить миграцию вручную или убедиться, что таблица `persons` создана.

5. Установить зависимости и сгенерировать Swagger-документацию:
    ```bash
    go mod tidy
    swag init -g cmd/main.go -o docs
    ```

6. Запустить проект:
    ```bash
    go run cmd/main.go
    ```

---

## 📚 Описание проекта

- **Язык**: Go (Golang)
- **Фреймворк**: Gin
- **БД**: PostgreSQL
- **Работа с БД**: sqlx
- **Документация API**: Swagger (через swaggo)
- **Docker**: используется для поднятия PostgreSQL

---

## 🛠 Реализованные запросы

| Метод | URL | Описание |
|------|-----|----------|
| `POST` | `/person/create` | Создать нового человека. Обогащение через API (возраст, пол, национальность). |
| `DELETE` | `/person/delete/{id}` | Удалить человека по ID. |
| `PUT` | `/person/update` | Полностью обновить данные человека. |
| `GET` | `/person/getPerson/{id}` | Получить информацию о человеке по ID. |
| `GET` | `/person/getAll` | Получить список всех людей с фильтрами и пагинацией. |

---

## 🔎 Swagger документация

После запуска проекта Swagger доступен по адресу:
http://localhost:8080/swagger/index.html#/Person/post_person_create



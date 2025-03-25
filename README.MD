# Тестовое задание для компании Gravitum
## Users Service
Users Service — это REST API для управления пользователями (CRUD), написанный на Go. Приложение использует PostgreSQL для хранения данных и Docker/Docker Compose для контейнеризации и упрощённого запуска.
## Содержание
1. Содержание
2. Структура репозитория
3. Технологии
4. Требования
5. Установка и запуск
6. Переменные окружения
7. Использования API
8. Тестирование
9. Вопросы и улучшения

## Особенности проекта
1. CRUD-операции: создание, получение, обновление и удаление пользователей.
2. Слой сервисов для бизнес-логики, слой репозиториев для доступа к БД, handlers для обработки HTTP-запросов.
3. PostgreSQL в качестве СУБД.
4. Контейнеризация с помощью Docker и Docker Compose.
5. Тесты (модульные) для обеспечения качества кода.

## Структура репозитория
```
├── cmd/
│   └── main.go               # Точка входа в приложение
├── internal/
│   ├── models/
│   │   └── user.go           # Определение структуры User
│   ├── repository/
│   │   ├── db.go             # Подключение к PostgreSQL
│   │   └── user_repository.go
│   ├── service/
│   │   └── user_service.go
│   │   └── user_service_test.go
│   └── handler/
│       └── user_handler.go
│       └── user_handler_test.go
├── Dockerfile                # Файл для сборки Docker-образа
├── docker-compose.yml        # Композиция для запуска app и db
├── .env.example              # Пример файла с переменными окружения
├── go.mod
├── go.sum
└── README.md                 # Настоящий файл с документацией
```
## Технология
1. Go 1.23
2. PostgreSQL
3. Docker/Docker Compose
4. GORM 

## Требования
1. Установленный Docker desktop и Docker Compose

## Установка и запуск
Запуск в Docker
1. Склонировать репозиторий:
```
git clone https://github.com/UndeadDxD/users.git
cd users
```
2. Скопировать .env.example в .env и заполнить реальными данными: <br>
    Заменить DB_PASSWORD на любой с @ на конце. <br> В качестве примера можно использовать : 123Ff321@
3. Запустить Docker Compose:
```
docker-compose up --build
```
Это соберёт образ приложения и поднимет контейнеры:
  * db (PostgreSQL) на порту 5432
  * app (Go-приложение) на порту 8080
4. Проверить работу: <br>
   Приложение будет доступно по адресу http://localhost:8080.
5. Переменные окружения <br>
   В проекте используются переменные окружения для конфигурации подключения к БД, порта приложения и т.д. Например (в .env.example): 
    * DB_HOST=db хост или имя сервиса (если в Docker Compose).
    * DB_PORT=5432 порт PostgreSQL.
    * DB_USER=postgres учётные данные PostgreSQL.
    * DB_PASSWORD=password пароль для бд
    * DB_NAME=users_db название базы данных.
6. Использование API <br>
   Приложение предоставляет REST-эндпоинты для управления пользователями. Примерные роуты (по умолчанию http://localhost:8080):
   * POST /users
   Создать пользователя. <br>
   Body (JSON):
   ```
   {
   "name": "Alice",
   "email": "alice@example.com"
   }
   ```
   * GET /users/{id}
   Получить пользователя по ID.<br>
   * PUT /users/{id}
   Обновить данные пользователя по ID.<br>
   Body (JSON):<br>
   ```
   {
   "name": "Alice Updated",
   "email": "alice.updated@example.com"
   }
   ```
   * DELETE /users/{id}<br>
   Удалить пользователя по ID.
   * При успешном выполнении возвращаются соответствующие коды HTTP (200 OK, 201 Created, 204 No Content, 404 Not Found и т.д.).
7. Тестирование<br>
   Для запуска тестов (модульных) используйте команду:<br>
   ```
   cd internal/handler
   go test
   
   cd internal/service
   go test
   ```


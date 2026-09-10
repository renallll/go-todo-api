# Go Todo API

RESTful API untuk aplikasi Todo List yang dibangun menggunakan Go dengan Gin Framework, PostgreSQL, GORM, dan JWT Authentication.

Project ini dibuat sebagai backend project untuk mempraktikkan pengembangan REST API, database integration, authentication, Docker containerization, API documentation, testing, dan CI menggunakan GitHub Actions.

## Features

* User Registration
* User Login
* JWT Authentication
* Protected API endpoints
* Task CRUD
* Task ownership
* PostgreSQL database
* GORM ORM
* Repository and Service Layer
* Swagger API Documentation
* Docker & Docker Compose
* PostgreSQL Healthcheck
* Unit Testing
* GitHub Actions CI

## Tech Stack

| Technology       | Usage                        |
| ---------------- | ---------------------------- |
| Go               | Backend programming language |
| Gin              | HTTP web framework           |
| PostgreSQL       | Relational database          |
| GORM             | ORM                          |
| JWT              | Authentication               |
| Swagger / Swaggo | API documentation            |
| Docker           | Containerization             |
| Docker Compose   | Multi-container environment  |
| GitHub Actions   | Continuous Integration       |

## Project Structure

```text
todo-api/
│
├── .github/
│   └── workflows/
│       └── go.yml
│
├── cmd/
│   └── main.go
│
├── database/
│   └── database.go
│
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── handlers/
│   ├── auth_handler.go
│   └── task_handler.go
│
├── middleware/
│   └── auth.go
│
├── models/
│   ├── task.go
│   └── user.go
│
├── repositories/
│   └── task_repository.go
│
├── routes/
│   └── routes.go
│
├── services/
│   ├── task_service_test.go
│   └── task_services.go
│
├── utils/
│   └── jwt.go
│
├── .dockerignore
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── go.sum
```

> `.env` digunakan untuk konfigurasi lokal dan sengaja tidak disimpan di repository.

## Architecture

Project menggunakan pendekatan layered architecture untuk memisahkan tanggung jawab setiap bagian aplikasi.

```text
Client
   │
   ▼
Routes
   │
   ▼
Middleware
   │
   ▼
Handlers
   │
   ▼
Services
   │
   ▼
Repositories
   │
   ▼
PostgreSQL
```

### Layer Responsibilities

**Handlers**

Menangani HTTP request dan response dari client.

**Services**

Berisi business logic aplikasi.

**Repositories**

Menangani operasi database menggunakan GORM.

**Models**

Mendefinisikan struktur data seperti User dan Task.

**Middleware**

Menangani JWT authentication untuk protected endpoints.

## Authentication Flow

```text
Register
   │
   ▼
Create User
   │
   ▼
Login
   │
   ▼
Generate JWT
   │
   ▼
Client receives Token
   │
   ▼
Authorization: Bearer <token>
   │
   ▼
Auth Middleware
   │
   ▼
Protected API
```

## API Endpoints

### Authentication

| Method | Endpoint    | Authentication | Description       |
| ------ | ----------- | -------------- | ----------------- |
| POST   | `/register` | No             | Register new user |
| POST   | `/login`    | No             | Login user        |

### Tasks

| Method | Endpoint     | Authentication | Description      |
| ------ | ------------ | -------------- | ---------------- |
| GET    | `/tasks`     | Yes            | Get user's tasks |
| GET    | `/tasks/:id` | Yes            | Get task by ID   |
| POST   | `/tasks`     | Yes            | Create new task  |
| PUT    | `/tasks/:id` | Yes            | Update task      |
| DELETE | `/tasks/:id` | Yes            | Delete task      |

## API Request & Response

### 1. Register

**POST `/register`**

Request:

```json
{
  "username": "renald",
  "password": "password123"
}
```

Example response:

```json
{
  "message": "User registered successfully"
}
```

### 2. Login

**POST `/login`**

Request:

```json
{
  "username": "renald",
  "password": "password123"
}
```

Example response:

```json
{
  "token": "your-jwt-token"
}
```

Token tersebut digunakan untuk mengakses endpoint yang membutuhkan authentication.

### Authorization Header

```http
Authorization: Bearer <your-jwt-token>
```

### 3. Get Tasks

**GET `/tasks`**

Request:

```http
GET /tasks
Authorization: Bearer <your-jwt-token>
```

Example response:

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "description": "Learn Go REST API",
    "completed": false
  }
]
```

### 4. Get Task by ID

**GET `/tasks/:id`**

Example:

```http
GET /tasks/1
Authorization: Bearer <your-jwt-token>
```

### 5. Create Task

**POST `/tasks`**

Request:

```json
{
  "title": "Learn PostgreSQL",
  "description": "Practice PostgreSQL with GORM"
}
```

Example response:

```json
{
  "id": 1,
  "title": "Learn PostgreSQL",
  "description": "Practice PostgreSQL with GORM",
  "completed": false
}
```

### 6. Update Task

**PUT `/tasks/:id`**

Request:

```json
{
  "title": "Learn Go and PostgreSQL",
  "description": "Build a REST API using Go",
  "completed": true
}
```

### 7. Delete Task

**DELETE `/tasks/:id`**

Example:

```http
DELETE /tasks/1
Authorization: Bearer <your-jwt-token>
```

## Environment Variables

Repository ini hanya menyediakan `.env.example`.

Copy `.env.example` menjadi `.env` untuk konfigurasi lokal.

### Windows PowerShell

```powershell
Copy-Item .env.example .env
```

Example `.env.example`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=todo_db

JWT_SECRET=your_secret_key

PORT=8080
```

> Jangan commit file `.env` ke Git. File `.env` digunakan hanya untuk environment lokal.

## Running Locally

### 1. Clone Repository

```bash
git clone git@github.com:renallll/go-todo-api.git
cd go-todo-api
```

### 2. Create Environment File

Windows PowerShell:

```powershell
Copy-Item .env.example .env
```

Kemudian sesuaikan konfigurasi database dan JWT secret di `.env`.

### 3. Install Dependencies

```bash
go mod download
```

### 4. Run PostgreSQL

Pastikan PostgreSQL sudah berjalan di komputer.

Kemudian jalankan:

```bash
go run ./cmd
```

API akan berjalan pada:

```text
http://localhost:8080
```

## Running with Docker

Project sudah dilengkapi dengan Dockerfile dan Docker Compose.

Build dan jalankan seluruh service:

```bash
docker compose up --build
```

Untuk menjalankan di background:

```bash
docker compose up -d --build
```

Stop container:

```bash
docker compose down
```

Melihat container:

```bash
docker compose ps
```

Melihat logs:

```bash
docker compose logs -f
```

### Docker Architecture

```text
                 ┌─────────────────┐
                 │     Client      │
                 └────────┬────────┘
                          │
                          ▼
                 ┌─────────────────┐
                 │    Go API       │
                 │   Gin + GORM    │
                 └────────┬────────┘
                          │
                          ▼
                 ┌─────────────────┐
                 │   PostgreSQL    │
                 │      16         │
                 └─────────────────┘
```

PostgreSQL menggunakan Docker healthcheck agar API menunggu database siap sebelum melakukan koneksi.

## Swagger Documentation

API documentation tersedia menggunakan Swagger / Swaggo.

Setelah aplikasi berjalan, buka:

```text
http://localhost:8080/swagger/index.html
```

Swagger dapat digunakan untuk:

* Melihat seluruh endpoint
* Melihat request dan response
* Mencoba API secara langsung
* Mengirim JWT Authorization
* Testing endpoint tanpa Postman

Jika Swagger perlu di-generate ulang:

```bash
swag init -g cmd/main.go
```

## Testing

Menjalankan seluruh unit test:

```bash
go test ./...
```

Menjalankan test dengan verbose output:

```bash
go test ./... -v
```

Project memiliki unit test untuk service layer.

## GitHub Actions

Project menggunakan GitHub Actions untuk menjalankan proses Continuous Integration.

Workflow akan melakukan pengecekan project ketika terdapat push atau pull request.

Contoh proses CI:

```text
Push / Pull Request
        │
        ▼
GitHub Actions
        │
        ├── Setup Go
        │
        ├── Download Dependencies
        │
        ├── Build
        │
        └── Run Tests
```

## Git Branching

Pengembangan project dilakukan menggunakan beberapa feature branch:

```text
main
│
├── feature/auth-jwt
├── feature/clean-architecture
├── feature/dockerize-api
├── feature/github-actions
├── feature/postgresql
└── feature/task-ownership
```

Feature dikembangkan pada branch masing-masing kemudian diintegrasikan ke `main` melalui Pull Request.

## Security

Beberapa praktik keamanan yang digunakan:

* Password tidak disimpan secara plaintext
* JWT digunakan untuk authentication
* Protected endpoints menggunakan authentication middleware
* `.env` tidak disimpan di repository
* `.env.example` digunakan sebagai template konfigurasi
* Task hanya dapat diakses oleh user yang memiliki task tersebut

> Jika secret/password pernah tersimpan di Git history, secret tersebut sebaiknya tetap dianggap compromised dan diganti.

## Future Improvements

Beberapa pengembangan yang dapat dilakukan selanjutnya:

* Refactor authentication ke Clean Architecture
* User Repository dan User Service
* DTO untuk request dan response
* Centralized error handling
* Authentication unit tests
* Integration tests
* Refresh token
* Password hashing menggunakan bcrypt
* Pagination untuk task list
* API rate limiting
* Logging
* Deployment ke cloud

## Learning Goals

Project ini dibuat untuk memperdalam kemampuan dalam:

* Go programming
* REST API development
* Gin Framework
* PostgreSQL
* GORM
* JWT Authentication
* Clean Architecture
* Docker
* Swagger
* Unit Testing
* Git & GitHub
* GitHub Actions
* API design and development

## Author

**Rhenald Adrian Nainggolan**

Informatics Engineering Graduate
Backend Development | Software Engineering | Data Analytics

GitHub: [renallll](https://github.com/renallll)

## License

This project is available for educational and portfolio purposes.

# auth_exam

## 📌 Description

`auth_exam` is a simple authentication and authorization service written in Golang.  
It is built for **interview preparation** and **educational purposes**.

This service includes:
- User signup and login
- JWT token generation
- Role-based access control (e.g., superadmin)
- Protected routes using middleware
- Swagger API documentation
- PostgreSQL integration

---

## ⚙️ Technologies Used

- **Go (Golang)** – Gin framework
- **PostgreSQL** – for persistent storage
- **JWT** – for secure token authentication
- **Swaggo** – for Swagger/OpenAPI docs
- **godotenv** – to load environment variables

---

## 📁 Project Structure


---

## 🚀 Getting Started

### ✅ Prerequisites

- Go 1.20+
- PostgreSQL installed and running
- Git

### 📥 Installation

```bash
git clone https://github.com/web-gopro/auth_exam.git
cd auth_exam
go mod tidy
go run cmd/main.go

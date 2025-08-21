# TodoApp Golang

## 🚀 Features

- RESTful endpoints
- JSON request/response
- Environment variable configuration
- Docker support

---

## 📦 Requirements

Before running, make sure you have installed:

- [Go](https://go.dev/dl/) (>=1.20)
- [Docker](https://www.docker.com/) (optional, for containerized run)
- [Git](https://git-scm.com/)

---

## ⚙️ Setup & Run Locally

1. **Clone the repository**

   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   ```

2. **Create .env file base on .env.example**

3. **Install dependencies**

   ```bash
    go mod tidy
   ```

4. **Run the server**

   ```bash
    go run main.go
   ```

5. **Run the server**

   ```bash
    curl http://localhost:8080/health
   ```

---

## 🐳 Run with Docker

1. **Build Docker image**

   ```bash
   docker build -t go-api .
   ```

2. **Run container**

   ```bash
   docker run -p 8080:8080 --env-file .env go-api
   ```

---

## 📂 Project Structure

- `main.go`: Defines the server entry point
- `Dockerfile`: Used to build the Docker image
- `storage/`: Contains in-memory data storage
- `models/`: Defines the data models
- `handlers/`: Handles API routes

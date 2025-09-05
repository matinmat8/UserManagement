# Authentication Service (Go + Redis + Swagger)

This project is a simple authentication service built with **Golang**, **Redis**, and documented using **Swagger**.  
It supports OTP login, JWT token management, and user listing with pagination and search.  

---

## 📦 Prerequisites
- [Docker](https://docs.docker.com/get-docker/) installed  
- [Docker Compose](https://docs.docker.com/compose/) installed  

---

## 🚀 How to Run (Dockerized)

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/auth-service.git
   cd auth-service


2 - Generate Swagger docs (only needed during development):

    go install github.com/swaggo/swag/cmd/swag@latest
    swag init -g main.go -o ./docs


The generated docs/ folder should already be in the repo.
If you’re running just with Docker, you don’t need to run this step.


3- Start the application with Docker Compose:
```
  docker-compose up --build
```

4- The services will be available at:
Go API → http://localhost:8080
Redis → localhost:6379


📖 API Documentation (Swagger)
Swagger UI is served directly from the app.
Once the container is running, open:
👉 http://localhost:8080/swagger/index.html


🛠 Environment Variables

The app uses the following environment variables (configured in docker-compose.yml):

| Variable     | Default | Description        |
| ------------ | ------- | ------------------ |
| `REDIS_HOST` | redis   | Redis service name |
| `REDIS_PORT` | 6379    | Redis port         |


🧹 Useful Commands

- Rebuild containers:
``` docker-compose up --build ```

- Stop containers:
```docker-compose down```

- View logs:
```docker-compose logs -f <container_id>```

- Enter Redis CLI:
```docker exec -it redis redis-cli```



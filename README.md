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



---

## ⚡ Error Handling

This project uses a **centralized error handling middleware** with Gin.  
Instead of returning raw errors, we **panic with a custom `PanicMessage` struct**. The middleware recovers from panics and translates them into meaningful JSON responses.

### 🔹 How it Works
1. Each part of the application (repositories, services, controllers) can `panic(utils.PanicMessage{MessageKey: <key>})` when something goes wrong.
2. The middleware (`middleware/ErrorHandling.go`) intercepts the panic using `recover()`.
3. It looks up the error message in a **message template map** (`pkg/templates`).
4. A structured JSON response is returned to the client with the correct HTTP status code and user-friendly message.
5. The error is also logged with depth information using `utils/logger`.

### 🔹 Example Response
If OTP was already sent, the response might look like:

```json
{
  "fa_message": "کد یکبار مصرف قبلاً ارسال شده است",
  "en_message": "OTP already sent"
}
```
🔹 Why This is Useful
- Keeps controller code clean (no repetitive if err != nil handling everywhere).
- Ensures consistent error messages across the whole project.
- Makes it easier to log and debug errors.
- Provides user-friendly API responses instead of raw stack traces.

🗂 File Logger (Zerolog + Lumberjack)

Logging is handled by a combination of:

- Zerolog → High-performance structured logging.
- Lumberjack  → Log file rotation & compression.

Features
- Logs are written to logs/auth.log.
- Each file rotates at 200 MB.
- Old logs are compressed automatically.
- Errors are written with the file name, line number, and details for easier debugging.

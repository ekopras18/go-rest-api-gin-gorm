# To Do List

- [x] CRUD
- [x] Auth JWT
- [x] Middleware
- [ ] Rate Limiter
- [x] CORS
- [ ] Docker

# Structure of the Project

```bash
go-rest-api-gin-gorm/
│── cmd/
│   ├── api/
│   │   ├── ...
│── config/
│   ├── ...
│── docs/
│   ├── ...
│── internal/
│   ├── app/
│   │   ├── ...
│   ├── database/
│   │   ├── ...
│   ├── repository/
│   │   ├── ...
│   ├── service/
│   │   ├── ...
│   ├── controller/
│   │   ├── ...
│   ├── models/
│   │   ├── ...
│   ├── dto/
│   │   ├── ...
│── pkg/
│   ├── utils/
│   │   ├── ...
│── .air.toml
│── .env.example
│── .gitignore
│── go.mod
│── go.sum
│── README.md
│── server.log
```

## Installation

- Clone this repository
```bash
git clone https://github.com/ekopras18/go-rest-api-gin-gorm.git
```
- Go into the repository
```bash
cd go-rest-api-gin-gorm
```

- Install dependencies
```bash
go mod tidy
```

- Run the application
    - run with the default environment
        ```bash
        go run cmd/api/main.go 
        ````
  
    - or, Run with hot reload (install air first)
        ```bash
        go install github.com/cosmtrek/air@latest
        ```
    - then
      ```bash
      air .
      ```
      

## Swagger UI

- Open your browser and navigate to `http://localhost:8080/docs/index.html` to view the API documentation.
- You can also use the Swagger UI to test the API endpoints directly from your browser.
- Make sure to have the server running before accessing the Swagger UI.
- To generate the Swagger documentation, you can use the following command:
    ```bash
    swag init -g cmd/api/main.go
    ```

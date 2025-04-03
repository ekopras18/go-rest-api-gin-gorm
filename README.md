```bash
your-app/
│── cmd/
│   ├── api/
│   │   ├── main.go
│── config/
│   ├── config.go
│   ├── config.yaml
│── internal/
│   ├── app/
│   │   ├── server.go
│   │   ├── routes.go
│   │   ├── middleware.go
│   ├── database/
│   │   ├── db.go
│   │   ├── migration.go
│   ├── repository/
│   │   ├── user_repository.go
│   ├── service/
│   │   ├── user_service.go
│   ├── controller/
│   │   ├── user_handler.go
│   ├── entities/
│   │   ├── user_entities.go
│── pkg/
│   ├── logger/
│   │   ├── logger.go
│   ├── utils/
│   │   ├── hash.go
│── tests/
│   ├── user_test.go
│── .env
│── go.mod
│── go.sum
│── README.md
```
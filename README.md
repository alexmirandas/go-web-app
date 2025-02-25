# Go Web App 222

This is a simple Go web application with unit tests.

## Features
- Simple HTTP server
- `/` endpoint returning "Hello, World!"
- `/health` endpoint returning "OK" for health check
- Unit tests for API handlers

## Project Structure
```bash
go-web-app/
│── main.go
│── main_test.go
│── go.mod
│── go.sum
│── README.md
└── handlers/
    ├── handlers.go
    ├── handlers_test.go

````


## Prerequisites
- Go installed (1.18+)

## Installation
Clone the repository and initialize dependencies:
```sh
git clone https://github.com/your-repo/go-web-app.git
cd go-web-app
go mod tidy
```

## Run
```bash
go run main.go
```

Un comentario mas
## Running Tests
```bash
go test -v ./handlers
```

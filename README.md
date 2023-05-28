# Bookify - Book store API service

### Interactive API documentation (provided by Swagger UI, Redoc, Rapidoc)
```
- http://localhost:8080/docs/swagger
- http://localhost:8080/docs/redoc
- http://localhost:8080/docs/rapidoc
```


## Run Seeder
```terminal
go mod vendor

go run main.go seed

or

go run main.go seed --truncate=true or -t=true
```

## Local
```terminal
go mod vendor

go run main.go serve
```

## Docker
```terminal
make development
```
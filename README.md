# go_backend

Getting started

```bash
> docker-compose up -d

Creating go-backend_postgres_1 ... done
Creating go-backend_pgadmin_1  ... done
Creating go-backend_mongo_1    ... done
Creating go-backend_app_1      ... done
Creating go-backend_mongo-express_1 ... done
```

Then attach Visual Studio Code to `go-backend_app_1` through docker extension.

To debug change the `HTTP_PORT` to another port inside the remote session.

To run all tests with coverage and check the coverage report:

```bash
> go test -coverprofile=coverage.out ./... 
> go tool cover -html=coverage.out -o coverage.html
```

To run only the unit tests (integration tests will be skipped):

```bash
> go test -short -coverprofile=coverage.out ./...
```
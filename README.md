# go-jwt

Basic JWT Authentication in golang (following [this tutorial](https://www.youtube.com/watch?v=ma7rUS_vW9M))

## Repo Strcuture

```
├── go.mod
├── go.sum
├── initializers
│   └── connectToDB.go
├── models
│   └── userModel.go
├── controllers
│   └── usersController.go
├── middleware
│   └── requireAuth.go
├── main.go <-- main file
├── migrate
│   └── migrate.go <-- standalone script for database migration
```

## Steps to run

(assuming database is already set up)

1. Download binary for your OS from [releases page](https://github.com/ShobanChiddarth/go-jwt/releases)
2. Give it executable permissions (if in Linux, `chmod +x`)
3. Set environment variables like [sample.env](./sample.env)
4. Run it
5. Base URL is `locahost:8002` (or whatever port you changed to in env variable `PORT`), send requests to it according to [API DOCS](./API-DOCS.md)

## Steps to build from source

1. Clone this repo
2. Set environment variables correctly for your database (see [sample.env](./sample.env))
3. Run [build-migrate.bash](./build-migrate.bash)
4. Run the database migration executable (of your platform) to seed the database
5. Run [build.bash](./build.bash)
6. Run the executable (of your platform)

## Steps to run without building

Same as above but instead of running the build scripts and running executables do `go run main.go` or `go run migrate/migrate.go`


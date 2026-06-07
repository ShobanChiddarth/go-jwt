docker run -d -p 127.0.0.1:5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=PASSWORD -e POSTGRES_DB=go-jwt --name postgres-go-jwt --name postgres-db postgres:16

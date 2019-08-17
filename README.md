# Go + React example

A simple Golang REST API + React UI.

## Installation

After cloning the repo, assuming you have `docker` installed on your machine, do the following steps in the project folder:

1. For local testing with MySQL run by Docker set env vars `MYSQL_PASS` for DB password and `CLEARDB_DATABASE_URL`as the MySQL URL. 
2. Run `docker-compose up` or `docker-compose up --build` if you want to rebuild everything and ignore the cache. 

## Tests

Steps:
1. Generate mocks with `go generate ./...`
2. Run all tests with: `go test -v ./...`
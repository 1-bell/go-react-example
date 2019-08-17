# Go + React example

A simple Golang REST API + React UI.

## Installation

### For the API:
After cloning the repo, assuming you have `docker` installed on your machine, do the following steps in the project folder:

1. For local testing with MySQL run by Docker set env vars `MYSQL_PASS` for DB password and `MYSQL_URL` as the MySQL URL. 
2. Run `docker-compose up` or `docker-compose up --build` if you want to rebuild everything and ignore the cache.

It will start serving requests on port `8080`.

### For the UI:
1. It's going to use node to serve the static files, so make sure you have `node` installed.
2. Go to the root folder with `cd ui` and then run `npm start`.

It will start serving static files on port `3000`.

## Tests

Steps:
1. API: generate Go mocks with `go generate ./...` and then run all Go tests with: `go test -v ./...`
2. UI: go to the root folder with `cd ui` and then run `npm test`.

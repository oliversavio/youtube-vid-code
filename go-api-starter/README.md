# This is Part 4 of the series on building production read APIs
 - [Part 1](https://medium.com/code-uncomplicated/building-apis-in-go-beyond-hello-world-5d501d6403de)
 - [Part 2](https://medium.com/code-uncomplicated/you-should-unit-test-your-api-endpoints-its-easy-a685abc87982)
 - [Part 3](https://www.youtube.com/watch?v=pP2DKCKR4CQ)

## System Requirements
 - Go 1.16
 - Install `go install github.com/swaggo/swag/cmd/swag@latest`
 - Docker CE 

## Build Docker Image
```
make build
```

## Run Docker Image on default port 9000
```
make run
```

## Run Nginx on port 8080 with 2 API instances load-balanced
```
make start-server
```

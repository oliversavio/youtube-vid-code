# A series of posts on Building APIs 
 - [Part 1: Building APIs beyond Hello World](https://medium.com/code-uncomplicated/building-apis-in-go-beyond-hello-world-5d501d6403de)
 - [Part 2: Unit Test for API Endpoints](https://medium.com/code-uncomplicated/you-should-unit-test-your-api-endpoints-its-easy-a685abc87982)
 - [Part 3: Monitoring with Prometheus and Grafana](https://www.youtube.com/watch?v=pP2DKCKR4CQ)
 - [Part 4: Setting up Nginx as a reverse proxy, load-balancer and rate-limiter](https://medium.com/code-uncomplicated/why-you-should-consider-nginx-in-front-of-your-microservice-9deb6b5996cb)

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


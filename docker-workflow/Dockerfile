from golang:1.16 AS builder

LABEL MAINTAINER="Oliver Mascarenhas"

WORKDIR /go/src/
COPY . .
RUN GOOS=linux go build -a -o bin/app *.go

FROM debian:buster-slim
WORKDIR /api/
COPY --from=builder /go/src/bin/app .

EXPOSE 3000
CMD ["./app"]


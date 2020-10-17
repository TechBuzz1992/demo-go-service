FROM golang:latest

LABEL maintainer="harsh.dusane@gmail.com"

WORKDIR /demo-go-service

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]


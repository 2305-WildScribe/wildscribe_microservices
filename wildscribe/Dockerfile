FROM golang:1.21.3-alpine as builder

COPY go.mod go.sum ./

WORKDIR /wildscribe
COPY ./wildscribe .

RUN  go build -o main.exe $(go list ./cmd/...)

FROM alpine:latest
COPY /wildscribe/main.exe .
EXPOSE 8080
CMD ["./main.exe"]

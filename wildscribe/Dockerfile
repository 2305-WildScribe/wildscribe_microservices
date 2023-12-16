# Builder stage
FROM golang:1.21.3-alpine as builder

WORKDIR /wildscribe
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.exe $(go list ./wildscribe/cmd/...)

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /wildscribe/main.exe .

EXPOSE 8080
CMD ["./main.exe"]

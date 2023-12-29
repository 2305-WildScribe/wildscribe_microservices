# # Builder stage
FROM golang:1.21.3-alpine as builder

WORKDIR /user
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.exe $(go list ./user/cmd/...)

# Final stage
FROM alpine:latest


WORKDIR /user
COPY --from=builder /user/main.exe .

EXPOSE 8082
CMD ["./main.exe"]

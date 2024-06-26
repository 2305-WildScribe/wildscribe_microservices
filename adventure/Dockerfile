# # Builder stage
FROM golang:1.21.3-alpine as builder

WORKDIR /adventure
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.exe $(go list ./adventure/cmd/...)

# Final stage
FROM alpine:latest


WORKDIR /adventure
COPY --from=builder /adventure/main.exe .

EXPOSE 8082
CMD ["./main.exe"]
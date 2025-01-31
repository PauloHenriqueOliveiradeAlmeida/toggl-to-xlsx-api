FROM golang:1.23-alpine as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8000
CMD ["./main"]

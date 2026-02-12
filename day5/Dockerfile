FROM golang:1.25.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog-api cmd/api/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/blog-api .

EXPOSE 8080
ENTRYPOINT ["/app/blog-api"]

FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/delete-service ./cmd

FROM scratch
COPY --from=builder /bin/delete-service /delete-service
EXPOSE 8084
ENTRYPOINT ["/delete-service"]

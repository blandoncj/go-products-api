FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/update-service ./cmd

FROM scratch
COPY --from=builder /bin/update-service /update-service
EXPOSE 8083
ENTRYPOINT ["/update-service"]

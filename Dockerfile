FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/migration ./migration 
COPY --from=builder /app/assets ./assets 

EXPOSE 8080 

# Set environment variables
ENV DB_USER=postgres
ENV DB_PASSWORD=Honshou!@#123
ENV DB_NAME=kreditplus_db
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DEFAULT_SEARCH_PATH=public
ENV SERVER_PORT=8045
ENV MIGRATION_PATH=./migration

# Run the binary
CMD ["./main"]
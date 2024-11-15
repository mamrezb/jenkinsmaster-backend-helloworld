# Use an official Golang image as a builder
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

# Download dependencies and build the app
RUN go mod download && go build -o backend-app .

# Run stage: using a minimal image for the final container
FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/backend-app .
EXPOSE 8080

CMD ["./backend-app"]

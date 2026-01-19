# 1. Build stage
FROM golang:1.21-alpine AS builder

# 2. Set working directory
WORKDIR /app

# 3. Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# 4. Copy the rest of the source code
COPY . .

# 5. Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notely main.go

# 6. Final lightweight image
FROM alpine:latest

# 7. Install CA certificates
RUN apk --no-cache add ca-certificates

# 8. Set working directory
WORKDIR /root/

# 9. Copy the compiled binary from builder stage
COPY --from=builder /app/notely .
# Copy static files if any
COPY --from=builder /app/static ./static

# 10. Expose the application port
EXPOSE 8080

# 11. Run the application
CMD ["./notely"]


## STAGE 1
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project to the container
COPY . .

# Download dependencies
RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/users ./services/users/users.go



## STAGE 2
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/users .

EXPOSE 8010
EXPOSE 8011
    
# Set the entrypoint for the microservicehow do i run
CMD ["./users"]
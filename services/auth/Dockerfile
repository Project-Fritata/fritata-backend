## STAGE 1
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project to the container
COPY . .

# Download dependencies
RUN go mod download

# Init swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app/services/auth
RUN swag init -pd -pdl 3 -g auth.go

# Build the binary
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/auth ./services/auth/auth.go



## STAGE 2
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/auth .

EXPOSE 8000
    
# Set the entrypoint for the microservicehow do i run
CMD ["./auth"]
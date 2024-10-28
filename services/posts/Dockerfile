## STAGE 1
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project to the container
COPY . .

# Download dependencies
RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/posts ./services/posts/posts.go



## STAGE 2
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/posts .

EXPOSE 8020
    
# Set the entrypoint for the microservicehow do i run
CMD ["./posts"]
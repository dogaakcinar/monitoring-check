# Stage 1: Build Stage
FROM golang:1.19.2-alpine AS build

WORKDIR /go/src/app

# Copy only the necessary files for dependency fetching
COPY go.mod go.sum ./

# Fetch dependencies if using modules
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the statically linked Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app

# Stage 2: Production Stage
FROM alpine:latest

# Copy the built binary from the previous stage
COPY --from=build /go/bin/app /app

# Set the working directory inside the container
WORKDIR /

# Expose a port if your Go application listens on a specific port
EXPOSE 8080

# Define the command to run your application
CMD ["/app"]

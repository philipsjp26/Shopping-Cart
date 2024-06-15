# Stage 1: Build stage
FROM golang:1.21-alpine AS build

RUN mkdir -p app
# Copy and download dependencies

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

RUN go mod tidy && go mod download && go mod vendor

# Build the Go application
RUN go build -o myapp .

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/myapp .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
CMD ["sh","-c", "/app/myapp serve"]
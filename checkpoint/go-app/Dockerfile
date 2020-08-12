FROM golang:1.14-alpine AS builder

# Set build environmet variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Copy the code into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o main cmd/main.go

WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM alpine:3.12

COPY --from=builder /dist/main /

# Command to run
ENTRYPOINT ["/main"]

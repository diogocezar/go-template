# STEP 1: Build the executable binary
FROM golang:alpine AS builder

# Install necessary tools: git for dependencies and upx for binary compression
RUN apk update && apk add --no-cache git upx

WORKDIR /app

# Copy the source code
COPY . .

# Fetch dependencies
RUN go mod download

# Build the binary with size optimization flags
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./main ./cmd/main.go

# Compress the binary with upx
RUN upx --best --lzma ./main

# STEP 2: Build a minimal runtime image
FROM scratch

# Copy the compressed binary from the builder stage
COPY --from=builder /app/main /main

# Run the binary
ENTRYPOINT ["/main"]

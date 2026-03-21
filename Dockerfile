# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /src

# Copy go.mod and go.sum first for better layer caching.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build a static binary.
RUN CGO_ENABLED=0 go build -o /out/api ./cmd/api

# Stage 2: Run
FROM scratch AS runtime

# Copy the binary from the builder stage.
COPY --from=builder /out/api /api

# Expose the port the API listens on.
EXPOSE 8080

# Run the application binary.
ENTRYPOINT ["/api"]
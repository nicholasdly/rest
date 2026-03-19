FROM golang:1.25-alpine AS build
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/api

FROM alpine:3.20
WORKDIR /
COPY --from=build /out/api /api

EXPOSE 8080
ENTRYPOINT ["/api"]
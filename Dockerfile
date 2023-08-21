#build stage
FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

#final stage
FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
COPY --from=builder /app .
ENTRYPOINT ["./app"]

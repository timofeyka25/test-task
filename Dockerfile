FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/app .

COPY ./migrations ./migrations

COPY config*.yaml /root/

RUN if [ ! -f /root/config.yaml ] && [ -f /root/config.example.yaml ]; then \
      cp /root/config.example.yaml /root/config.yaml; \
    fi

CMD ["./app"]

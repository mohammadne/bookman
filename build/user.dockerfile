FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /bin/app ./cmd/root.go

FROM alpine:latest

WORKDIR /bin/

COPY --from=builder /bin/app .

EXPOSE 8081

LABEL org.opencontainers.image.source="https://github.com/mohammadne/bookman-user"

ENTRYPOINT ["/bin/app"]

CMD ["server", "--env=dev"]

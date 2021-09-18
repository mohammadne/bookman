FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

RUN go mod download && make ent-generate

RUN go build -o /bin/app ./cmd/root.go

FROM alpine:latest

WORKDIR /bin/

COPY --from=builder /bin/app .

EXPOSE 8080

LABEL org.opencontainers.image.source="https://github.com/mohammadne/bookman-library"

ENTRYPOINT ["/bin/app"]

CMD ["server", "--env=dev"]

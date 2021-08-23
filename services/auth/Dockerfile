FROM golang:1.16.0 AS builder

WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./dist/bookman-auth ./cmd/root.go

FROM golang:1.16.0 AS bin

WORKDIR /app

# update and install dependency
RUN apt-get update

# Copy our static executable.
COPY --from=builder /app/.env /app/.env
COPY --from=builder /app/dist/bookman-auth /app/bookman-auth

# sourcing .env file to set environment-variables
RUN source .env

# removing .env file, because it contains secrets
RUN rm -f .env

VOLUME /app/downloads

ENTRYPOINT ["/app/bookman-auth"]

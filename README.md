<!-- LOGO -->
<p align="center">
  <img src="assets/logo.png" />
</p>

<!-- BADGES -->
<p align="center">
  <img src="https://img.shields.io/github/release/mohammadne/bookman.svg?style=for-the-badge">
  <img src="https://img.shields.io/codecov/c/gh/mohammadne/bookman?logo=codecov&style=for-the-badge">
  <img src="https://img.shields.io/github/license/mohammadne/bookman?style=for-the-badge">
  <img src="https://img.shields.io/github/stars/mohammadne/bookman?style=for-the-badge">
  <img src="https://img.shields.io/github/downloads/mohammadne/bookman/total.svg?style=for-the-badge">
</p>

<!-- TITLE -->
# BOOK MAN
> an online REST renting book platform which you can authenticate,
> order, reserve a book in your account.
>
> it's a microservices project with highly focus on code architecture.

## Microservices:

1. [`User`](https://img.shields.io) is a service for user managment

2. [`Auth`](https://img.shields.io) is bookman authentication service

3. [`Library`](https://img.shields.io) is bookman core service, responsible
for serving books and authors

<!-- 4. [`Rent`](https://img.shields.io) -->

<!-- 5. [`Notification`](https://img.shields.io) -->

## GUIDE

### install proto-buf tools

``` zsh
# protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 

# protoc-grpc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
```

### generate proto-buf files

``` zsh
protoc --proto_path=. --go_out=. service.proto
protoc --proto_path=. --go-grpc_out=. service.proto

# OR

protoc --proto_path=. --go_out=. --go-grpc_out=. service.proto
```

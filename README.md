<!-- TITLE -->
<h1 align="center">BOOK MAN</h1>

<!-- BADGES -->
<p align="center">
  <img src="https://img.shields.io/github/release/mohammadne/bookman.svg?style=for-the-badge">
  <img src="https://img.shields.io/codecov/c/gh/mohammadne/bookman?logo=codecov&style=for-the-badge">
  <img src="https://img.shields.io/github/license/mohammadne/bookman?style=for-the-badge">
  <img src="https://img.shields.io/github/stars/mohammadne/bookman?style=for-the-badge">
  <img src="https://img.shields.io/github/downloads/mohammadne/bookman/total.svg?style=for-the-badge">
</p>

<p align="center">
  <img src="assets/logo.png" />
</p>

# Introduction
> `Bookman` is an online REST renting book platform.
>
> it's microservices and uses clean-architecture

## Microservices:

1. `User Service` is user managment service

2. `Auth Service` is authentication service

3. `Library Service` will list all availables books

4. `Rent Service` is responsible for rent a book to a user (SOON)

5. `Notification Service` contains email and other means notification (SOON)

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

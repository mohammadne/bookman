<!-- TITLE -->
<h1 align="center">BOOK MAN</h1>

<p align="center">
  <img src="https://img.shields.io/github/workflow/status/mohammadne/bookman/test?label=test&logo=github&style=flat-square">
  <img src="https://img.shields.io/codecov/c/gh/mohammadne/bookman?logo=codecov&style=flat-square">
  <img src="https://img.shields.io/docker/image-size/mohammadne/bookman/latest?style=flat-square&logo=docker">
  <img src="https://img.shields.io/docker/pulls/mohammadne/bookman?style=flat-square&logo=docker">
  <img src="https://img.shields.io/docker/pulls/mohammadne/bookman?style=flat-square&logo=docker">
  <img src="https://pkg.go.dev/badge/github.com/mohammadne/bookman">
</p>

<!-- BADGES -->
<p align="center">
  <img src="https://img.shields.io/github/release/mohammadne/bookman.svg?style=for-the-badge">
  <img src="https://img.shields.io/github/license/mohammadne/bookman?style=for-the-badge">
  <img src="https://img.shields.io/github/stars/mohammadne/bookman?style=for-the-badge">
  <img src="https://img.shields.io/github/downloads/mohammadne/bookman/total.svg?style=for-the-badge">
</p>

<p align="center">
  <img src="assets/bookman.png" />
</p>

## Introduction
> `Bookman` is an online REST renting book platform.
>
> it's microservices and uses clean-architecture


### Ports
- `external` ports use HTTP-REST 
- `internal` ports use gRPC 

## Microservices:

- `User Service`
is user managment service

- `Auth Service`
is authentication service

- `Book Service`
will list all availables books

- `Rent Service`
is responsible for rent a book to a user

- `Notification Service`
contains email and other means notification


# BOOKMAN User API

## how to run
`go run app/*.go`

we will use MVC pattern:

1. Controller 
input layer which is responsible to provide endpoints to
interact with user api, it will handle the requests.
it will validate the request and if validations passed,
then it will pass the routine to the service

2. Service

we have no knowledge about web frameworks here

## DTO
an object that we will transfer from persistance to application layer and backward

## DAO
persistance and retrieve logic of DTO and it's an access level to database
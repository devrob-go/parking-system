# Parking System

Parking system is a gRPC gateway based project to create a structure of production grade software which can be use as micro service if needed. Will develop further over this.

## Prerequisites for development

- Go v1.19+
- Docker v20+
- Node.js v12.17+
- Docker Compose v2.23.2+

## Getting started

### Steps to run Parking System with docker

Using docker compose run these commands in terminal and wait until everything is initialized

```
$ go mod tidy
$ cd deploy/compose
$ docker compose up --build -d
```

that's it, you are good to go..

cURL example to test the project:
```
curl --location --request POST '127.0.0.1:8181/parking-lot/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Name": "New Parking Lot",
    "TotalSpace": 30
}'
```
Left out initial tasks:
- more test cases will be added
- data transforms will be more optimised and readable
- authentication system and more..
 
Note: Postman collection and enviroment has been added under directory `postman`
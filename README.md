# Golang Microservices with gRPC Protocol Buffer Communication
This is a Google Remote Procedure Call (gRPC) API built using Golang uses Protocol Buffer communication and built with makefile features hello request & response, and get user by id request and response.

## Technologies  
![Golang](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![gRPC](https://img.shields.io/badge/grpc-%2300ADD8.svg?style=for-the-badge&logo=grpc&logoColor=white)  


## Contain 2 Different Go Backend Services
### Server
* Receiving gRPC call from client and return message
### Client
* Calling gRPC server end points

## Proto File
### hello.proto
* Request and response hello message
###  user.proto
* Request user by id and response with requested user id

##  gRPC Server Endpoint
| Proto     | API Endpoint               | Description                                      | Services             |
|------------|----------------------------|--------------------------------------------------|-------------------|
| **hello**   | `localhost:50051/UserService/GetUser`                | Get user by ID                             | User Service    |
| **user**    | `localhost:50051/Greeter/Hello`            | Hello greeter with response         | Hello Service    |

# Getting Started
This guide will help you set up and run the golang server & client services

## Prerequisites
Ensure you have the following installed:

- [Golang](https://go.dev/dl/) (latest version)
- [Protobuf](https://github.com/protocolbuffers/protobuf/releases) (latest version)

### 1. Install Dependencies
Install the required packages:
```sh
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get google.golang.org/grpc/reflection
```

### 2.  Make Protobuf CLient and Server
```sh
make proto_client
make proto_server
```

### 3.  Run the Server
```sh
cd server
go run main.go
```

### 4.  Run the Client
```sh
cd client
go run main.go
```

# Galleries
![GetUser](https://github.com/user-attachments/assets/f02068a5-a3cb-4873-beb0-b5a25621ec62)
![sayHello](https://github.com/user-attachments/assets/61e2120c-6c9e-4fc6-a47f-84e3b7fb8a78)





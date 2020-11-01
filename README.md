# MruV API Go package

This repository contains code generated from https://github.com/MruV-RP/mruv-protos protobuf files.

## How to use this repository?
This repository is a go module package, so if you want to use MruV API package in your Go project, you should add this module to your project as a dependency by using this command in project root directory:
```
$ go get github.com/MruV-RP/mruv-pb-go
```

## Example code
This example will connect to MruV API on address 127.0.0.1:50051 and will run remote procedure IsAccountExist
```go
package main

import (
	"context"
	"log"

	"github.com/MruV-RP/mruv-pb-go/accounts"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := accounts.NewMruVAccountsServiceClient(conn)

	// Send gRPC request and print out its response.
	name := "Account Login"
	reply, err := c.IsAccountExist(context.Background(), &accounts.IsAccountExistRequest{Login: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if reply.Exists {
		log.Println("Account exists :)")
	} else {
		log.Println("There is no such account :(")
	}
}
```

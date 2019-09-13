package main

import (
	"fmt"
	"os"

	"github.com/TranThienChi/go-grpc-http-rest-microservice/pkg/cmd/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/sinomiko/brpc-go"
	"github.com/sinomiko/brpc-go/examples/echo"
	bstd "github.com/sinomiko/brpc-go/protocol/brpc-std"
)

func main() {
	clientConn, err := brpc.Dial(bstd.ProtocolName, "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	client := echo.NewEchoServerClient(clientConn)
	resp, err := client.Echo(context.Background(), &echo.EchoRequest{
		Message: "hello",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}

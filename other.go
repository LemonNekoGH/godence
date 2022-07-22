package godence

import (
	"context"
	"fmt"

	flowGrpc "github.com/onflow/flow-go-sdk/access/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var flowCli *flowGrpc.Client

func initFlowClient() {
	client, err := flowGrpc.NewClient(
		"localhost:3569",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	flowCli = client
	fmt.Println("Flow client init success.")
}

package main

import (
	pb "actorclient/proto/daprinternal"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}

	c := pb.NewDaprInternalClient(conn)
	_, err = c.CallActor(context.Background(), &pb.CallActorEnvelope{
		ActorType: "Cat",
		ActorID:   "Hobbit",
		Method:    "Eat",
	})
	if err != nil {
		fmt.Println(err)
	}
}

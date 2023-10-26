package main

import (
	// "google.golang.org/protobuf/proto"
	"context"
	"flag"
	// "fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"

  api "lanboard/protobufs"
)

var (
       addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
  score := &api.Score {
    Name: "Me",
    Epoch: int32(time.Now().Unix()),
  }

  flag.Parse();

  conn, err := grpc.Dial(*addr, grpc.WithInsecure());

  if err != nil {
    log.Fatalf("Did not connect: %v", err);
  }

  defer conn.Close();

  client := api.NewLeaderboardClient(conn);

  resp, err := client.Insert(context.Background(), score);

  if err != nil {
    log.Fatalf("Could not add score: %v", err);
  }

  log.Printf("Response: %v", resp);
}

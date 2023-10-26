package main

import (
  "context"
  "flag"
  "log"
  "time"

  "google.golang.org/grpc"

  api "lanboard/protobufs"
)

var (
  addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
  // Parse the command line arguments. 
  flag.Parse();

  // Connect to the gRPC server.
  conn, err := grpc.Dial(*addr, grpc.WithInsecure());

  //  Check for connection failure. 
  if err != nil {
    log.Fatalf("Did not connect: %v", err);
  }

  defer conn.Close();

  // Create a new client to talk to the leadboard.
  client := api.NewLeaderboardClient(conn);

  score := &api.Score {
    Name: "Me",
    Epoch: int32(time.Now().Unix()),
  }

  // Call the insert function to add a score
  resp, err := client.Insert(context.Background(), score);

  // Check gor add score failure
  if err != nil {
    log.Fatalf("Could not add score: %v", err);
  }

  // Print out whether it is okay or a fail. 
  log.Printf("Response: %v", resp.Result);
}

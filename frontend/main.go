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

type Leaderboard struct {
	conn   *grpc.ClientConn
	client api.LeaderboardClient
}

func NewLeaderboard() (*Leaderboard, error) {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return &Leaderboard{
		conn:   conn,
		client: api.NewLeaderboardClient(conn),
	}, nil
}

func (c *Leaderboard) InsertScore(score *api.Score) (*api.InsertionResponse, error) {
  return c.client.Insert(context.Background(), score)
}

func (c *Leaderboard) Close() error {
  return c.conn.Close()
}

func main() {
	// Parse the command line arguments.
	flag.Parse()

  client, err := NewLeaderboard();
  
  if err != nil {
    log.Fatalf("Did not c onnect: %v", err)
  }

  defer client.Close()

  score := &api.Score {
    Name: "Me",
    Epoch: int32(time.Now().Unix()),
  }

  resp, err := client.InsertScore(score)

  if err != nil {
    log.Fatalf("Failed to send message: %v", err)
  }

  if resp.Result != api.InsertionResult_OKAY {
    log.Fatalf("Result not okay: %v", resp.Result)
  }

  log.Println("Sent request.")
}

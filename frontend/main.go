package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	api "lanboard/protobufs"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	webaddr = flag.String("webaddr", "localhost:8000", "the address to host on")
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

func scoreHandler(client *Leaderboard) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    log.Println("Adding score!")

		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("name")
		epoch := int32(time.Now().Unix())

		resp, err := client.InsertScore(&api.Score{Name: name, Epoch: epoch})

		if err != nil {
			http.Error(w, "Failed to add score", http.StatusInternalServerError)
			return
		}

		if resp.Result != api.InsertionResult_OKAY {
			http.Error(w, "Result not okay: %v", http.StatusInternalServerError)
			return
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
  template := template.Must(template.ParseFiles("static/index.html"));

  template.Execute(w, nil);
}

func main() {
	flag.Parse()

	client, err := NewLeaderboard()

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer client.Close()

  fs := http.FileServer(http.Dir("static"))
  static := http.StripPrefix("/static/", fs)

	http.HandleFunc("/add-score/", scoreHandler(client))
  http.Handle("/static/", static)
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(*webaddr, nil))
}

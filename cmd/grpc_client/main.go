package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	desc "github.com/spv-dev/go-grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	host   = "localhost:50051"
	noteID = 12
)

func main() {
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %s", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Can't close connection. Err: %s", err)
		}
	}()

	c := desc.NewNoteV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("failed to get note by id: %s", err)
	}

	log.Printf(color.RedString("Note info: \n"), color.GreenString("%+v", r.GetNote()))
}

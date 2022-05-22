package main

import (
	"context"
	"flag"
	"log"

	"github.com/bricks-cloud/bricks-cli/pkg/cue"
	"github.com/bricks-cloud/bricks-cli/pkg/terraform"
)

var (
	filePath string
)

func init() {
	flag.StringVar(&filePath, "file", "", "the path of cue file")
}
func main() {
	flag.Parse()
	js, err := cue.MarshalJSON(filePath)
	if err != nil {
		log.Fatalf("Failed to Marshal cue to json: %v", err)
	}

	if err := terraform.Run(context.Background(), js); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successful Deploy!")
}

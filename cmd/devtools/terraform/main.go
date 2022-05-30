package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/bricks-cloud/bricks/pkg/cue"
	"github.com/bricks-cloud/bricks/pkg/terraform"
)

var (
	filePath string
	workdir  string
)

func init() {
	flag.StringVar(&filePath, "file", "", "the path of cue file")
	flag.StringVar(&workdir, "workdir", ".workdir", "the path of working dir")

}
func main() {
	flag.Parse()
	js, err := cue.MarshalJSON(filePath)
	if err != nil {
		log.Fatalf("Failed to Marshal cue to json: %v", err)
	}

	if err := os.MkdirAll(workdir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create work dir: %v", err)
	}

	if err := terraform.Run(context.Background(), workdir, js); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successful Deploy!")
}

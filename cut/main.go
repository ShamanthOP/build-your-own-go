package main

import (
	"context"
	"cut/cmd"
	"log"
	"os"
)

func main() {
	cmd := cmd.GetCmd()

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

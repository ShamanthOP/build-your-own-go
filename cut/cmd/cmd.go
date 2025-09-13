package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v3"
)

var cmd = &cli.Command{
	Name:        "cut",
	Description: "cut out selected portions of each line of a file",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "delimiter",
			Aliases: []string{"d"},
			Usage:   "Specify the delimiter for the cut command",
			Value:   "\t",
		},
		&cli.StringFlag{
			Name:    "fields",
			Aliases: []string{"f"},
			Usage:   "Specify the fields to cut",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {

		fields := c.String("fields")
		delimiter := c.String("delimiter")

		parts := strings.Split(fields, ",")
		if len(parts) == 1 {
			parts = strings.Fields(fields)
		}
		ints := make([]int, len(parts))
		for i, p := range parts {
			val, _ := strconv.Atoi(p)
			ints[i] = val
		}

		var reader *bufio.Reader
		if c.NArg() > 0 && c.Args().Get(0) != "-" {
			file, err := os.Open(c.Args().Get(0))
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot open file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()
			reader = bufio.NewReader(file)
		} else {
			reader = bufio.NewReader(os.Stdin)
		}

		for {
			line, err := reader.ReadString('\n')

			if err != nil && err != io.EOF {
				log.Fatal(err)
			}

			parts := strings.Split(line, delimiter)
			for i, field := range ints {
				if field < 1 || field > len(parts) {
					log.Fatalf("Invalid field number: %d", field)
				}
				fmt.Print(parts[field-1])
				if i < len(ints)-1 {
					fmt.Print(delimiter)
				} else {
					fmt.Println()
				}
			}

			if err == io.EOF {
				break
			}
		}

		return nil
	},
}

func GetCmd() *cli.Command {
	return cmd
}

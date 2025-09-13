package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "compressor",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(args)

		if len(args) == 0 {
			return fmt.Errorf("no command specified")
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

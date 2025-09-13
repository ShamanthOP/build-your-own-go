package cmd

import (
	"compressor/huffman"

	"github.com/spf13/cobra"
)

var decompressCmd = &cobra.Command{
	Use:   "decompress filename",
	Short: "Decompresses the provided file",
	Args:  cobra.ExactArgs(1),
	RunE:  decompress,
}

func init() {
	rootCmd.AddCommand(decompressCmd)
}

func decompress(cmd *cobra.Command, args []string) error {
	filename := args[0]

	err := huffman.Decode(filename)

	if err != nil {
		panic(err)
	}

	return nil
}

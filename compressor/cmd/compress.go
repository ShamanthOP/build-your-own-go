package cmd

import (
	"compressor/huffman"

	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress filename",
	Short: "Compresses the provided file",
	Args:  cobra.ExactArgs(1),
	RunE:  compress,
}

var outputFilename string

func init() {
	compressCmd.Flags().StringVarP(&outputFilename, "output", "o", "output.bin", "specify the output file name")
	rootCmd.AddCommand(compressCmd)
}

func compress(cmd *cobra.Command, args []string) error {
	filename := args[0]

	err := huffman.Encode(filename, outputFilename)

	if err != nil {
		panic(err)
	}

	return nil
}

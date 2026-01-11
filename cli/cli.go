package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Options struct {
	Input string
	Output string
	IncludeDir string
}

func NewRootCmd() (*cobra.Command, *Options) {
	opts := &Options{}
	cmd := &cobra.Command {
		Use: "blunderbuss",
		Short: "A compiler for blunderbuss",
	}

	cmd.Flags().StringVarP(&opts.Input, "input", "i", "", "source file")
	cmd.Flags().StringVarP(&opts.Output, "output", "o", "a.asm", "output file")
	cmd.Flags().StringVarP(&opts.IncludeDir, "include", "I", "", "include directory")

	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("include")

	return cmd, opts
}

func run(opts *Options) {
	fmt.Println(opts)
}


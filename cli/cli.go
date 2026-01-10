package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Options struct {
	Input string
	Output string
}

func NewRootCmd() (*cobra.Command, *Options) {
	opts := &Options{}
	cmd := &cobra.Command {
		Use: "blunderbuss",
		Short: "A compiler for blunderbuss",
	}

	cmd.Flags().StringVar(&opts.Input, "input", "", "source file")
	cmd.Flags().StringVar(&opts.Output, "output", "a.asm", "output file")

	cmd.MarkFlagRequired("input")

	return cmd, opts
}

func run(opts *Options) {
	fmt.Println(opts)
}


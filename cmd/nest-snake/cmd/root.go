package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDefaultSnakeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "nest-snake",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("SsssSSssSsSSss")
		},
	}
	return cmd
}

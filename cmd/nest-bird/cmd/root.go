package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDefaultBirdCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "nest-bird",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Piou piou")
		},
	}
	return cmd
}

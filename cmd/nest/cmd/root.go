package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"go.zenithar.org/nest/pkg/cmd/plugin"
	"go.zenithar.org/nest/pkg/handler"
)

func newNestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "nest",
	}
	cmd.AddCommand(plugin.NewCmdPlugin())
	return cmd
}

func NewDefaultNestCommand() *cobra.Command {
	cmd := newNestCommand()

	pluginHandler := handler.NewDefaultPluginHandler([]string{"nest"})
	if len(os.Args) > 1 {
		cmdPathPieces := os.Args[1:]

		// only look for suitable extension executables if
		// the specified command does not already exist
		if _, _, err := cmd.Find(cmdPathPieces); err != nil {
			if err := handler.HandlePluginCommand(pluginHandler, cmdPathPieces); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		}
	}

	return cmd
}

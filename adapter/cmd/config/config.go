package config

import (
	"github.com/spf13/cobra"
)

// AddSubCommands adds the sub-commands of config command.
func AddSubCommands(root *cobra.Command) {
	root.AddCommand(printCmd)
}

package config

import (
	"fmt"

	"github.com/blackhorseya/assessment-bito/pkg/configx"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(configx.C.String()) //nolint:forbidigo // it's just for print
	},
}

package restful

import (
	"github.com/blackhorseya/assessment-bito/pkg/adapterx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCmd is to create a new restful command.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "start a restful service",
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.GetViper()

			var (
				service adapterx.Servicer
				err     error
			)
			if v.GetString("gods") == "memory" {
				service, err = NewMemory(v)
				cobra.CheckErr(err)
			} else {
				service, err = NewRBTree(v)
				cobra.CheckErr(err)
			}

			err = service.Start()
			cobra.CheckErr(err)

			err = service.AwaitSignal()
			cobra.CheckErr(err)
		},
	}
	cmd.Flags().String("gods", "rbtree", "gods type(memory, rbtree)")

	return cmd
}

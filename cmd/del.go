package cmd

import (
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:     "del [key]",
	Aliases: []string{"clear", "rm"},
	Short:   "Clear out the value of a key",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// NOTE: viper doesn't implement an Unset function, yet.
		// See: https://github.com/spf13/viper/pull/519
		key := args[0]
		setCmd.Run(cmd, []string{key, ""})
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}

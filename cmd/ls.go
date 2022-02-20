package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"all", "keys", "list"},
	Short:   "List all keys",
	Run: func(cmd *cobra.Command, args []string) {
		for _, key := range viper.AllKeys() {
			fmt.Println(key)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

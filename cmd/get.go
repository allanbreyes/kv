package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get the value for a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		val := viper.Get(key)
		if val != nil {
			fmt.Println(val)
		} else {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

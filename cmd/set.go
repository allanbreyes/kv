package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:     "set [key] [value]",
	Aliases: []string{"add", "put"},
	Short:   "Set a value to a key",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		val := args[1]
		viper.Set(key, val)
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(val)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

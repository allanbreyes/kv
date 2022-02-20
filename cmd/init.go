package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		handleConfig()
	},
}

func handleConfig() {
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = viper.ReadInConfig()
	fmt.Println(viper.ConfigFileUsed())
}

func init() {
	rootCmd.AddCommand(initCmd)
}

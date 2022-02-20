package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Run: func(cmd *cobra.Command, args []string) {
		handleConfig()
	},
}

func handleConfig() {
	err := viper.WriteConfigAs(path.Join(ConfigDir, fmt.Sprintf("%s.%s", DefaultDataStore, "yml")))
	_ = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(viper.ConfigFileUsed())
}

func init() {
	rootCmd.AddCommand(initCmd)
}

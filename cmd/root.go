package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DefaultDataStore = ".kv.yml"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kv",
	Short: "A key-value store CLI",
	Long: `A lightweight key-value store command-line program.

Example:
  $ kv put foo bar
  bar

  $ kv get foo
  bar
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	inputDataStore string
)

func setConfig(input string) string {
	viper.SetConfigType("yml")
	if input != "" {
		viper.SetConfigFile(input)
		viper.AddConfigPath(filepath.Dir(input))
	} else {
		dir, err := os.UserHomeDir()
		if err != nil {
			dir, err = os.UserConfigDir()
		}
		if err != nil {
			dir, err = os.Getwd()
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.SetConfigFile(path.Join(dir, DefaultDataStore))
	}

	if err := viper.ReadInConfig(); err != nil {
		return ""
	}
	return viper.ConfigFileUsed()
}

func initConfig() {
	setConfig(inputDataStore)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&inputDataStore,
		"file",
		"",
		fmt.Sprintf("data file (default is $HOME/%s.yml", DefaultDataStore),
	)
}

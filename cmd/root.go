package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DefaultDataStore = ".kv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kv",
	Short: "A key-value store CLI",
	Long: `A lightweight key-value store command-line program, powered by SQLite.

Example:
  $ kv put foo bar
  bar

  $ kv get foo
  bar
`,
	// Run: func(cmd *cobra.Command, args []string) { },
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
	ConfigDir      string
	inputDataStore string
)

func initConfig() {
	viper.SetConfigName(DefaultDataStore)
	viper.SetConfigType("yml")
	if inputDataStore != "" {
		viper.SetConfigFile(inputDataStore)
	}

	var err error
	ConfigDir, err = os.UserHomeDir()
	if err != nil {
		ConfigDir, err = os.UserConfigDir()
	}
	if err != nil {
		ConfigDir, err = os.Getwd()
	}

	viper.AddConfigPath(ConfigDir)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
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

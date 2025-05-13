package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wyrth-io/whit/internal/logging"
)

var cfgFile string

var whitVersion = "0.2.3"

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "whit",
	Short: "Whit is a an awesome app that does magic things",
	Long: `Whit is a an awesome app that does magic things

  The --loglevel can be set to one of:
  * info
  * warning
  * error
  * fatal
  * debug
  `,
	// this is what is run if no subcommand or arguments have been given
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println(whitVersion)
		} else {
			cmd.HelpFunc()(cmd, args)
		}
	}, PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logging.SetLevelFromName(Loglevel)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if !logging.Loglevel(Loglevel) {
			return fmt.Errorf("loglevel '%s' is not a recognized loglevel", Loglevel)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Loglevel set on command line (defaults to "warning").
var Loglevel string
var showVersion bool

func init() {
	cobra.OnInitialize(initConfig)

	// support persisted flags - global for this application and allow giving config file as option
	flags := RootCmd.PersistentFlags()
	flags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.whit.yaml)")
	flags.StringVar(&Loglevel, "loglevel", "warning", "sets log filtering to this level and above")

	RootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Prints the current version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".whit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".whit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Common cmd line flags/values.

// Format is used to accept the name of a format. Different commands may accept different formats, so this is
// just a variable declaration.
var format string

// Outfile is used to get the name of a file to write to in commands that have such an operation.
var outFile string

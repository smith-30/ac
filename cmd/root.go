package cmd

import (
	"fmt"
	"os"

	"github.com/smith-30/acc/color"
	"github.com/spf13/cobra"
)

var (
	// this var is setting by Makefile
	appName = ""
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// for logging panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(color.Redf("sorry unexpected error is occurred. please report to the repository. err: %s", err))
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// common settings ...
}

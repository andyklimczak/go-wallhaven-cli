/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-wallhaven",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-wallhaven.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&Destination, "destination", "d", "~/.wallpapers", "destination directory")
	rootCmd.PersistentFlags().StringVarP(&ApiKey, "apikey", "a", "", "destination directory")
	rootCmd.PersistentFlags().StringVarP(&Username, "username", "u", "", "username of the wallhaven user who owns the collection")
	rootCmd.PersistentFlags().StringVarP(&CollectionLabel, "collection-label", "c", "", "collection label")
	threads := runtime.NumCPU()
	downloadCmd.Flags().IntVarP(&Threads, "threads", "t", threads, "number of threads")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.MarkPersistentFlagRequired("collection-label")
}

var Verbose bool
var Destination string
var ApiKey string
var Username string
var CollectionLabel string
var Threads int

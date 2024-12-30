/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/andyklimczak/go-wallhaven/internal/downloader"
	"github.com/andyklimczak/go-wallhaven/internal/logger"
	"github.com/andyklimczak/go-wallhaven/internal/wallhaven"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download content from a Wallhaven collection",
	Long:  `gohaven download -username "user" -collection-label "Desktop"`,
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.DefaultGoHavenLogger(Verbose)
		client := wallhaven.NewClient(Username, ApiKey, log)
		collections, err := client.CollectionsForApikey()
		if err != nil {
			log.Fatal("Unable to get collections: %w", err)
		}
		collection, err := collections.GetByLabel(CollectionLabel)
		if err != nil {
			log.Fatal("Unable to get collection by label: %w", err)
		}
		searchResult, err := client.ListingsForCollection(collection)
		if err != nil {
			log.Fatal("Unable to get listings for collection: %w", err)
		}
		dwnldr := downloader.NewDefaultWallhavenDownloader(Threads, log)
		err = dwnldr.DownloadFromCollection(collection, searchResult, Destination)
		if err != nil {
			log.Fatal("Unable to download collection: %w", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

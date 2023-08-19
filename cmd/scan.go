/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/eargollo/dupfinder/internal/dupfinder"
	"github.com/eargollo/dupfinder/pkg/dupfile"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan one or more folders for duplicated files",
	Long: `Scans one or more folders for duplicated files.
Pass each folder as a separate argument such as:
dupfinder scan /first/folder  /another/folder
`,
	Run: func(cmd *cobra.Command, args []string) {
		cachePath, err := dupfinder.CachePath()
		if err != nil {
			log.Fatalf("could not execute: %v", err)
		}

		df, err := dupfile.New(dupfile.WithPaths(args), dupfile.WithCache(cachePath))
		if err != nil {
			log.Fatal(err)
		}

		result := df.Run()

		fmt.Println("\n\nDuplicates list by size:")
		for i, dup := range result {
			fmt.Printf("Duplicate %d Size %d Files %d MD5 %x\n", i, dup[0].Size, len(dup), dup[0].Hash)
			for _, file := range dup {
				fmt.Printf("\t%s\n", file.Path)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

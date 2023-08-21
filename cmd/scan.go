/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

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
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatalf("invalid output file '%s': %v", output, err)
		}

		//Create output file
		filename := dupfinder.OutputFileName(output)
		log.Printf("Creating output file '%s", filename)
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("could not create output file '%s", err)
		}
		defer f.Close()

		cachePath, err := dupfinder.CachePath()
		if err != nil {
			log.Fatalf("could not execute: %v", err)
		}

		df, err := dupfile.New(dupfile.WithPaths(args), dupfile.WithCache(cachePath))
		if err != nil {
			log.Fatal(err)
		}

		result := df.Run()

		fmt.Fprintln(f, "Duplicates list by size:")
		for i, dup := range result {
			fmt.Fprintf(f, "Duplicate %d Size %d Files %d MD5 %x\n", i, dup[0].Size, len(dup), dup[0].Hash)
			for _, file := range dup {
				fmt.Fprintf(f, "    '%s'\n", file.Path)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.PersistentFlags().String("output", "duplicates.txt", "Set output file for duplicate list. Default is duplicates.txt. Iterates number in file if output file exists.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

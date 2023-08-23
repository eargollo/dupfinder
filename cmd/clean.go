/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/eargollo/dupfinder/internal/cleaner"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Scan one or more folders for duplicated files",
	Long: `Scans one or more folders for duplicated files.
Pass each folder as a separate argument such as:
dupfinder scan /first/folder  /another/folder
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Pass the file with choices as argument")
			return
		}
		if len(args) > 1 {
			fmt.Println("Pass only one file at a time")
			return
		}

		cl := cleaner.New()

		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			log.Fatalf("error reading dry-run flag: %v", err)
		}

		err = cl.LoadFile(args[0])
		if err != nil {
			log.Fatalf("Error processing file %s: %v", args[0], err)
		}
		actions := cl.Actions()

		for i, action := range actions {
			if action.Type != cleaner.Keep {
				log.Printf("%d: Group %d - Action %s '%s'", i, action.Group, strings.ToUpper(action.Type.String()), action.Filename)
				if !dryRun {
					err := action.Execute()
					if err != nil {
						log.Printf("ERROR: Could not perform action %s to file '%s': %v", action.Type.String(), action.Filename, err)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.PersistentFlags().Bool("dry-run", false, "Dry run mode, show the actions that would be taken only.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

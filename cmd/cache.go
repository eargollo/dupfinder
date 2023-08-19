/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"sort"
	"strings"

	"github.com/eargollo/dupfinder/internal/dupfinder"
	"github.com/eargollo/dupfinder/pkg/dupfile"
	"github.com/spf13/cobra"
)

// cacheCmd represents the cache command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage file signature (MD5) cache ",
	Long: `dupfinder caches the file signature (MD5) based on file absolute path
and file size. This approach speeds up multiple executions. The cache might need
to be cleaned in case file changes between executions without changing size. This
command allows you to inspect and clean the cache.
For example:
- Get information on the cache:
   dupfinder cache
- Clean the cache:
   dupfinder cache --clean
`,
	Run: func(cmd *cobra.Command, args []string) {
		cachePath, err := dupfinder.CachePath()
		if err != nil {
			log.Fatalf("could not execute: %v", err)
		}

		db, err := dupfile.NewMD5Cache(cachePath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		log.Printf("Cache DB path %s", cachePath)
		log.Printf("Cache has %d items", db.Count())

		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		if verbose {
			list := db.List([]string{})
			logList(list)
		}

		clean, err := cmd.Flags().GetBool("clean")
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		if clean {
			if len(args) == 0 {
				log.Printf("Add prefix to be cleaned as argument. To clean all run 'dupfinder cache clean /'")
				return
			}
			log.Printf("Cleaning cache with prefixes %s", strings.Join(args, ","))

			total := db.Count()
			toDelete := db.List(args)
			if verbose {
				log.Printf("Deleting %d items", len(toDelete))
				logList(toDelete)
			}
			db.Delete(toDelete)
			after := db.Count()
			log.Printf("Cleaning completed: %d items cleaned (from %d to %d remaining)", total-after, total, after)
		}
	},
}

func init() {
	rootCmd.AddCommand(cacheCmd)
	cacheCmd.PersistentFlags().Bool("clean", false, "Clean files cache")
	cacheCmd.PersistentFlags().Bool("verbose", false, "Verbose mode, showing cache items")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func logList(list []string) {
	sort.StringSlice.Sort(list)
	for _, key := range list {
		log.Print(key)
	}
}

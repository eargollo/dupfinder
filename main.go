package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eargollo/dupfinder/pkg/dupfile"
)

func main() {
	paths := os.Args[1:]
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Think of a way to avoid this becoming stale
	cacheDir := homedir + "/.dupfile"
	df, err := dupfile.New(dupfile.WithPaths(paths), dupfile.WithCache(cacheDir))
	if err != nil {
		log.Fatal(err)
	}

	result := df.Run()

	fmt.Println("\n\nDuplicates list by size:")
	for i, dup := range result {
		fmt.Printf("Duplicate %d Size %d Files %d MD5 %x\n", i, dup[0].Size, len(dup), dup[0].Md5)
		for _, file := range dup {
			fmt.Printf("\t%s\n", file.Path)
		}
	}
}

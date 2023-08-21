package dupfinder

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/eargollo/dupfinder/pkg/dupfile"
)

func CachePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return homedir, err
	}

	cacheDir := homedir + "/.dupfile"
	return cacheDir, nil
}

func OutputFileName(file string) string {
	// If file doesn't exist, returns file
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		// file does not exist
		return file
	}

	// If it exists iterate on numbers until first free number is found
	extension := filepath.Ext(file)
	base := strings.TrimSuffix(file, extension)
	name := filepath.Base(base)
	pattern := base + "_*" + extension

	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalf("Could not set an output file: %v", err)
	}

	reg := regexp.QuoteMeta(name+"_") + "([0-9]+)" + regexp.QuoteMeta(extension)
	r, err := regexp.Compile(reg)
	if err != nil {
		log.Fatalf("Could not set an output file: %v", err)
	}

	max := 0
	for _, f := range files {
		res := r.FindStringSubmatch(f)
		if len(res) > 0 {
			val, _ := strconv.Atoi(res[1])
			if val > max {
				max = val
			}
		}
	}

	return fmt.Sprintf("%s_%d%s", base, max+1, extension)
}

type SummaryResult struct {
	Duplicates int
	Size       int64
	Groups     int
}

func Summary(result [][]*dupfile.File) SummaryResult {
	sum := SummaryResult{}
	sum.Groups = len(result)
	for _, dups := range result {
		tot := len(dups)
		sum.Duplicates += tot - 1
		sum.Size += int64(tot-1) * dups[0].Size
	}
	return sum
}

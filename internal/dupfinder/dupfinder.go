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

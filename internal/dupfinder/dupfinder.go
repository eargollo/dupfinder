package dupfinder

import "os"

func CachePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return homedir, err
	}

	cacheDir := homedir + "/.dupfile"
	return cacheDir, nil
}

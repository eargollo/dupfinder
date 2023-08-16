package dupfile

import (
	"crypto/md5"
	"io"
	"os"
)

type File struct {
	Path string
	Name string
	Size int64
	md5  []byte
}

func (f File) AbsPath() string {
	return f.Path
}

func (fl *File) MD5(cache *MD5Cache) ([]byte, error) {
	if len(fl.md5) == 0 {
		// Check if MD5 is in cache
		if cache != nil {
			md5 := cache.Get(fl.Path, fl.Size)
			if md5 != nil {
				fl.md5 = md5
				return fl.md5, nil
			}
		}

		// Calculate and store MD5
		f, err := os.Open(fl.AbsPath())
		if err != nil {
			return []byte{}, err
		}
		defer f.Close()

		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			return []byte{}, err
		}
		fl.md5 = h.Sum(nil)
		//Add to cache
		if cache != nil {
			cache.Put(fl)
		}
	}

	return fl.md5, nil
}

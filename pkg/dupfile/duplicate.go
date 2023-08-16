package dupfile

import "fmt"

type Dedup struct {
	paths     []string
	cachePath string
	cache     *MD5Cache
}

func New(opts ...DedupOption) (*Dedup, error) {
	// Default
	d := Dedup{paths: []string{"."}}

	for _, opt := range opts {
		opt(&d)
	}

	// Cache
	if d.cachePath != "" {
		var err error
		d.cache, err = NewMD5Cache(d.cachePath)
		if err != nil {
			return &d, fmt.Errorf("could not initialize cache: %w", err)
		}
	}
	return &d, nil
}

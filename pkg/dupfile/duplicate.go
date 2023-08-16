package dupfile

type Dedup struct {
	paths []string
}

func New(opts ...DedupOption) *Dedup {
	// Default
	d := Dedup{paths: []string{"."}}
	for _, opt := range opts {
		opt(&d)
	}
	return &d
}

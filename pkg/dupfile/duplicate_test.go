package dupfile

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		opts []DedupOption
		want *Dedup
	}{
		{name: "Defaults", opts: []DedupOption{}, want: &Dedup{paths: []string{"."}}},
		{name: "With paths", opts: []DedupOption{WithPaths([]string{"/Home", "/Volumes"})}, want: &Dedup{paths: []string{"/Home", "/Volumes"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

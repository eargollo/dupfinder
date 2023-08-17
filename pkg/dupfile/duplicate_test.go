package dupfile

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name    string
		opts    []DedupOption
		want    *Dedup
		wantErr bool
	}{
		{name: "Defaults", opts: []DedupOption{}, want: &Dedup{paths: []string{"."}}, wantErr: false},
		{
			name:    "With paths",
			opts:    []DedupOption{WithPaths([]string{"/Home", "/Volumes"})},
			want:    &Dedup{paths: []string{"/Home", "/Volumes"}},
			wantErr: false,
		},
		{
			name: "With cache",
			opts: []DedupOption{
				WithPaths([]string{"/Home", "/Volumes"}),
				WithCache(tempDir),
			},
			want:    &Dedup{paths: []string{"/Home", "/Volumes"}, cachePath: tempDir},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

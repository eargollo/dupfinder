package dupfinder

import "testing"

func TestOutputFileName(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"No files", "duplicates.txt", "duplicates.txt"},
		{"Default", "./testdata/test1/duplicates.txt", "./testdata/test1/duplicates_1.txt"},
		{"Multiple", "./testdata/test2/duplicates.txt", "./testdata/test2/duplicates_3.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OutputFileName(tt.path); got != tt.want {
				t.Errorf("OutputFileName() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

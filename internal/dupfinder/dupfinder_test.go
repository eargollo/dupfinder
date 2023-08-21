package dupfinder

import (
	"reflect"
	"testing"

	"github.com/eargollo/dupfinder/pkg/dupfile"
)

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

func TestSummary(t *testing.T) {
	tests := []struct {
		name   string
		result [][]*dupfile.File
		want   SummaryResult
	}{
		{name: "No duplicates", result: [][]*dupfile.File{}, want: SummaryResult{}},
		{
			name: "Multi duplicates",
			result: [][]*dupfile.File{
				[]*dupfile.File{
					&dupfile.File{Size: 800},
					&dupfile.File{Size: 800},
					&dupfile.File{Size: 800},
				},
				[]*dupfile.File{
					&dupfile.File{Size: 5000},
					&dupfile.File{Size: 5000},
				},
				[]*dupfile.File{
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
					&dupfile.File{Size: 1000},
				},
			},
			want: SummaryResult{Duplicates: 12, Size: 15600, Groups: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Summary(tt.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Summary() = %v, want %v", got, tt.want)
			}
		})
	}
}

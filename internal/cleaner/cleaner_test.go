package cleaner

import (
	"reflect"
	"testing"
)

func TestCleaner_load(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []Action
		wantErr bool
	}{
		{name: "Nothing", content: "Nothing here", want: []Action{}, wantErr: false},
		{
			name: "Keep files",
			content: `Found 1 duplicate groups with 2 duplicated files taking 7.4 GiB of extra space:
Duplicate 0 Size 732631040 Files 2 MD5 8cb9641ad22300dbec61e0299116911a8275018afa35d92e6f213dc40464375b
[]    '/files/a/filea.avi'
[]    '/files/b/filea.avi'
[]    '/files/a/filec.avi'
`,
			want: []Action{
				{Group: 0, Filename: "/files/a/filea.avi", Type: Keep, To: ""},
				{Group: 0, Filename: "/files/b/filea.avi", Type: Keep, To: ""},
				{Group: 0, Filename: "/files/a/filec.avi", Type: Keep, To: ""},
			},
			wantErr: false},
		{
			name: "Err missing file",
			content: `Found 1 duplicate groups with 2 duplicated files taking 7.4 GiB of extra space:
Duplicate 0 Size 732631040 Files 2 MD5 8cb9641ad22300dbec61e0299116911a8275018afa35d92e6f213dc40464375b
[]    '/files/a/filea.avi'
[]
[]    '/files/a/filec.avi'
`,
			want: []Action{
				{Group: 0, Filename: "/files/a/filea.avi", Type: Keep, To: ""},
			},
			wantErr: true,
		},
		{
			name: "Keep and delete files",
			content: `Found 1 duplicate groups with 2 duplicated files taking 7.4 GiB of extra space:
Duplicate 0 Size 732631040 Files 2 MD5 8cb9641ad22300dbec61e0299116911a8275018afa35d92e6f213dc40464375b
[]    '/files/a/filea.avi'
[]    '/files/b/filea.avi'
[]    '/files/a/filec.avi'
Duplicate 1 Size 9183 Files 4 MD5 8cb9641ad22300dbec61e0299116911a8275018afa35d92e6f213dc40464375b
[K]    '/files/a/filea.avi'
[D]    '/files/b/filea.avi'
[d]    '/files/a/filec.avi'
`,
			want: []Action{
				{Group: 0, Filename: "/files/a/filea.avi", Type: Keep, To: ""},
				{Group: 0, Filename: "/files/b/filea.avi", Type: Keep, To: ""},
				{Group: 0, Filename: "/files/a/filec.avi", Type: Keep, To: ""},
				{Group: 1, Filename: "/files/a/filea.avi", Type: Keep, To: ""},
				{Group: 1, Filename: "/files/b/filea.avi", Type: Delete, To: ""},
				{Group: 1, Filename: "/files/a/filec.avi", Type: Delete, To: ""},
			},
			wantErr: false},
		{
			name: "Bug #13",
			content: `Found 29 duplicate groups with 84 duplicated files taking 5.0 MiB of extra space:
Duplicate 0 Size 130677 Files 27 MD5 4dec5ec062e0e0aacd7adfe827c632120568c57cbdfa566040551003a14840d0
[d]    '/files/a.jpg'
[d]    '/files/b.jpg'
[d]    '/files/c.jpg'`,
			want: []Action{
				{Group: 0, Filename: "/files/a.jpg", Type: Delete, To: ""},
				{Group: 0, Filename: "/files/b.jpg", Type: Delete, To: ""},
				{Group: 0, Filename: "/files/c.jpg", Type: Delete, To: ""},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			err := c.Load(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(c.actions, tt.want) {
				t.Errorf("New() = %v, want %v", c.actions, tt.want)
			}
		})
	}
}

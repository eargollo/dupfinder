package dupfile

import (
	"reflect"
	"testing"
)

func TestNewMD5Cache(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantErr bool
	}{
		{name: "Invalid file", args: args{""}, wantNil: true, wantErr: true},
		{name: "Valid path", args: args{t.TempDir()}, wantNil: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMD5Cache(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMD5Cache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("NewMD5Cache() = %v, wantNil %v", got, tt.wantNil)
			}
		})
	}
}

func TestCache(t *testing.T) {
	c, err := NewMD5Cache(t.TempDir())
	if err != nil {
		t.Fatalf("Could not create cache: %v", err)
	}

	aFile := &File{Path: "/my/file", Name: "file", Size: 100, Md5: []byte("abc")}
	bFile := &File{Path: "/my/other/file", Name: "file", Size: 200, Md5: []byte("efg")}
	c.Put(aFile)
	c.Put(bFile)

	res := c.Get("/not/in/there", 50)
	if res != nil {
		t.Errorf("should return nil when file does not exist")
	}

	res = c.Get("/my/file", 100)
	if !reflect.DeepEqual(res, aFile.Md5) {
		t.Errorf("Get() = %v, want %v", res, aFile)
	}

	res = c.Get("/my/other/file", 200)
	if !reflect.DeepEqual(res, bFile.Md5) {
		t.Errorf("Get() = %v, want %v", res, bFile)
	}

	res = c.Get("/my/file", 101)
	if res != nil {
		t.Error("should get nil if size doesn't match")
	}

	res = c.Get("/my/file", 100)
	if res != nil {
		t.Error("should get nil after a mismatch")
	}
}

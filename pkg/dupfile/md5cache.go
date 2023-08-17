package dupfile

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/akrylysov/pogreb"
)

type MD5Cache struct {
	db *pogreb.DB
}

type CacheRecord struct {
	Size int64  `json:"size"`
	Hash []byte `json:"hash"`
}

func NewMD5Cache(file string) (*MD5Cache, error) {
	db, err := pogreb.Open(file, nil)
	if err != nil {
		// log.Fatal(err)
		return nil, fmt.Errorf("error openning or creating cache database %s: %w", file, err)
	}

	return &MD5Cache{db: db}, nil
}

func (c *MD5Cache) Close() {
	c.db.Close()
}

func (c *MD5Cache) Get(path string, size int64) []byte {
	record := CacheRecord{}
	data, _ := c.db.Get([]byte(path))
	if data == nil {
		return nil
	}
	// deserialize
	err := json.Unmarshal(data, &record)
	if err != nil {
		return nil
	}
	if record.Size != size {
		c.db.Delete([]byte(path))
		return nil
	}
	return record.Hash
}

func (c *MD5Cache) Put(file *File) {
	record := CacheRecord{Size: file.Size, Hash: file.Md5}
	data, _ := json.Marshal(record)
	_ = c.db.Put([]byte(file.Path), data)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

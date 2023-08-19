package dupfile

import (
	"encoding/json"
	"fmt"
	"strings"

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

func (c *MD5Cache) List(filters []string) []string {
	compare := len(filters) != 0
	keys := []string{}

	iterator := c.db.Items()
	k, _, err := iterator.Next()
	for err == nil {
		key := string(k)
		if compare {
			for _, filter := range filters {
				if strings.HasPrefix(key, filter) {
					keys = append(keys, key)
					break
				}
			}
		} else {
			keys = append(keys, key)
		}
		// strings.HasPrefix(s, prefix)
		k, _, err = iterator.Next()
	}

	return keys
}

func (c *MD5Cache) Delete(files []string) {
	for _, file := range files {
		key := []byte(file)
		data, _ := c.db.Get(key)
		if data != nil {
			c.db.Delete(key)
		}
	}
}

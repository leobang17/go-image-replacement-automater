package cache

import (
	"bufio"
	"path/filepath"
	"time"
)

const (
	timestampLayout = time.RFC3339
)

type cache struct {
	cacheDir string 
	cacheFilePath string
}

func Cache() *cache {
	cacheDir := filepath.Join(".", "cached")
	cacheFilePath := filepath.Join(cacheDir, "go-image-replacement-automater")

	return &cache {
		cacheDir: cacheDir,
		cacheFilePath: cacheFilePath,
	}
}

func (c *cache) ReadCache() time.Time {
	cacheFile, err := c.openCacheFile()
	if err != nil {
		panic(err)
	}
	defer cacheFile.Close()

	scanner := bufio.NewScanner(cacheFile)
	var firstLine string
	if scanner.Scan() {
		firstLine = scanner.Text()	
	}
	cachedTime, _ := time.Parse(timestampLayout, firstLine)
	return cachedTime
}

func (c *cache) WriteCache() error {

	return nil
}
package cache

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	timestampLayout string 	= time.RFC3339
	tempFilePath string 		= "tempfile"
)

type cache struct {
	cacheDir string 
	cacheFilePath string
}

func Cache() *cache {
	cacheDir := filepath.Join(".", ".cached")
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
	if scanner.Scan() {
		cachedTime, _ := time.Parse(timestampLayout, scanner.Text())
		return cachedTime
	}
	
	return time.Time{}
}

func (c *cache) WriteCache() error {
	cacheFile, err := c.openCacheFile()
	if err != nil {
		fmt.Println("Cache update failed while opening cache file.")
		return err
	}
	defer cacheFile.Close()

	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		fmt.Println("Cache update failed while updating cache.")
		return err
	}
	defer tempFile.Close()

	var updateErr error
	defer c.rollback(tempFile, &updateErr)
	c.updateCache(cacheFile, tempFile)
	c.flush(tempFile)

	return updateErr
}
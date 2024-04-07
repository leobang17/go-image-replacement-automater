package cache

import (
	"fmt"
	"os"
)

func (c *cache) openCacheFile() (*os.File, error) {
	if err := c.checkCacheDirIfNotCreate(); err != nil {
		return nil, err
	}

	if _, err := os.Stat(c.cacheFilePath); os.IsNotExist(err) {
		if cacheFile, err := os.Create(c.cacheFilePath); err != nil {
			return nil, fmt.Errorf("crashed during making cache file: %v", err)
		}	else {
			return cacheFile, nil
		}
	}	

	if cacheFile, err := os.Open(c.cacheFilePath); err != nil {
		return nil, fmt.Errorf("crashed opening cache file: %v", err)
	} else {
		return cacheFile, nil
	}
}

func (c *cache) checkCacheDirIfNotCreate() error {
	if _, err := os.Stat(c.cacheDir); os.IsNotExist(err) {
		if err := os.Mkdir(c.cacheDir, 0755); err != nil {
			return err
		}
 	}	
	return nil
}
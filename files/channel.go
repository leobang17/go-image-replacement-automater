package files

import (
	"os"
	"path/filepath"
	"time"
)

type FileMeta struct {
	os.FileInfo
	Path string
}

func FileModified(lastRuntime time.Time) <-chan FileMeta {
	fileChan := make(chan FileMeta, 100)
	checker := NewExtChecker(".md", ".mdx")

	go func() {
		defer close(fileChan)
		filepath.Walk("./.document", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
      
			if checker.MatchWithFilepath(path) && info.ModTime().After(lastRuntime) {
				fileChan <- FileMeta{ 
					Path: path,
					FileInfo: info,
				}
			}
			return nil
		})
	}()

	return fileChan
}
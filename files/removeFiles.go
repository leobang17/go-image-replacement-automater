package files

import (
	"errors"
	"os"
	"sync"
)

type FilePathHolder interface {
	FilePath() string
}

func RemoveFiles[T FilePathHolder](files []T) error {
	wg := &sync.WaitGroup{}
	var err error

	for _, file := range files {
		wg.Add(1)
		go func(f FilePathHolder) {
			defer wg.Done()
			err = os.Remove(f.FilePath())
		}(file)
	}

	wg.Wait()
	if err != nil {
		return errors.New("something got wrong while removing files")
	}
	
	return nil
}
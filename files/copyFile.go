package files

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := createDir(dest); err != nil {
		return err
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	
	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}

	if err := destFile.Sync(); err != nil {
		return err
	}

	return nil
}

func createDir(path string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0766); err != nil {
		return err
	}

	return nil
}


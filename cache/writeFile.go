package cache

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func (c *cache) updateCache(file *os.File, tempFile *os.File) {
	currentTimeWithLineBreak := fmt.Sprintln(time.Now().Format(timestampLayout))
	if _, err := tempFile.WriteString(currentTimeWithLineBreak); err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if _, err := tempFile.WriteString(fmt.Sprintln(sc.Text())); err != nil {
			panic(err)
		}
	}
}

func (c *cache) flush(tempFile *os.File) {
	if err := os.Rename(tempFile.Name(), c.cacheFilePath); err != nil {
		panic(err)
	}
	tempFile.Sync()
}

func (c *cache) rollback(tempFile *os.File, updateError *error) {
	if err := recover(); err != nil {
		tempFile.Close()
		os.Remove(tempFilePath)
		*updateError = fmt.Errorf("error occured updating... \nRollback Changes")
	}
}
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"image-replacement-automater/cache"
)

func main() {
	// 0. Escape if go program is run at an improper directory
		// - Check if current program is called in a directory that has ./package.json
	currentPath := getCurrentPath()
	if err := isValidDirectory(currentPath); err != nil {
		fmt.Println(err)
		return 
	}
	
	// 1. Check "Last Runtime" from cache
		// if ./cache exists && ./cache/go-image-replacement-automater exists >>> get last time
		// elif ./cache not exist >>> mkdir ./cache directory and touch ./cache/go-image-replacement-automator file 
		// elif ./cache exists but ./cache/go-image-replacement-automater does not exists >> touch ./cache/go-image-replacement-automator file
			// both elif returns the zero time. 
			// all cases should record current time in the cache file.
			// mkdir > touch > writefile process should be run concurrently (in another goroutine) ; but no need to be waited
	cache := cache.Cache()
	lastRuntime := cache.ReadCache()
	fmt.Println(lastRuntime)

	// 2. Recursively probe under the given document directory <which could be served with a flag, but has a default value of ./document>
		// only deal with the file 1) .md or .mdx extension 2) modified after last run cached
		// this probing & 2-1 processes should be run concurrently at a seperate goroutine.
	modifiedChan := make(chan os.FileInfo, 100)
	probeDirectory := ""
	var isTarget func () bool
	
	go func() {
		defer close(modifiedChan)
		
		filepath.Walk(probeDirectory, func (path string, info os.FileInfo, err error) error {
			if err != nil { return err }	

			if isTarget() {
				modifiedChan <- info
			}
			return nil
		})
	}()
	
	// 2-1. Deal with the target files >>> grep !()[] or <img /> syntax
		// if image is locally sourced >>> copy image files under ./imgs directory
			// if ./imgs directory does not exist >>> mkdir 
			// if copy failed (for some reason. maybe the image path is unavailable or sth.) >>> log Error, but DO NOT END the go process
		// Overwrite image route to relative path (./imgs/{}.png kinda sth) 
		// if image local source 
	type parsedInfo struct {
		fullSyntax string
		imageSource string
		description string
	}
	var parseImageInfo func (line string) (*parsedInfo, error)
	var reconstructSyntax func (parsed *parsedInfo) string
	var wg sync.WaitGroup
	for fileInfo := range modifiedChan {
		go func() {
			file, err := os.Open(fileInfo.Name())
			if err != nil {
				return 
			}
			defer file.Close()
			
			var updatedStrings []string = []string{}
			sc := bufio.NewScanner(file)
			
			for sc.Scan() {
				lineText := sc.Text()
				parsed, err := parseImageInfo(lineText)
				// does contain image syntax
				if err != nil {
					reconstructed := reconstructSyntax(parsed)
					lineText = strings.Replace(lineText, parsed.fullSyntax, reconstructed, -1)
					go func () {
						// copy image from absolute path > ./imgs
					}()
				}
				updatedStrings = append(updatedStrings, lineText)
			}

			if err := sc.Err(); err != nil {
				panic(err)
			}
			updatedContent := strings.Join(updatedStrings, "\n")
			if err := os.WriteFile(fileInfo.Name(), []byte(updatedContent), 0644); err != nil {
				panic(err)
			}
		}()
		
	}

	wg.Wait()
	// 3. Wait for all files to be modified & print the result 
		// -) total run time for this run
		// a) total files traversed
		// b) modified files after last run
		// c) modified files during this run
		// d) modified counts during this run
		// e) error occured during this run

	cache.WriteCache()
}

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {		
		panic(err)
	}
	return path
}

func isValidDirectory(currentPath string) error {
	packageJsonPath := filepath.Join(currentPath, "package.json")
	if _, err := os.Stat(packageJsonPath); os.IsNotExist(err) {
		return errors.New("not a proper directory for node blog... package.json not found")
	}
	
	return nil
}
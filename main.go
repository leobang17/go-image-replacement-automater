package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	
	// 2. Recursively probe under the given document directory <which could be served with a flag, but has a default value of ./document>
		// only deal with the file 1) .md or .mdx extension 2) modified after last run cached
		// this probing & 2-1 processes should be run concurrently at a seperate goroutine.

	// 2-1. Deal with the target files >>> grep !()[] or <img /> syntax
		// if image is locally sourced >>> copy image files under ./imgs directory
			// if ./imgs directory does not exist >>> mkdir 
			// if copy failed (for some reason. maybe the image path is unavailable or sth.) >>> log Error, but DO NOT END the go process
		// Overwrite image route to relative path (./imgs/{}.png kinda sth) 
		// if image local source 
	
	// 3. Wait for all files to be modified & print the result 
		// -) total run time for this run
		// a) total files traversed
		// b) modified files after last run
		// c) modified files during this run
		// d) modified counts during this run
		// e) error occured during this run
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
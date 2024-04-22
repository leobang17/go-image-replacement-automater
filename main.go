package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	atomicity "image-replacement-automater/atomicity"
	cache "image-replacement-automater/cache"
	files "image-replacement-automater/files"
	imagetag "image-replacement-automater/imageTag"
	"image-replacement-automater/images"
)

func main() {

	// 0. Escape if go program is run at an improper directory
		// - Check if current program is called in a directory that has ./package.json
	// currentPath := files.GetCurrentPath()
	// if err := files.IsValidDirectory(currentPath); err != nil {
	// 	fmt.Println(err)
	// 	return 
	// }
	
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
	modifiedChan := files.FileModified(lastRuntime)

	// 2-1. Deal with the target files >>> grep !()[] or <img /> syntax
		// if image is locally sourced >>> copy image files under ./imgs directory
			// if ./imgs directory does not exist >>> mkdir 
			// if copy failed (for some reason. maybe the image path is unavailable or sth.) >>> log Error, but DO NOT END the go process
		// Overwrite image route to relative path (./imgs/{}.png kinda sth) 
		// if image local source 

	var wg sync.WaitGroup
	for meta := range modifiedChan {
		wg.Add(1)
		go func(meta files.FileMeta) {
			defer wg.Done()

			file, err := os.Open(meta.Path)
			if err != nil {
				return 
			}
			defer file.Close()
			
			type log struct {
				OriginalSource string
				CopiedSource string
				FileName string
				LineNumber int
			}
			atomicSection := atomicity.NewAtomicSection[*log]()
			var updatedStrings []string = []string{}
			sc := bufio.NewScanner(file)
			changeCount := 0
			resultLogs, atomicErr := atomicSection(func (runnable atomicity.Runnable[*log]) {
				lineCount := 1
				for sc.Scan() {
					lineText := sc.Text()
					imageTags, _ := imagetag.ParseImageTag(lineText)
					// iterate if contain image syntax
					for _, itag := range imageTags {
						switch images.ResolveSrcType(itag.Source) {
						case images.Web: 
							// do web thing...
							
						// copy image from absolute path > ./imgs	
						case images.Absolute:
							changeCount += 1
							newTag := itag.ConstructRelativeTag()
							lineText = strings.Replace(lineText, itag.FullTag, newTag.FullTag, -1)
							runnable(func () (*log, error) {
								if err := files.CopyFile(itag.Source, newTag.Source); err != nil {
									return nil, err
								}
								return &log {
									FileName: meta.Path,
									LineNumber: lineCount,
									OriginalSource: itag.Source,
									CopiedSource: newTag.Source,
								}, nil
							})
						case images.Relative:
							// do nothing if relative...
						}	
						
						updatedStrings = append(updatedStrings, lineText)
						lineCount += 1
					}
				}
				updatedContent := strings.Join(updatedStrings, "\n")
				
				if err := sc.Err(); err != nil {
					panic(err)
				}

				if changeCount > 0 {
					if err := os.WriteFile(meta.Path, []byte(updatedContent), 0644); err != nil {
						panic(err)
					}
				}
			})


		}(meta)
		
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

package files

import "path/filepath"

type extChecker struct {
	extMap map[string] bool
}

func NewExtChecker(exts ...string) *extChecker {
	newChecker := extChecker{
		extMap: make(map[string]bool),
	}
	for _, ext := range exts {
		newChecker.extMap[ext] = true
	}

	return &newChecker
}

func (e *extChecker) Match(ext string) bool {
	return e.extMap[ext]
}

func (e *extChecker) MatchWithFilepath(path string) bool {
	fileExt := filepath.Ext(path)
	return e.Match(fileExt)
}
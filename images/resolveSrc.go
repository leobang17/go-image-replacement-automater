package images

import (
	"net/url"
	"path/filepath"
)


type PathType int

const (
	Web PathType = iota
	Absolute          
	Relative           
	None
)

func ResolveSrcType(path string) PathType {
	if isWebURL(path) {
		return Web
	}
	
	if filepath.IsAbs(path) {
		return Absolute
	}

	if filepath.IsLocal(path) {
		return Relative
	}

	return None
}

func isWebURL(path string) bool {
	u, err := url.Parse(path)
	return err == nil && u.Scheme != "" && u.Host != ""
}
package imagetag

import (
	"errors"
	"regexp"
)


type markdownParser struct {
	regex *regexp.Regexp
}

func newMarkdownParser() *markdownParser {
	mp := &markdownParser{
		regex: regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`),
	}
	return mp
}

func (mp *markdownParser) Parse(line string) (*ImageTag, error) {
	var matches []string
	if matches = mp.regex.FindStringSubmatch(line); matches == nil || len(matches) < 3 {
		return nil, errors.New("not an markdown image tag")	
	} 

	imageTag := &ImageTag{
		FullTag: matches[0],
		Description: matches[1],
		Source: matches[2],
	}

	return imageTag, nil
}
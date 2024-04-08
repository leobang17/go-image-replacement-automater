package imagetag

import (
	"errors"
	"regexp"
)

type htmlParser struct {
	fullRegex *regexp.Regexp
	srcRegex *regexp.Regexp
	altRegex *regexp.Regexp
}

func newHtmlParser() *htmlParser {
	hp := &htmlParser {
		fullRegex: regexp.MustCompile(`<img\s+[^>]*?>`),
		srcRegex: regexp.MustCompile(`src="(.*?)"`),
		altRegex: regexp.MustCompile(`alt="(.*?)"`),			
	}
	return hp
}

func (hp *htmlParser) Parse(line string) (*ImageTag, error) {
	imageTag := new(ImageTag)
	fullTag := hp.fullRegex.FindString(line)
	if fullTag == "" {
		return nil, errors.New("not an image tag")
	}
	imageTag.FullTag = fullTag
	
	srcMatch := hp.srcRegex.FindStringSubmatch(fullTag)
	if srcMatch != nil {
		imageTag.Source = srcMatch[1]
	}

	altMatch := hp.altRegex.FindStringSubmatch(fullTag)
	if altMatch != nil {
		imageTag.Description = altMatch[1]
	}

	return imageTag, nil
}
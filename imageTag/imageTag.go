package imagetag

import (
	"errors"
	"fmt"
)

type ImageTag struct {
	FullTag string
	Description string
	Source string
}

type ImageParser interface {
	Parse(line string) (*ImageTag, error)
}

func ParseImageTag(line string) ([]ImageTag, error) {
	parsedTags := []ImageTag{}
	mp := newMarkdownParser()
	hp := newHtmlParser()

	mresult, merr := mp.Parse(line)
	if merr == nil {
		parsedTags = append(parsedTags, *mresult)
	}

	hresult, herr := hp.Parse(line)
	if herr == nil {
		parsedTags = append(parsedTags, *hresult)
	}
	
	if merr != nil && herr != nil {
		return nil, errors.New("not an image tag")
	}

	return parsedTags, nil
}


func Reconstruct(tag *ImageTag) string {
	return fmt.Sprintf("![%s](%s)", tag.Description, tag.Source)
}
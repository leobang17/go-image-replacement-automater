package imagetag

import (
	"errors"
	"fmt"
	"path/filepath"
)

type ImageTag struct {
	FullTag string
	Description string
	Source string
}


func (t ImageTag) ConstructRelativeTag() ImageTag {
	relativeSrc := filepath.Join("./", "imgs", filepath.Base(t.Source))
	newTag := ImageTag {
		Description: t.Description,
		Source: relativeSrc,
	} 
	
	newTag.reconstruct()
	return newTag
}
	
func (t *ImageTag) reconstruct() {
	t.FullTag = fmt.Sprintf("![%s](%s)", t.Description, t.Source)
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
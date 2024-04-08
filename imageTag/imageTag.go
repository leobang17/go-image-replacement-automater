package imagetag

import "fmt"

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
	return parsedTags, nil
}


func Reconstruct(tag *ImageTag) string {
	return fmt.Sprintf("![%s](%s)", tag.Description, tag.Source)
}
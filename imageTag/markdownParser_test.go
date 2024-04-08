package imagetag

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)


func Test_MarkdownImageParser(t *testing.T) {
	markdownParser := newMarkdownParser()
	source := `/User/leobang17/go-image/test.png`
	description := `this is description`
	
	t.Run("Success: src, alt both exists", func (t *testing.T)  {
		fullTag := fmt.Sprintf(`![%s](%s)`, description, source)
		rawLine := fmt.Sprintf(`blah blah %s blah blah`, fullTag)
		
		expected := ImageTag{
			FullTag:  fullTag,
			Source: source,
			Description: description,
		}

		imageTag, err := markdownParser.Parse(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		} 
	
		if diff := cmp.Diff(*imageTag, expected); diff != "" {
			t.Error("Parsed result not correct ", diff)
		}
	})

	t.Run("Success: src only exists", func (t *testing.T)  {
		fullTag := fmt.Sprintf(`![](%s)`, source)
		rawLine := fmt.Sprintf(`blah blah %s blah blah`, fullTag)
		
		expected := ImageTag{
			FullTag:  fullTag,
			Source: source,
			Description: "",
		}

		imageTag, err := markdownParser.Parse(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		} 
	
		if diff := cmp.Diff(*imageTag, expected); diff != "" {
			t.Error("Parsed result not correct \n", diff)
		}
	})
	
	t.Run("Failed: not an image tag", func (t *testing.T) {
		_, err := markdownParser.Parse("fwekfjwenfjknewjknfwekfn")
		if err == nil {
			t.Error("Should return error")
		}
	})

	t.Run("Failed: weired syntax", func (t *testing.T) {
		rawLine := fmt.Sprintf(`![description fwe%s)`, source)

		_, err := markdownParser.Parse(rawLine)
		if err == nil {
			t.Error("should return error")
		}
	})
}
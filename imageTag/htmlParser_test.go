package imagetag

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_HtmlImageParser(t *testing.T) {
	htmlParser := newHtmlParser()
	source := `/User/leobang17/go-image/test.png`
	description := `this is description`
	t.Run("Success: img, alt both exists. sequence img > alt", func (t *testing.T)  {
		rawLine := fmt.Sprintf(`blah blah <img src="%s" alt="%s" /> blah blah`, source, description)
		
		expected := ImageTag{
			FullTag:  `<img src="/User/leobang17/go-image/test.png" alt="this is description" />`,
			Source: source,
			Description: description,
		}

		imageTag, err := htmlParser.Parse(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		} 
	
		if diff := cmp.Diff(*imageTag, expected); diff != "" {
			t.Error("Parsed result not correct ", diff)
		}
	})

	t.Run("Success: img, alt both exists. sequence alt > img", func (t *testing.T)  {
		fullTag := fmt.Sprintf(`<img alt="%s" src="%s" fwefew />`, description, source)
		rawLine := fmt.Sprintf(`blah blah %s blah blah`, fullTag)
		
		expected := ImageTag{
			FullTag:  fullTag,
			Source: source,
			Description: description,
		}

		imageTag, err := htmlParser.Parse(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		} 
	
		if diff := cmp.Diff(*imageTag, expected); diff != "" {
			t.Error("Parsed result not correct ", diff)
		}
	})
	
	t.Run("Success: src only exists", func (t *testing.T) {
		fullTag := fmt.Sprintf(`<img src="%s" fwefew />`, source)
		rawLine := fmt.Sprintf(`blah blah %s blah blah`, fullTag)
		
		expected := ImageTag{
			FullTag:  fullTag,
			Source: source,
			Description: "",
		}
	
		imageTag, err := htmlParser.Parse(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		} 
	
		if diff := cmp.Diff(*imageTag, expected); diff != "" {
			t.Error("Parsed result not correct ", diff)
		}
	})

	t.Run("Failed: not an image tag", func (t *testing.T) {
		_, err := htmlParser.Parse("fwekfjwenfjknewjknfwekfn")
		if err == nil {
			t.Error("Should return error")
		}
	})
}
package imagetag

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ParseImageTag_with_markdown_syntax(t *testing.T) {
	source := `/User/leobang17/go-image/test.png`
	description := `this is description`

	t.Run("Success: src, alt both exists", func (t *testing.T)  {
		fullTag := fmt.Sprintf(`![%s](%s)`, description, source)
		rawLine := fmt.Sprintf(`blah blah %s blah blah`, fullTag)

		expected := []ImageTag{{
			FullTag:  fullTag,
			Source: source,
			Description: description,
		}}

		imageTag, err := ParseImageTag(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		}

		if diff := cmp.Diff(imageTag, expected); diff != "" {
			t.Error("Parsed result not correct ", diff)
		}
	})

	t.Run("Success: src only exists", func (t *testing.T)  {
		fullTag := fmt.Sprintf(`![](%s)`, source)
		rawLine := fmt.Sprintf(`blah blah %s blah blah`, fullTag)

		expected := []ImageTag{{
			FullTag:  fullTag,
			Source: source,
			Description: "",
		}}

		imageTag, err := ParseImageTag(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		}

		if diff := cmp.Diff(imageTag, expected); diff != "" {
			t.Error("Parsed result not correct \n", diff)
		}
	})

	t.Run("Failed: not an image tag", func (t *testing.T) {
		_, err := ParseImageTag("fwekfjwenfjknewjknfwekfn")
		if err == nil {
			t.Error("Should return error")
		}
	})

	t.Run("Failed: weired syntax", func (t *testing.T) {
		rawLine := fmt.Sprintf(`![description fwe%s)`, source)

		_, err := ParseImageTag(rawLine)
		if err == nil {
			t.Error("should return error")
		}
	})
}

func Test_ParseImageTag_with_html_syntax(t *testing.T) {
	source := `/User/leobang17/go-image/test.png`
	description := `this is description`
	t.Run("Success: img, alt both exists. sequence img > alt", func (t *testing.T)  {
		rawLine := fmt.Sprintf(`blah blah <img src="%s" alt="%s" /> blah blah`, source, description)

		expected := []ImageTag{{
			FullTag:  `<img src="/User/leobang17/go-image/test.png" alt="this is description" />`,
			Source: source,
			Description: description,
		}}

		imageTag, err := ParseImageTag(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		}

		if diff := cmp.Diff(imageTag, expected); diff != "" {
			t.Error("Parsed result not correct ", diff)
		}
	})

	t.Run("Success: img, alt both exists. sequence alt > img", func (t *testing.T)  {
		fullTag := fmt.Sprintf(`<img alt="%s" src="%s" fwefew />`, description, source)
		rawLine := fmt.Sprintf(`blah blah %s blah blah`, fullTag)

		expected := []ImageTag{{
			FullTag: fullTag,
			Source: source,
			Description: description,
		}}

		imageTag, err := ParseImageTag(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		}

		if diff := cmp.Diff(imageTag, expected); diff != "" {
			t.Error("Parsed result not correct ", diff)
		}
	})

	t.Run("Success: src only exists", func (t *testing.T) {
		fullTag := fmt.Sprintf(`<img src="%s" fwefew />`, source)
		rawLine := fmt.Sprintf(`blah blah %s blah blah`, fullTag)

		expected := []ImageTag{{
			FullTag:  fullTag,
			Source: source,
			Description: "",
		}}

		imageTag, err := ParseImageTag(rawLine)
		if err != nil {
			t.Error("Parse() should not return error")
		}

		if diff := cmp.Diff(imageTag, expected); diff != "" {
			t.Error("Parsed result not correct ", diff)
		}
	})

	t.Run("Failed: not an image tag", func (t *testing.T) {
		_, err := ParseImageTag("fwekfjwenfjknewjknfwekfn")
		if err == nil {
			t.Error("Should return error")
		}
	})
}

func Test_ImageTag_ConstructRelativeTag(t *testing.T) {    
	originTag := ImageTag{
			Description: "Image of me",
			Source:      "/Users/leobang/selfie.png",
			FullTag: 		 `<img alt="Image of me" src="/Users/leobang/selfie.png"`,
	}

	expected := ImageTag{
		Description: "Image of me",
		Source:      "imgs/selfie.png",
		FullTag: 		 `![Image of me](imgs/selfie.png)`,
	}
	
	actual := originTag.ConstructRelativeTag()

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Error(diff)
	}
}
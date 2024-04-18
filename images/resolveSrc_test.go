package images

import (
	"cmp"
	"testing"
)

func Test_resolveSrc(t *testing.T) {
	tests := []struct {
		name 		 string
		path     string
		expected PathType
	}{
		{"HTTP Url", "http://leobang.me/image.png", Web},
		{"HTTPS Url",  "https://leobang.me/image.png", Web},
		{"Absolute Path",  "/Users/leobang/images/image.png", Absolute},
		{"Relative Path", "./imgs/image.png", Relative},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := ResolveSrcType(test.path)
			if diff := cmp.Compare(test.expected, result); diff != 0 {
				t.Error(diff)
			}
		})
	}
}
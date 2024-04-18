package files

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_copyfile_success(t *testing.T) {
	
	src := "sourceFile.txt"
	srcContent := []byte("this is original text")
	srcFile, err := os.CreateTemp(".", src)
	if err != nil {
			t.Error("error while creating temp src file", err)
	}
	defer srcFile.Close()

	t.Cleanup(func() {
		os.Remove(srcFile.Name()) // 테스트 후 원본 파일 삭제
	})

	_, err = srcFile.Write(srcContent)
	if err != nil {
			t.Fatal(err)
	}
	srcFile.Close()

	dest := "./imgs/destFile.txt"
	if err := CopyFile(srcFile.Name(), dest); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.RemoveAll(dest)
		os.Remove("./imgs")
	})

	result, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(srcContent), string(result)); diff != "" {
		t.Error(diff)
	}
}

func Test_copyfile_fail(t *testing.T) {
	src := "sourceFile.txt"

	dest := "./imgs/destFile.txt"
	if err := CopyFile(src, dest); err == nil {
		t.Error("error should occur if there is no src file.")
	}	
}
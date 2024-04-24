package files

import (
	"os"
	"testing"
)

// 테스트에 사용할 파일 경로와 이름
var testPaths = []string { 
	"testfile1.txt",
	"testfile2.txt",
	"testfile3.txt",
	"testfile4.txt",
}

func TestMain(m *testing.M) {
		setup()
    code := m.Run()
    // teardown()
    os.Exit(code)
}

func setup() {
	for _, p := range testPaths {
		file, err := os.Create(p)
		if err != nil {
			 panic(err)
		}
		file.Close()
	}
}

func teardown() {
		for _, p := range testPaths {
			os.Remove(p)
		}
}

func Test_RemoveFiles_success(t *testing.T) {
    logs := make([]*log, len(testPaths))
		for i, path := range testPaths {
			logs[i] = &log{CopiedSource: path} 
		}	

    err := RemoveFiles(logs)
    if err != nil {
        t.Error(err)
    }

		for _, path := range testPaths {
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				t.Errorf("file should be all deleted but %s is still remaining", path)
			}
		}
}

type log struct {
	OriginalSource string
	CopiedSource string
	FileName string
	LineNumber int
}

func (l *log) FilePath() string {
	return l.CopiedSource
}
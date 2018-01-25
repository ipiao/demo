package basePkgT

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestIOUtil(t *testing.T) {
	var reader = strings.NewReader("dsasgddfsgdf")
	//var nonpcloser = ioutil.NopCloser(reader)
	bs, err := ioutil.ReadAll(reader)
	t.Log(string(bs), err)

	fileinfos, err := ioutil.ReadDir("../..")
	t.Log(err)
	for _, file := range fileinfos {
		t.Log(file)
	}

	bbs, err := ioutil.ReadFile("../main.go")
	t.Log(string(bbs), err)

}

func TestIOTempDirFile(t *testing.T) {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("./", "example")
	t.Log(dir)
	if err != nil {
		t.Fatal(err)
	}

	//	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		t.Fatal(err)
	}
}

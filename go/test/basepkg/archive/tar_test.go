package basePkgT

import (
	"archive/tar"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestTar(t *testing.T) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)
	// Create a new tar archive.
	tw := tar.NewWriter(buf)
	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			t.Fatal(err)
		}
	}
	// Make sure to check the error on Close.
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	// Open the tar archive for reading.
	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)
	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			t.Fatal(err)
		}
		t.Log()
	}
}

func TestTar2(t *testing.T) {
	fw, err := os.Create("test.tar")
	if err != nil {
		t.Fatal(err)
	}
	defer fw.Close()

	tw := tar.NewWriter(fw)
	defer tw.Close()

	files, err := ioutil.ReadDir("../../resources/")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			return
		}
		body := readFile("../../resources/" + file.Name())

		hdr := &tar.Header{
			Name: file.Name(),
			Mode: 0600,
			Size: int64(len(body)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			t.Logf("%v", err)
			return
		}
		if _, err := tw.Write([]byte(body)); err != nil {
			t.Logf("%v", err)
			return
		}
	}
	t.Log("done")
}

func readFile(path string) string {
	var content string
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%v", err)
		return ""
	}

	defer file.Close()

	inputReader := bufio.NewReader(file)
	for {
		inputString, err := inputReader.ReadString('\n')
		content += inputString
		if err == io.EOF {
			break
		}
	}

	return content
}

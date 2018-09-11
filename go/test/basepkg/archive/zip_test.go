package basePkgT

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	// zip.RegisterCompressor(1, func(w io.Writer) (io.WriteCloser, error) {
	// 	return gzip.NewWriter(w), nil
	// })

	// zip.RegisterDecompressor(1, func(r io.Reader) io.ReadCloser {
	// 	gr, _ := gzip.NewReader(r)
	// 	return gr
	// })

	f, _ := os.Create("files1.gz")
	defer f.Close()
	zw := zip.NewWriter(f)
	defer zw.Close()

	files, _ := ioutil.ReadDir("files")
	for _, file := range files {
		zh, err := zip.FileInfoHeader(file)
		if err != nil {
			t.Fatal(err)
		}
		zh.Method = zip.Deflate
		zh.Name = "files/" + file.Name()
		_, err = zw.CreateHeader(zh)
		if err != nil {
			t.Fatal(err)
		}
		// io.Copy(hw, zh)

		zf, err := zw.Create(zh.Name)
		if err != nil {
			t.Fatal(err)
		}

		zzf, _ := os.Open(zh.Name)
		defer zzf.Close()
		io.Copy(zf, zzf)

	}
}

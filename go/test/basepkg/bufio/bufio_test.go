package basePkgT

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestBufferRead(t *testing.T) {
	f, _ := os.Open("2.txt")
	defer f.Close()
	br := bufio.NewReaderSize(f, 15)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				t.Log(err)
				break
			}
			t.Fatal(err)
		}
		t.Log(string(line))
	}
}

func TestBufferWrite(t *testing.T) {
	f, _ := os.Create("2.txt")
	defer f.Close()
	bw := bufio.NewWriterSize(f, 15)
	defer bw.Flush()

	_, err := bw.WriteString("这是第一行hahah\n")
	if err != nil {
		t.Fatal(err)
	}
	_, err = bw.WriteString("这是第二行\n")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(bw.Available())
	t.Log(bw.Buffered())
}

func TestBufferReadWrite(t *testing.T) {
	f, _ := os.Open("1.txt")
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	rw := bufio.NewReadWriter(r, w)

	done := make(chan struct{})
	go func() {
		for {
			line, _, err := rw.ReadLine()
			if err != nil {
				t.Log(err)
				done <- struct{}{}
			}
			t.Log("Got:", string(line))
			n, err := rw.Write(line)
			rw.WriteByte('\n')
			rw.Flush()
			t.Log("Write:", n, err)
		}
	}()
	<-done
}

func TestBufferScan(t *testing.T) {
	f, _ := os.Open("1.txt")
	defer f.Close()

	bs := bufio.NewScanner(f)
	bs.Split(bufio.ScanWords)

	for bs.Scan() {
		t.Log(bs.Text())
	}
}

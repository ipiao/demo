package basePkgT

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestOs(t *testing.T) {
	err := os.Chdir(".") // cd
	t.Log(err)
	err = os.Chmod("test.sh", 0777) //chmod
	t.Log(err)
	cmd := exec.Command("sh", "test.sh") // shell
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	t.Log(err)
	t.Logf("output is %s", out.String())

	os.Chtimes("test.sh", time.Now().Add(time.Second*6), time.Now().Add(time.Second*5)) // 修改文件的访问时间和修改时间
	// var envs = os.Environ()
	// t.Log("envs:", envs)
	//os.Clearenv() // 删除所有的环境变量
	executable, err := os.Executable()
	t.Log("executable:", executable, err)

	// os.Exit(0)
	goroot := os.ExpandEnv("hello $GOROOT")
	t.Log(goroot)
	gor := os.Getenv("PWD")
	t.Log(gor)

	ss := os.Expand("hello $name", func(s string) string {
		if s == "name" {
			return "tom"
		}
		return ""
	})
	t.Log(ss)

	uid := os.Getuid()
	t.Log(uid)
	gid := os.Getgid()
	t.Log(gid)
	groups, err := os.Getgroups()
	t.Log(groups, err)
	psgesize := os.Getpagesize()
	t.Log(psgesize)
	pid := os.Getpid()
	t.Log(pid)
	ppid := os.Getpid()
	t.Log(ppid)

	pwd, err := os.Getwd()
	t.Log(pwd)
	hostname, err := os.Hostname()
	t.Log(hostname, err)

	filename := "a-nonexistent-file"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Log("file does not exist")
	}

	// err = os.Link("test.sh", "newtest.sh")
	// t.Log(err)
	// err = os.Remove("test.sh")
	// t.Log(err)
	// err = os.Rename("newtest.sh", "test.sh")
	// t.Log(err)

	tempDir := os.TempDir()
	t.Log(tempDir)

	err = os.Truncate("../../resources/images/idcard_back.jpg", 1024*13)
	t.Log(err)
}

func TestOsFile(t *testing.T) {
	file, err := os.Create("test.txt") // 0666 except write and create
	defer file.Close()
	t.Log(file, err)
	file.Write([]byte("hello"))
	file.Close()

	fd := file.Fd()
	t.Log(fd)
	newFile := os.NewFile(fd, "newtest.txt")
	_, err = newFile.WriteString("adsgffdgdfs")
	defer newFile.Close()
	t.Log(err)
}

func TestPipe(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Logf("Error: Couldn't create the named pipe: %s\n", err)
	}
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			t.Logf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		t.Logf("Read %d byte(s). [file-based pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		t.Logf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	t.Logf("Written %d byte(s). [file-based pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}

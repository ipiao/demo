package basePkgT

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestExecBase(t *testing.T) {
	var err error
	//res, err := exec.LookPath("../basepkg")
	//t.Log(res, err)
	err = os.Chdir("..") // cd ..
	t.Log(err)
	//err = os.Chmod("basepkg", os.ModePerm)
	t.Log(err)
	cmd := exec.Command("./basepkg")
	t.Log("path:", cmd.Path)
	t.Log("args:", cmd.Args)
	t.Log("env:", cmd.Env)
	t.Log("dir:", cmd.Dir)
	//cmd.Stdin = strings.NewReader("./basepkg")
	var out bytes.Buffer
	var er bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &er
	err = cmd.Run()
	t.Log(err)
	t.Log("stdin:", cmd.Stdin)
	t.Log("stdout:", out.String())
	t.Log("stderr:", er.String())
	t.Log("ExtraFiles:", cmd.ExtraFiles)
	t.Log("SysProcAttr:", cmd.SysProcAttr)
	t.Log("Process:", cmd.Process)
	t.Log("ProcessState:", cmd.ProcessState)
}

func TestExecCommandContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		t.Log(err)
	}
}

func TestExecCombinedOutput(t *testing.T) {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", stdoutStderr)
}

func TestExecOutput(t *testing.T) {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("The date is %s\n", out)
}

func TestExecStartWait(t *testing.T) {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	t.Log("Waiting for command to finish...")
	err = cmd.Wait()
	t.Logf("Command finished with error: %v", err)
}

func TestExecStderrPipe(t *testing.T) {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	t.Logf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func TestExecStdoutPipe(t *testing.T) {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

func TestExecStdoinPipe(t *testing.T) {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

func TestStdOutReadEOF(t *testing.T) {
	cmd := exec.Command("echo", "-n", "My first commmand comes from golang")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("Error:cound not obtain the stdout:%s\n", err)
	}

	if err := cmd.Start(); err != nil {
		t.Fatalf("Error:command can not startup:%s\n", err)
	}

	// var outBuff bytes.Buffer
	// for {
	// 	tempOutput := make([]byte, 5)
	// 	n, err := stdout.Read(tempOutput)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			t.Log("EOF")
	// 			break
	// 		} else {
	// 			t.Fatalf("Error: cound not read from pipe:%s\n", err)
	// 		}
	// 	}
	// 	if n > 0 {
	// 		outBuff.Write(tempOutput[:n])
	// 	}
	// }
	outBuff := bufio.NewReader(stdout)
	output, _, err := outBuff.ReadLine()
	t.Log(string(output))
}

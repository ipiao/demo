package basePkgT

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
)

func TestSignal(t *testing.T) {
	sigcha1 := make(chan os.Signal, 1)
	sigcha2 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	sigs2 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	// sigs3 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigcha1, sigs1...)
	signal.Notify(sigcha2, sigs2...)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {

		for sig := range sigcha1 {
			fmt.Println("11111", sig)
			// os.Exit(0)
			close(sigcha1)
		}

		wg.Done()
	}()

	go func() {

		for sig := range sigcha2 {
			fmt.Println("2222222", sig)
			//os.Exit(0)
			close(sigcha2)
		}

		wg.Done()
	}()

	wg.Wait()
}

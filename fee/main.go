package main

import (
	"fmt"
	"github.com/smapig/go-ddd-sample/fee/ioc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server, err := ioc.InitializeServer("")
	if err != nil {
		panic(fmt.Errorf("server initialization failed: %w \n", err))
	}

	stopSignal := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		stopSignal <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		if err := server.StartHTTP(); err != nil {
			panic(fmt.Errorf("server HTTP starting failed: %w \n", err))
		}
	}()

	// Wait for stopping signal.
	err = <-stopSignal
	log.Printf("server stopped %s", err.Error())
}
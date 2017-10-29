package main

import (
	"math/rand"
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan bool, 1)
	go ProcessSignal(sig)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := r.Intn(40) + 10;
	time.Sleep((time.Second * time.Duration(randNum)))
	if randNum > 20 {
		panic("Too much!")
	}
	fmt.Println("Done")
}

func ProcessSignal(sig chan bool) {
	sigch := make(chan os.Signal)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	for {
		signalType := <-sigch
		fmt.Println("Received signal from channel : ", signalType)

		switch signalType {
		case syscall.SIGINT, syscall.SIGTERM:
			fmt.Println("got Termination signal - portable number 15")
			time.Sleep(time.Second * 3)
			sig <- true
			os.Exit(0)
			break
		}
	}
}
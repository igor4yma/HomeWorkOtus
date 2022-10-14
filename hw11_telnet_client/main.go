package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	timeout          time.Duration
	ErrNotEnoughArgs = errors.New("not enough arguments, should be 3 at least")
	sigs             = make(chan os.Signal, 1)
)

const (
	minLenArgs = 3
	maxLenArgs = 4
)

func init() {
	flag.DurationVar(&timeout, "timeout", 0, "connection timeout")
}

func main() {
	flag.Parse()
	if (len(os.Args) < minLenArgs) || (len(os.Args) > maxLenArgs) {
		log.Fatal(ErrNotEnoughArgs)
	}
	host := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]

	client := NewTelnetClient(
		net.JoinHostPort(host, port),
		timeout,
		os.Stdin,
		os.Stdout,
	)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	if err := client.Connect();err != nil {
	 log.Fatalln(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	go worker(client.Receive, cancelFunc)
	go worker(client.Send, cancelFunc)

	select {
	case <-sigs:
		cancelFunc()
		signal.Stop(sigs)
		return
	case <-ctx.Done():

		close(sigs)
		return
	}
}

func worker(handler func() error, cancelFunc context.CancelFunc) {
	if err := handler(); err != nil {
		cancelFunc()
	}
}

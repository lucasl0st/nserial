package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/lucasl0st/nserial/pkg/client"
	"github.com/lucasl0st/nserial/pkg/model"
)

var (
	port = flag.String("port", "/dev/ttyUSB0", "serial port")
	out  = flag.String("out", "out/", "output directory")
)

func main() {
	flag.Parse()

	c, err := client.New(client.WithPort(*port))
	if err != nil {
		panic(err)
	}

	interceptExit(c)

	err = c.Connect()
	if err != nil {
		panic(err)
	}

	data, err := c.GetData(consumer)
	if err != nil {
		panic(err)
	}

	err = writeData(data)
	if err != nil {
		panic(err)
	}
}

func writeData(b []byte) error {
	filename := path.Join(*out, "dump.bin")
	// nolint: gosec // this is user input
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	n, err := file.Write(b)
	if err != nil {
		return err
	}

	if n != len(b) {
		return errors.New("failed to write whole file")
	}

	return nil
}

func consumer(roll model.Roll) error {
	b, err := json.MarshalIndent(roll, "", "    ")
	if err != nil {
		return err
	}

	filename := path.Join(*out, fmt.Sprintf("roll-%d.json", roll.Number))
	// nolint: gosec // this is user input
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	n, err := file.Write(b)
	if err != nil {
		return err
	}

	if n != len(b) {
		return errors.New("failed to write whole file")
	}

	return nil
}

func interceptExit(client client.Client) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		err := client.Close()
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}()
}

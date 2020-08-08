package main

import (
	"fmt"
	"os"

	"github.com/ikascrew/mp4"
	"golang.org/x/xerrors"
)

func main() {

	f := "sample.mp4"

	err := run(f)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

}

func run(name string) error {

	fd, err := os.Open(name)
	if err != nil {
		return xerrors.Errorf("file open: %w", err)
	}
	defer fd.Close()

	_, err = mp4.Decode(fd)
	if err != nil {
		return xerrors.Errorf("file decode: %w", err)
	}

	return nil
}

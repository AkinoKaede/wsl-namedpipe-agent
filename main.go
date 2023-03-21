package main

import (
	"flag"
	"io"
	"os"

	"github.com/Microsoft/go-winio"
)

var (
	path = flag.String("path", "", "Path of the named pipe")
)

func main() {
	flag.Parse()
	conn, err := winio.DialPipe(*path, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	go func() {
		writePipe(io.Writer(conn))
	}()

	readPipe(io.Reader(conn))
}

func readPipe(pipe io.Reader) (int64, error) {
	return io.Copy(os.Stdout, pipe)
}

func writePipe(pipe io.Writer) (int64, error) {
	return io.Copy(pipe, os.Stdin)
}

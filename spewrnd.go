package main

import (
	"crypto/rand"
	"io"
	"os"
	"runtime"
)

func spewRandom(out io.Writer) error {
	for i := 0; i < 10; i++ {
		buf := make([]byte, 1*1024*1024)
		_, err := io.ReadFull(rand.Reader, buf)
		if err != nil {
			return err
		}

		_, err = out.Write(buf)
		if err != nil {
			return err
		}

		if os.Getenv("FORCEGC") == "1" {
			runtime.GC()
		}
	}

	return nil
}

func main() {
	err := spewRandom(os.Stdout)
	if err != nil {
		panic(err)
	}
}

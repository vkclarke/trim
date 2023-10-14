package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(bytes.TrimSpace(data))
}

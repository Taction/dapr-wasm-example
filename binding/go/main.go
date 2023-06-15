package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil && err != io.EOF {
		os.Stdout.WriteString(err.Error())
		return
	}

	os.Stdout.WriteString(strings.Join(os.Args[:1], " "))
	os.Stdout.Write(buf)
	os.Stdout.Write([]byte{'\n'})
}

package main

import (
	"fmt"
	"io"
	"os"
)

type cWriter struct {
	writer  io.Writer
	counter int64
}

func (c *cWriter) Write(p []byte) (n int, err error) {
	bytes, err := c.writer.Write(p)
	c.counter += int64(bytes)
	n = bytes
	//fmt.Printf("This Write() call has written %v bytes\n", n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var counter int64
	return w, &counter
}

func main() {
	mystrings := []string{"Hello I love you\n", "AAA\n", "13\n"}
	mycWriter := cWriter{}
	w, c := CountingWriter(os.Stdout)
	mycWriter.writer = w
	mycWriter.counter = *c

	for _, i := range mystrings {
		fmt.Printf("The string is: %v\n", i)
		fmt.Printf("Writing the string to the new writer...\n")

		_, err := mycWriter.Write([]byte(i))
		if err != nil {
			fmt.Printf("Can't write the string to stdout!\n")
			os.Exit(1)
		}

		fmt.Printf("Total nubmer of bytes written is: %v\n", mycWriter.counter)
	}
}

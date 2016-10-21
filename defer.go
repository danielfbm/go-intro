package main

import (
	"io"
	"os"
)

func main() {
	// defer is generally used to execute code
	// before a return statement or the end of
	// the code execution
	// Very useful to make your open and close
	// statements be very close
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

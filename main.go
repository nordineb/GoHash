package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

func sha512hash(value string) string {
	h := sha512.Sum512([]byte(ssn))
	return base64.StdEncoding.EncodeToString(h[:])
}

func main() {
	flag.Parse()

	filename := flag.Arg(0)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		var hash = sha512hash(string(line))
		fmt.Printf("%s - %s \n", line, hash)
	}
}

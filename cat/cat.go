package main

import (
	"io"
	"log"
	"os"
)

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, nil, err
	}

	// the getFile function returns a cleanup function, a common pattern in go
	return file, func() {
		file.Close()
	}, err
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing file to end of function
	defer closer()

	data := make([]byte, 2048)

	for {
		count, err := f.Read(data)

		os.Stdout.Write(data[:count])

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}
	}
}

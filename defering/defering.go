package main

import (
	"io"
	"log"
	"os"
)

// A common pattern in Go and many functional languages is to return a cleanup
// function when a function accesses an external resouce. This function
// returns a function that closes the file that it opens.
func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, nil, err
	}

	return file, func() {
		file.Close()
	}, err
}

// This function performs the same as the cat unix commnand. The thing to note
// is that the the closing of the file is scheduled for after the function
// completes using defer. You can run this with:
// go run ./defering/defering.go ./defering/file.txt
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

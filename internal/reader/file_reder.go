package reader

import (
	"fmt"
	"log"
	"os"
)

func Reader(file string) []byte {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return body
}

func Writer(file string, data []byte) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("unable to write file: %v", err)
	}

	n, err := f.Write(data)
	if err != nil {
		log.Fatalf("unable to write file: %v", err)
	}
	fmt.Printf("wrote %d bytes\n", n)
	defer f.Close()
	f.Sync()

}

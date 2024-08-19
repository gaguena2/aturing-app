package reader

import (
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

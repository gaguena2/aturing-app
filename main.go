package main

import (
	"fmt"

	"github.com/gaguena/aturing/internal/crypter"
	"github.com/gaguena/aturing/internal/reader"
)

func main() {
	provider := crypter.NewCrypter("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47")

	file := reader.Reader("./certificado")
	encoded := provider.Encrypt(file)
	fmt.Println(string(encoded))
	decrypt := provider.Decrypt(string(encoded))
	fmt.Println(string(decrypt))
}

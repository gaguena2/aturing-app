package main

import (
	"fmt"

	"github.com/gaguena/aturing/internal/crypter"
	"github.com/gaguena/aturing/internal/http"
	"github.com/gaguena/aturing/internal/reader"
)

func main() {
	secret, _ := http.GetSecret()
	provider := crypter.NewCrypter(*secret)

	file := reader.Reader("./certificado")
	encoded := provider.Encrypt(file)
	fmt.Println(string(encoded))
	decrypt := provider.Decrypt(string(encoded))
	fmt.Println(string(decrypt))
}

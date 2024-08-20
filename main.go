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

	//file := reader.Reader("./certificado")
	//encoded := provider.Encrypt(file)
	//reader.Writer("./encript", encoded)
	//decrypt := provider.Decrypt(string(encoded))
	file2 := reader.Reader("./encript")
	decrypt := provider.Decrypt(string(file2))
	fmt.Println(string(decrypt))
}

package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

func GetSecret() (*string, error) {
	config := Loader()

	client := &http.Client{}

	req, err := http.NewRequest("GET", config["HOST_SECRET"], nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil, err
	}
	req.Header.Add("X-Vault-Token", config["HOST_TOKEN"])
	req.Header.Add("Content-Type", "application/json")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil, err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return nil, err
	}
	field := gjson.ParseBytes(body).Get("data.data.secret").Str
	return &field, nil

}

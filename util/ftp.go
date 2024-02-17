package util

import (
	"io"
	"net/http"
	"os"
)

func GetOnlineResource(url string) []byte {
	// Grab the file from the url and return the contents
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := &http.Client{}

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	// Read the body of the response
	body, _ := io.ReadAll(resp.Body)

	return body
}

func StoreFile(content []byte, path string) {
	file, _ := os.Create(path)
	defer file.Close()
	// Write the body of the response to the file
	file.Write(content)
}

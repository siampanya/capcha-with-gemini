package service

import (
	"io/ioutil"
	"net/http"
	"os"
)

func GetFilesFromhttpImages(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body as a byte slice
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Create a temporary file to write the content
	tmpFile, err := os.CreateTemp("", "tmp")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	// Write the content to the temporary file
	_, err = tmpFile.Write(body)
	if err != nil {
		return nil, err
	}

	// Read the content from the temporary file
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, err
	}

	return content, nil

}

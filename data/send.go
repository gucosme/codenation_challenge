package data

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// SendData reads the contents off the given file and send it as a
// multipart-formdata request to the given url
func SendData(url string, file *os.File) ([]byte, error) {
	f, _ := os.Open(file.Name())
	defer f.Close()

	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("answer", filepath.Base(file.Name()))

	if err != nil {
		return nil, err
	}

	io.Copy(part, f)
	writer.Close()

	request, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return content, nil
}

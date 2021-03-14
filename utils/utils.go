package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func URLToFile(url, path string) error {
	body, err := BodyFromURL(url)
	if err != nil {
		return err
	}

	return WriteToFile(path, body)
}

func WriteToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		return err
	}

	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatal(err)
		return err
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func BodyFromURL(url string) (string, error) {
	bodyBuffer, err := BufferFromURL(url)
	if err != nil {
		return "", err
	}

	return string(bodyBuffer), nil
}

func BufferFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return make([]byte, 1), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// some random []byte
		return make([]byte, 1), err
	}

	return body, nil
}

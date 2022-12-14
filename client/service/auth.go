package service

import (
	"bytes"
	"net/http"
)

func Login(data []byte) (*http.Response, error) {
	url := "http://localhost:8000/api/login"
	result, error := http.Post(url, "application/json", bytes.NewBuffer(data))
	if error != nil {
		return nil, error
	}
	return result, nil
}

func Register(data []byte) (*http.Response, error) {
	url := "http://localhost:8000/api/register"
	result, error := http.Post(url, "application/json", bytes.NewBuffer(data))
	if error != nil {
		return nil, error
	}
	return result, error
}

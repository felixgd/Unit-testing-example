package service

import (
	"errors"
)

type response struct {
	text string
}

func NewService(text string) ServiceInterface {
	return &response{text}
}

func (r *response) ReturnText(s string) (string, error) {
	if len(s) > 10 {
		return "", errors.New("Name is too long.")
	}

	r.text = "Hello! " + s

	return r.text, nil
}

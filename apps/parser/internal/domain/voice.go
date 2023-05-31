package domain

import "errors"

type Voice struct {
	id string
}

func NewVoice(id string) (*Voice, error) {
	if id == "" {
		return nil, errors.New("voice id is empty")
	}

	return &Voice{id: id}, nil
}

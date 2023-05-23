package domain

import "errors"

type SpokenText struct {
	text string
}

func NewSpokenText(text string) (*SpokenText, error) {
	if text == "" {
		return nil, errors.New("text is empty")
	}
	return &SpokenText{text: text}, nil
}

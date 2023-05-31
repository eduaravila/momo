package domain

import (
	"errors"
	"strconv"
)

type SoundBite struct {
	id int
}

func NewSoundBiteFromString(id string) (*SoundBite, error) {
	if id == "" {
		return nil, errors.New("sound bite id is empty")
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return NewSoundBite(idInt)
}

func NewSoundBite(id int) (*SoundBite, error) {
	if id < 0 {
		return nil, errors.New("sound bite id is empty")
	}

	return &SoundBite{id: id}, nil
}

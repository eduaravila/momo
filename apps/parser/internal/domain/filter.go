package domain

import "errors"

type Filter struct {
	id string
}

func NewFilter(id string) (*Filter, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return &Filter{id: id}, nil
}

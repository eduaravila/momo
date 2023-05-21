package adapter

import (
	"regexp"

	"github.com/eduaravila/momo/apps/parser/internal/domain"
)

// /(\w+):|\{[\d.]\}|\[[\d]+\]|[\w\.\,]+/g

type RegexpToAudioParser struct {
}

func NewRegexpToAudioParser() RegexpToAudioParser {
	return RegexpToAudioParser{}
}

func (r RegexpToAudioParser) Parse(input string) ([]string, error) {
	return r.ToStringArray(input)
}

// TODO: implement this, convert the string array to an audio object
// include all the filters, voices and chunks
func (RegexpToAudioParser) ToAudio(input []string) (*domain.Audio, error) {
	return domain.NewAudio(), nil
}

func (RegexpToAudioParser) ToStringArray(input string) ([]string, error) {
	r, err := regexp.Compile(`(\w+):|\{[\d\w.]+\}|\[[\d\w]+\]|[\w.,]+`)

	if err != nil {
		return []string{}, err
	}

	result := r.FindAllString(input, -1)

	return result, nil
}

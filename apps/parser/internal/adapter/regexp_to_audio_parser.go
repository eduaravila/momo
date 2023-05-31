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

type StringToAudioParser interface {
	Parse(input string) (domain.Segment, error)
}

// TODO: implement this, convert the string array to an audio object
// include all the filters, voices and chunks
func (RegexpToAudioParser) ToAudio(input []string) (*domain.Audio, error) {

	var stringToAudioParser StringToAudioParser = NewTextToAudioChainParser()
	audio := domain.NewAudio()
	currentChunk := domain.NewChunk()

	for _, s := range input {
		chunk, err := stringToAudioParser.Parse(s)
		if err != nil {
			continue
		}

		audio.AddChunk(chunk)
	}

	return audio, nil
}

// ^(\w+):|^(\{\s*[\d\w.]+\s*\})|^(\[\s*[\d\w]+\s*\])|[\w\.\,]+|

func (RegexpToAudioParser) ToStringArray(input string) ([]string, error) {
	r, err := regexp.Compile(`(\w+):|\{\s*[\d\w.]+\s*\}|\[\s*[\d\w]+\s*\]|[\w.,]+`)

	if err != nil {
		return []string{}, err
	}

	result := r.FindAllString(input, -1)

	return result, nil
}

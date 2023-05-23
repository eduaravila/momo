package adapter

import (
	"errors"
	"regexp"
	"strings"

	"github.com/eduaravila/momo/apps/parser/internal/domain"
)

type TextToAudioHandler interface {
	Handle(input string) (domain.Segment, error)
}

type StringToAudioParserChain struct {
	TextToAudioHandler
}

func NewTextToAudioChainParser() *StringToAudioParserChain {
	return &StringToAudioParserChain{}
}

func (s StringToAudioParserChain) Parse(input string) (domain.Segment, error) {
	return s.Handle(input)
}

type StringToSpokenTextParser struct {
	Next TextToAudioHandler
}

func (p *StringToSpokenTextParser) Parse(input string) (domain.Segment, error) {

	//\{\s*[\d\w.]+\s*\} // filter with string or number
	isProbablyAFilter, err := regexp.Compile(`\{\s*[\d\w.]+\s*\}`)

	if err != nil {
		return nil, err
	}

	ok := isProbablyAFilter.MatchString(input)

	if ok {
		if p.Next != nil {
			return p.Next.Handle(input)
		}
		return nil, errors.New("not a spoken text")
	}

	r, err := regexp.Compile(`[\w.,]+`)

	if err != nil {
		return nil, err
	}

	ok = r.MatchString(input)

	if !ok {
		if p.Next != nil {
			return p.Next.Handle(input)
		}
		return nil, errors.New("not a spoken text")
	}

	spokenText := strings.TrimSpace(input)

	return domain.NewSpokenText(spokenText)
}

type StringToFilterParser struct {
	Next TextToAudioHandler
}

func (s StringToFilterParser) Handle(input string) (domain.Segment, error) {

	r, err := regexp.Compile(`\{\s*[\d.]+\s*\}`)

	if err != nil {
		return nil, err
	}

	ok := r.MatchString(input)

	if !ok {
		if s.Next != nil {
			return s.Next.Handle(input)
		}
		return nil, errors.New("not a spoken text")
	}

	audioFilter := strings.ReplaceAll(input, "{", "")
	audioFilter = strings.ReplaceAll(audioFilter, "}", "")
	audioFilter = strings.TrimSpace(audioFilter)

	return domain.NewSpokenText(audioFilter)
}

func NewSpokenTextHandler() *StringToFilterParser {
	return &StringToFilterParser{}
}

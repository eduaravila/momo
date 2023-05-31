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
	return &StringToAudioParserChain{NewStringToSpokenTextParserWith(
		NewStringToVoiceParserWith(
			NewStringToFilterParserWith(
				NewStringToSoundBite(),
			),
		),
	)}
}

func (s StringToAudioParserChain) Parse(input string) (domain.Segment, error) {
	return s.Handle(input)
}

type StringToSpokenTextParser struct {
	Next TextToAudioHandler
}

func NewStringToSpokenTextParserWith(next TextToAudioHandler) *StringToSpokenTextParser {
	return &StringToSpokenTextParser{Next: next}
}

func NewStringToSpokenTextParser() *StringToSpokenTextParser {
	return &StringToSpokenTextParser{}
}

func (p *StringToSpokenTextParser) Handle(input string) (domain.Segment, error) {

	//\{\s*[\d\w.]+\s*\} // filter with string or number
	isProbablyAFilter, err := regexp.Compile(`\{\s*[\d\w.]+\s*\}`)

	if err != nil {
		return nil, err
	}

	ok := isProbablyAFilter.MatchString(input)

	if ok {
		return NewStringToFilterParser().Handle(input)
	}

	isProbablyASoundBite, err := regexp.Compile(`\[\s*[\d\w]+\s*\]`)

	if err != nil {
		return nil, err
	}

	ok = isProbablyASoundBite.MatchString(input)

	if ok {
		return NewStringToSoundBite().Handle(input)
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

func NewStringToFilterParser() *StringToFilterParser {
	return &StringToFilterParser{}
}

func NewStringToFilterParserWith(next TextToAudioHandler) *StringToFilterParser {
	return &StringToFilterParser{Next: next}
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

type StringToVoiceParser struct {
	Next TextToAudioHandler
}

func NewStringToVoiceParserWith(next TextToAudioHandler) *StringToVoiceParser {
	return &StringToVoiceParser{Next: next}
}

func NewStringToVoiceParser() *StringToVoiceParser {
	return &StringToVoiceParser{}
}

func (s StringToVoiceParser) Handle(input string) (domain.Segment, error) {

	r, err := regexp.Compile(`\w+:`)

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

	voice := strings.ReplaceAll(input, ":", "")
	voice = strings.TrimSpace(voice)

	return domain.NewVoice(voice)
}

type StringToSoundBite struct {
	Next TextToAudioHandler
}

func NewStringToSoundBiteWith(next TextToAudioHandler) *StringToSoundBite {
	return &StringToSoundBite{Next: next}
}

func NewStringToSoundBite() *StringToSoundBite {
	return &StringToSoundBite{}
}

func (s StringToSoundBite) Handle(input string) (domain.Segment, error) {
	isProbablyASoundBite, err := regexp.Compile(`\[\s*[\d\w]+\s*\]`)

	if err != nil {
		return nil, err
	}

	ok := isProbablyASoundBite.MatchString(input)

	if !ok {
		if s.Next != nil {
			return s.Next.Handle(input)
		}
		return nil, errors.New("not a sound bite")
	}

	soundBite := strings.ReplaceAll(input, "[", "")
	soundBite = strings.ReplaceAll(soundBite, "]", "")
	soundBite = strings.TrimSpace(soundBite)

	return domain.NewSoundBiteFromString(soundBite)
}

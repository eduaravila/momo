package domain

type Audio struct {
	chunks []Chunk
}

type Segment interface {
}

type Chunk struct {
	filters  []Filter
	segments []Segment
	voice    string
}

func NewAudio() *Audio {
	return &Audio{}
}

func (a *Audio) AddChunk(chunk Chunk) {
	a.chunks = append(a.chunks, chunk)
}

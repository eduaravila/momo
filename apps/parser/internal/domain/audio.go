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

func NewChunk() *Chunk {
	return &Chunk{}
}

func (c *Chunk) AddSegment(segment Segment) {
	c.segments = append(c.segments, segment)
}

func (c *Chunk) AddFilter(filter Filter) {
	c.filters = append(c.filters, filter)
}

func NewAudio() *Audio {
	return &Audio{}
}

func (a *Audio) AddChunk(chunk Chunk) {
	a.chunks = append(a.chunks, chunk)
}

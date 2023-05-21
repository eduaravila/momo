package adapter_test

import (
	"testing"

	"github.com/eduaravila/momo/apps/parser/internal/adapter"
	"github.com/stretchr/testify/require"
)

func TestRegexpToAudioParser_Parse(t *testing.T) {
	tests := []struct {
		input string
		out   []string
	}{
		{
			input: "hello",
			out:   []string{"hello"},
		},
		{
			input: "{7}[76]forsen:help us, chat, he is an imposter. nani: {11} we are trapped chat, only, you, can save, us. [76], help us.",
			out: []string{
				"{7}",
				"[76]",
				"forsen:",
				"help",
				"us,",
				"chat,",
				"he",
				"is",
				"an",
				"imposter.",
				"nani:",
				"{11}",
				"we",
				"are",
				"trapped",
				"chat,",
				"only,",
				"you,",
				"can",
				"save,",
				"us.",
				"[76]",
				",",
				"help",
				"us.",
			},
		},
	}

	parser := adapter.NewRegexpToAudioParser()

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			out, err := parser.Parse(test.input)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(out) != len(test.out) {
				t.Errorf("expected %d elements, got %d", len(test.out), len(out))
			}
			require.ElementsMatch(t, test.out, out)
		})
	}
}

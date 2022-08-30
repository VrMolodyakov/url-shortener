package service

import (
	"testing"

	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	shorener := NewShortener(logging.GetLogger("debug"))
	testCases := []struct {
		title    string
		randomId uint64
		want     string
	}{
		{
			title:    "Encode method testing",
			randomId: 42,
			want:     "Q",
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			got := shorener.Encode(test.randomId)
			assert.Equal(t, test.want, got)

		})
	}
}

func TestDecode(t *testing.T) {
	shorener := NewShortener(logging.GetLogger("debug"))
	testCases := []struct {
		title string
		key   string
		want  uint64
	}{
		{
			title: "Decode method testing",
			key:   "someKeyBeenEncoded",
			want:  4269408008282604311,
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			got := shorener.Decode(test.key)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestDecodeEncode(t *testing.T) {
	shorener := NewShortener(logging.GetLogger("debug"))
	testCases := []struct {
		title    string
		randomId uint64
		want     uint64
	}{
		{
			title:    "Encode then Decode method testing",
			randomId: 4269408008282604311,
			want:     4269408008282604311,
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			encoded := shorener.Encode(test.randomId)
			got := shorener.Decode(encoded)
			t.Log("GOT:", got)
			assert.Equal(t, test.want, got)
		})
	}
}

package shortener

import (
	"math"
	"strings"

	"github.com/VrMolodyakov/url-shortener/pkg/logging"
)

type shortener struct {
	logger *logging.Logger
}

func NewShortener(logger *logging.Logger) *shortener {
	return &shortener{logger: logger}
}

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(num uint64) string {
	b := make([]byte, 0)
	size := uint64(len(alphabet))
	if num == 0 {
		return string(alphabet[0])
	}

	for num > 0 {
		remainder := num % size
		b = append(b, alphabet[remainder])
		num = num / size
	}

	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func Decode(token string) uint64 {
	var r uint64
	charPos := float64(len(token)) - 1
	alphabetSize := float64(len(alphabet))
	for i := 0; i < len(token); i++ {
		symb := string(token[i])
		r += uint64(strings.Index(alphabet, symb)) * uint64(math.Pow(alphabetSize, charPos))
		charPos--
	}
	return r
}

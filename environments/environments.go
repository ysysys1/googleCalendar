package environments

import (
	"os"
	"strings"
)

type Environments interface {
	GetInvalidWordsSearchQueryString() []string
}

type environments struct{}

func NewEnvironments() Environments {
	return &environments{}
}

func (e environments) GetInvalidWordsSearchQueryString() []string {
	invalidWords := os.Getenv("INVALID_SEARCH_QUERY_STRINGS")

	return strings.Split(invalidWords, ",")
}

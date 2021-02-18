package markdown

import (
	"apishgo/pkg/apish"
	"os"
)

// MDGenerator Markdown document producer
type MDGenerator struct {
	filename string
}

// NewGenerator creates a new plugin that will generate a markdown file from
// the full API spec
func NewGenerator(filename string) MDGenerator {
	return MDGenerator{filename}
}

// CreateFromSpec produces the markdown file from the specification
func (m MDGenerator) CreateFromSpec(spec apish.APISpec) error {
	f, err := os.Open(m.filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

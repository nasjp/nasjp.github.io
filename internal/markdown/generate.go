package markdown

import (
	"bytes"
	"fmt"
	"io"
)

type generator struct {
	buf *bytes.Buffer
}

func generate(b *block) (io.Reader, error) {
	g := newGenerator()
	if err := g.generate(b); err != nil {
		return nil, err
	}
	return g.buf, nil
}

func newGenerator() *generator {
	return &generator{
		buf: bytes.NewBuffer(nil),
	}
}

func (g *generator) generate(b *block) error {
	for _, block := range b.blocks {
		switch block.kind {
		case heading:
			if err := g.pf("<h%d>", block.num); err != nil {
				return err
			}

			if err := g.generate(block); err != nil {
				return err
			}

			if err := g.p("</h1>"); err != nil {
				return err
			}
		case paragraph:
			if err := g.pf("<p>%s</p>", block.content); err != nil {
				return err
			}
		default:
			return ErrorGenerate
		}
	}

	return nil
}

func (g *generator) p(s string) error {
	if _, err := g.buf.WriteString(s); err != nil {
		return err
	}

	return nil
}

func (g *generator) pf(format string, a ...interface{}) error {
	return g.p(fmt.Sprintf(format, a...))
}

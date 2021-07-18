package markdown

import (
	"bytes"
	"fmt"
	"io"
)

type generator struct {
	buf *bytes.Buffer
}

func generate(nd *node) (io.Reader, error) {
	g := newGenerator()
	if err := g.gen(nd); err != nil {
		return nil, err
	}
	return g.buf, nil
}

func newGenerator() *generator {
	return &generator{
		buf: bytes.NewBuffer(nil),
	}
}

func (g *generator) gen(nd *node) error {
	for _, child := range nd.children {
		switch child.kind {
		case paragraph:
			if err := g.pf("<p>%s</p>", child.content); err != nil {
				return err
			}
		case heading1:
			if err := g.p("<h1>"); err != nil {
				return err
			}

			if err := g.gen(child); err != nil {
				return err
			}

			if err := g.p("</h1>"); err != nil {
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

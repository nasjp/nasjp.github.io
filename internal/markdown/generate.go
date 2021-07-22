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

			if err := g.generateInline(block); err != nil {
				return err
			}

			if err := g.pf("</h%d>", block.num); err != nil {
				return err
			}
		case blockQuote:
			if err := g.p("<blockquote>"); err != nil {
				return err
			}

			if err := g.generate(block); err != nil {
				return err
			}

			if err := g.p("</blockquote>"); err != nil {
				return err
			}
		case horizontalRule:
			if err := g.p("<hr>"); err != nil {
				return err
			}
		case paragraph:
			if err := g.p("<p>"); err != nil {
				return err
			}

			if err := g.generateInline(block); err != nil {
				return err
			}

			if err := g.p("</p>"); err != nil {
				return err
			}
		default:
			return ErrorGenerate
		}
	}

	return nil
}

func (g *generator) generateInline(b *block) error {
	for _, inline := range b.inlines {
		switch inline.kind {
		case strong:
			if err := g.pf("<strong>%s</strong>", inline.content); err != nil {
				return err
			}
		case emphasis:
			if err := g.pf("<em>%s</em>", inline.content); err != nil {
				return err
			}
		case inlineCode:
			if err := g.pf("<code>%s</code>", inline.content); err != nil {
				return err
			}
		case inlineLink:
			if err := g.pf(`<a href="%s">%s</a>`, inline.attributes["href"], inline.content); err != nil {
				return err
			}
		case inlineImage:
			if err := g.pf(`<img src="%s" alt="%s">`, inline.attributes["src"], inline.attributes["alt"]); err != nil {
				return err
			}
		case str:
			if err := g.pf("%s", inline.content); err != nil {
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

package markdown

import (
	"bufio"
	"io"
	"strings"
)

const (
	_ markdownElementKind = iota
	text
	hash
	// paragraph
	// blockquote
	// codeBlock
	// unorderedList
	// raw
	// hr
)

var checkers = []checker{
	checkText,
}

type markdownElementKind int

type markdown struct {
	makdownElements []*makdownElement
}

type makdownElement struct {
	kind markdownElementKind
	v    string
}

type context struct {
	v         string
	inProcess bool
	sc        *bufio.Scanner
	md        *markdown
}

type checker func(*context, string) (bool, tokenizer)

type tokenizer func() (*context, error)

func tokenize(r io.Reader) (*markdown, error) {
	ctx := newContext(r)
	for ctx.next() {
		var err error
		ctx, err = ctx.tokenize()
		if err != nil {
			return nil, err
		}
	}

	if ctx.v != "" {
		ctx.md.addText(ctx.v)
	}

	return ctx.md, nil
}

func newContext(r io.Reader) *context {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanRunes)

	return &context{
		sc: sc,
		md: &markdown{},
	}
}

func (ctx *context) next() bool {
	return ctx.sc.Scan()
}

func (ctx *context) tokenize() (*context, error) {
	v := strings.Join([]string{ctx.v, ctx.sc.Text()}, "")

	for _, check := range checkers {
		is, f := check(ctx, v)
		if !is {
			continue
		}

		ctx, err := f()
		if err != nil {
			return nil, err
		}

		return ctx, nil
	}

	return nil, ErrorTokenize
}

func (md *markdown) addText(v string) {
	md.makdownElements = append(md.makdownElements, &makdownElement{
		kind: text,
		v:    v,
	})
}

var _ checker = checkText

func checkText(ctx *context, v string) (bool, tokenizer) {
	f := func() (*context, error) {
		ctx.v = v
		return ctx, nil
	}

	return true, f
}

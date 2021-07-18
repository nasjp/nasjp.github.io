package markdown

import (
	"bufio"
	"io"
	"regexp"
)

type blockKind int

const (
	_ blockKind = iota
	paragraph
	heading
	italic
)

type inlineKind int

const (
	_ inlineKind = iota
)

type block struct {
	kind    blockKind
	num     int
	inlines []*inline
	blocks  []*block
	content string
}

type inline struct{}

type context struct {
	v         string
	inProcess bool
	sc        *bufio.Scanner
	document  *block
	cur       *block
}

type checker func(*context) (bool, parser)

type parser func() (*context, error)

func parse(r io.Reader) (*block, error) {
	doc := &block{}
	ctx := newContext(r, doc)
	for ctx.next() {
		ctx.v += ctx.sc.Text()
		var err error
		ctx, err = tokenizeContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	if ctx.v != "" {
		addParagraph(ctx)
	}

	return doc, nil
}

func newContext(r io.Reader, doc *block) *context {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanRunes)

	return &context{
		sc:       sc,
		document: doc,
		cur:      doc,
	}
}

func (b *block) withContext(ctx *context) *context {
	return &context{
		v:         ctx.v,
		inProcess: ctx.inProcess,
		sc:        ctx.sc,
		document:  ctx.document,
		cur:       b,
	}
}

func (ctx *context) next() bool {
	return ctx.sc.Scan()
}

func (ctx *context) withValue(v string) *context {
	return &context{
		v:         v,
		inProcess: ctx.inProcess,
		sc:        ctx.sc,
		document:  ctx.document,
		cur:       ctx.cur,
	}
}

func tokenizeContext(ctx *context) (*context, error) {
	for _, check := range []checker{
		checkItalic,
		checkHeading,
		checkParagraph,
	} {
		is, f := check(ctx)
		if !is {
			continue
		}

		ctx, err := f()
		if err != nil {
			return nil, err
		}

		return ctx, nil
	}

	return nil, nil
}

var (
	_ checker = checkHeading
	_ checker = checkParagraph
)

var (
	italicRegexp = regexp.MustCompile(`^\*(.*)\*$`)
)

func checkItalic(ctx *context) (bool, parser) {
	if !italicRegexp.MatchString(ctx.v) {
		return false, nil
	}

	submatches := italicRegexp.FindStringSubmatch(ctx.v)
	if len(submatches) != 2 {
		return false, nil
	}

	f := func() (*context, error) {
		ctx.v = submatches[1]
		if err := addItalic(ctx); err != nil {
			return nil, err
		}

		return ctx, nil
	}

	return true, f
}

func checkHeading(ctx *context) (bool, parser) {
	nums := map[string]int{
		"# ":      1,
		"## ":     2,
		"### ":    3,
		"#### ":   4,
		"##### ":  5,
		"###### ": 6,
	}

	num, ok := nums[ctx.v]

	if !ok {
		return false, nil
	}

	f := func() (*context, error) {
		ctx.v = ""

		if err := addHeading(ctx, num); err != nil {
			return nil, err
		}

		return ctx, nil
	}

	return true, f
}

func checkParagraph(ctx *context) (bool, parser) {
	f := func() (*context, error) {
		return ctx, nil
	}

	return true, f
}

func addItalic(ctx *context) error {
	if _, err := tokenizeContext(ctx.withValue(ctx.v[1 : len(ctx.v)-2])); err != nil {
		return err
	}

	ctx.cur.blocks = append(ctx.cur.blocks, &block{
		kind: italic,
	})

	return nil
}

func addHeading(ctx *context, num int) error {
	h := &block{
		kind: heading,
		num:  num,
	}

	ctx.cur.blocks = append(ctx.cur.blocks, h)

	ctx.cur = h

	return nil
}

func addParagraph(ctx *context) {
	ctx.cur.blocks = append(ctx.cur.blocks, &block{
		kind:    paragraph,
		content: ctx.v,
	})
}

package markdown

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type blockKind int

const (
	_ blockKind = iota
	paragraph
	heading
)

type inlineKind int

const (
	_ inlineKind = iota
	emphasis
	strong
	str
)

type block struct {
	kind    blockKind
	num     int
	inlines []*inline
	blocks  []*block
}

type inline struct {
	kind    inlineKind
	content string
}

type context struct {
	v          string
	inProgress bool
	sc         *bufio.Scanner
	document   *block
	cur        *block
}

type checker func(*context) (bool, parser)

type parser func() (*context, error)

func parse(r io.Reader) (*block, error) {
	doc := &block{}
	ctx := newContext(r, doc)
	for {
		next, err := parseBlock(ctx)
		if err != nil {
			return nil, err
		}

		if !next {
			break
		}
	}

	if ctx.v != "" {
		if err := addParagraph(ctx); err != nil {
			return nil, err
		}
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

func read(ctx *context) bool {
	if ctx.inProgress {
		if ctx.v == "" {
			return false
		}

		return true
	}

	if !ctx.sc.Scan() {
		return false
	}

	ctx.v += ctx.sc.Text()

	return true
}

func readLine(ctx *context) bool {
	if ctx.inProgress {
		if ctx.v == "" {
			return false
		}

		return true
	}

	var ret bool

	for ctx.sc.Scan() {
		ret = true

		v := ctx.sc.Text()
		if v == "\n" {
			break
		}
		ctx.v += v
	}

	return ret
}

func (ctx *context) withValue(v string) *context {
	return &context{
		v:          v,
		inProgress: ctx.inProgress,
		sc:         ctx.sc,
		document:   ctx.document,
		cur:        ctx.cur,
	}
}

func (ctx *context) inline() *context {
	return &context{
		inProgress: ctx.inProgress,
		sc:         ctx.sc,
		document:   ctx.document,
		cur:        ctx.cur,
	}
}

func parseBlock(ctx *context) (bool, error) {
	if !read(ctx) {
		return false, nil
	}

	checkers := []checker{
		checkHeading,
		checkParagraph,
	}

	for _, check := range checkers {
		is, parse := check(ctx)
		if !is {
			continue
		}

		if _, err := parse(); err != nil {
			return false, err
		}

		return true, nil
	}

	return true, nil
}

func parseInline(ctx *context) (bool, error) {
	if !readLine(ctx) {
		return false, nil
	}

	checkers := []checker{
		checkStrong,
		checkEmphasis,
		checkStr,
	}

	for _, check := range checkers {
		is, parse := check(ctx)
		if !is {
			continue
		}

		if _, err := parse(); err != nil {
			return false, err
		}

		return true, nil
	}

	return true, nil
}

var (
	// block
	_ checker = checkHeading
	_ checker = checkParagraph
	// inline
	_ checker = checkStrong
	_ checker = checkEmphasis
	_ checker = checkStr
)

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
	parser := func() (*context, error) {
		return ctx, nil
	}

	return true, parser
}

var emphasisRegexp = regexp.MustCompile(`^\*(.*)\*`)

func checkEmphasis(ctx *context) (bool, parser) {
	if !emphasisRegexp.MatchString(ctx.v) {
		return false, nil
	}

	submatches := emphasisRegexp.FindStringSubmatch(ctx.v)
	if len(submatches) != 2 {
		return false, nil
	}

	parser := func() (*context, error) {
		v := ctx.v

		ctx.v = strings.Trim(submatches[1], "*")
		if err := addEmphasis(ctx); err != nil {
			return nil, err
		}

		ctx.v = strings.TrimPrefix(v, submatches[1])

		return ctx, nil
	}

	return true, parser
}

var strongRegexp = regexp.MustCompile(`^\*\*(.*)\*\*`)

func checkStrong(ctx *context) (bool, parser) {
	if !strongRegexp.MatchString(ctx.v) {
		return false, nil
	}

	submatches := strongRegexp.FindStringSubmatch(ctx.v)
	if len(submatches) != 2 {
		return false, nil
	}

	parser := func() (*context, error) {
		v := ctx.v

		ctx.v = strings.Trim(submatches[1], "*")
		if err := addStrong(ctx); err != nil {
			return nil, err
		}

		ctx.v = strings.TrimPrefix(v, submatches[1])

		return ctx, nil
	}

	return true, parser
}

func checkStr(ctx *context) (bool, parser) {
	parser := func() (*context, error) {
		if err := addStr(ctx); err != nil {
			return nil, err
		}

		ctx.v = ""

		return ctx, nil
	}

	return true, parser
}

func addHeading(ctx *context, num int) error {
	h := &block{
		kind: heading,
		num:  num,
	}

	ctx.cur.blocks = append(ctx.cur.blocks, h)

	ctx.cur = h

	if _, err := parseInline(ctx); err != nil {
		return err
	}

	return nil
}

func addParagraph(ctx *context) error {
	p := &block{
		kind: paragraph,
	}

	ctx.cur.blocks = append(ctx.cur.blocks, p)

	ctx.cur = p

	ctx.inProgress = true

	if _, err := parseInline(ctx); err != nil {
		return err
	}

	return nil
}

func addEmphasis(ctx *context) error {
	ctx.cur.inlines = append(ctx.cur.inlines, &inline{
		kind:    emphasis,
		content: ctx.v,
	})

	return nil
}

func addStrong(ctx *context) error {
	ctx.cur.inlines = append(ctx.cur.inlines, &inline{
		kind:    strong,
		content: ctx.v,
	})

	return nil
}

func addStr(ctx *context) error {
	ctx.cur.inlines = append(ctx.cur.inlines, &inline{
		kind:    str,
		content: ctx.v,
	})

	return nil
}

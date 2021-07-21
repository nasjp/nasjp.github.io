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
	blockquote
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

type parser func() error

func parse(r io.Reader) (*block, error) {
	bl := &block{}
	ctx := newContext(r, bl)
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

	return bl, nil
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

func parseBlock(ctx *context) (bool, error) {
	if !readRune(ctx) {
		return false, nil
	}

	return parseCheckers(ctx, []checker{
		checkHeading,
		checkBlockquote,
		checkParagraph,
	})
}

func parseInline(ctx *context) (bool, error) {
	if !readLine(ctx) {
		return false, nil
	}

	return parseCheckers(ctx, []checker{
		checkStrong,
		checkEmphasis,
		checkStr,
	})
}

func readRune(ctx *context) bool {
	if ctx.inProgress {
		return ctx.v != ""
	}

	if !ctx.sc.Scan() {
		return false
	}

	ctx.v += ctx.sc.Text()

	return true
}

func readLine(ctx *context) bool {
	if ctx.inProgress {
		return ctx.v != ""
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

func parseCheckers(ctx *context, checkers []checker) (bool, error) {
	for _, check := range checkers {
		is, parse := check(ctx)
		if !is {
			continue
		}

		if err := parse(); err != nil {
			return false, err
		}

		return true, nil
	}

	return true, nil
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

	parser := func() error {
		ctx.v = ""

		return addHeading(ctx, num)
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

func checkBlockquote(ctx *context) (bool, parser) {
	if ctx.v != "> " {
		return false, nil
	}

	parser := func() error {
		ctx.v = ""

		return addBlockquote(ctx)
	}

	return true, parser
}

func addBlockquote(ctx *context) error {
	h := &block{
		kind: blockquote,
	}

	ctx.cur.blocks = append(ctx.cur.blocks, h)

	ctx.cur = h

	if _, err := parseBlock(ctx); err != nil {
		return err
	}

	return nil
}

func checkParagraph(ctx *context) (bool, parser) {
	parser := func() error {
		return nil
	}

	return true, parser
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

var strongRegexp = regexp.MustCompile(`^\*\*(.*)\*\*`)

func checkStrong(ctx *context) (bool, parser) {
	if !strongRegexp.MatchString(ctx.v) {
		return false, nil
	}

	submatches := strongRegexp.FindStringSubmatch(ctx.v)
	if len(submatches) != 2 {
		return false, nil
	}

	parser := func() error {
		v := ctx.v
		ctx.v = strings.Trim(submatches[1], "*")

		if err := addStrong(ctx); err != nil {
			return err
		}

		ctx.v = strings.TrimPrefix(v, submatches[1])

		return nil
	}

	return true, parser
}

func addStrong(ctx *context) error {
	ctx.cur.inlines = append(ctx.cur.inlines, &inline{
		kind:    strong,
		content: ctx.v,
	})

	return nil
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

	parser := func() error {
		v := ctx.v
		ctx.v = strings.Trim(submatches[1], "*")

		if err := addEmphasis(ctx); err != nil {
			return err
		}

		ctx.v = strings.TrimPrefix(v, submatches[1])

		return nil
	}

	return true, parser
}

func addEmphasis(ctx *context) error {
	ctx.cur.inlines = append(ctx.cur.inlines, &inline{
		kind:    emphasis,
		content: ctx.v,
	})

	return nil
}

func checkStr(ctx *context) (bool, parser) {
	parser := func() error {
		if err := addStr(ctx); err != nil {
			return err
		}

		ctx.v = ""

		return nil
	}

	return true, parser
}

func addStr(ctx *context) error {
	ctx.cur.inlines = append(ctx.cur.inlines, &inline{
		kind:    str,
		content: ctx.v,
	})

	return nil
}

func (bk blockKind) String() string {
	return map[blockKind]string{
		paragraph:  "[paragraph]",
		heading:    "[heading]",
		blockquote: "[blockquote]",
	}[bk]
}

func (ik inlineKind) String() string {
	return map[inlineKind]string{
		emphasis: "[emphasis]",
		strong:   "[strong]",
		str:      "[str]",
	}[ik]
}

package markdown

import (
	"bufio"
	"io"
)

type token struct {
	next *token
	kind tokenKind
	val  string
}

const (
	tokenP tokenKind = iota
)

type tokenKind int

func tokenize(r io.Reader) (*token, error) {
	head := &token{}
	tk := head
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanRunes)

	var str string
	for scanner.Scan() {
		c := scanner.Text()
		if c == "\n" {
			if str == "" {
				continue
			}

			tk = tk.p(str)
			str = ""

			continue
		}

		str += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return head.next, nil
}

func (tk *token) p(str string) *token {
	tk.next = &token{
		kind: tokenP,
		val:  str,
	}

	return tk.next
}

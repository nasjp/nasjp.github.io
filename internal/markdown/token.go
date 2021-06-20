package markdown

import (
	"bufio"
	"io"
)

const (
	tokenP tokenKind = iota
	tokenH1
	tokenH2
	tokenH3
	tokenH4
	tokenH5
	tokenH6
)

type tokenKind int

type token struct {
	next *token
	kind tokenKind
	val  string
}

func (kind tokenKind) increH() tokenKind {
	return map[tokenKind]tokenKind{
		tokenP:  tokenH1,
		tokenH1: tokenH2,
		tokenH2: tokenH3,
		tokenH3: tokenH4,
		tokenH4: tokenH5,
		tokenH5: tokenH6,
		tokenH6: tokenP,
	}[kind]
}

func (kind tokenKind) isH() bool {
	return map[tokenKind]bool{
		tokenH1: true,
		tokenH2: true,
		tokenH3: true,
		tokenH4: true,
		tokenH5: true,
		tokenH6: true,
	}[kind]
}

func tokenize(r io.Reader) (*token, error) {
	head := &token{}
	tk := head
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanRunes)

	var str string
	var noHeader bool

	kind := tokenP

	for scanner.Scan() {
		c := scanner.Text()
		if c == "#" && !noHeader {
			if str == "" {
				kind = kind.increH()
				if kind == tokenP {
					noHeader = true
					str = "#######"
				}
				continue
			}
		}

		if c == " " && kind.isH() {
			continue
		}

		if c == "\n" {
			if str == "" {
				continue
			}

			tk = tk.add(str, kind)
			str = ""
			kind = tokenP
			noHeader = false

			continue
		}

		str += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return head.next, nil
}

func (tk *token) add(str string, tkK tokenKind) *token {
	tk.next = &token{
		kind: tkK,
		val:  str,
	}

	return tk.next
}

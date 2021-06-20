package markdown

import (
	"io"
)

func ToHTML(r io.Reader) (io.Reader, error) {
	tk, err := tokenize(r)
	if err != nil {
		return nil, err
	}

	nd, err := parse(tk)
	if err != nil {
		return nil, err
	}

	return gen(nd)
}

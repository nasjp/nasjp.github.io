package markdown

import (
	"io"
)

func ToHTML(r io.Reader) (io.Reader, error) {
	block, err := parse(r)
	if err != nil {
		return nil, err
	}

	return nil, nil

	return generate(block)
}

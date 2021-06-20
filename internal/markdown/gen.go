package markdown

import (
	"bytes"
	"fmt"
	"io"
)

func gen(nd *node) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	for ; nd != nil; nd = nd.next {
		if _, err := buf.WriteString(fmt.Sprintf("<p>%s</p>\n", nd.content)); err != nil {
			return nil, err
		}
	}

	return buf, nil
}

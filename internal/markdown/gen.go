package markdown

import (
	"bytes"
	"fmt"
	"io"
)

func gen(nd *node) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	for ; nd != nil; nd = nd.next {
		switch nd.kind {
		case nodeP:
			if _, err := buf.WriteString(fmt.Sprintf("<p>%s</p>\n", nd.content)); err != nil {
				return nil, err
			}
		case nodeH1:
			if _, err := buf.WriteString(fmt.Sprintf("<h1>%s</h1>\n", nd.content)); err != nil {
				return nil, err
			}
		case nodeH2:
			if _, err := buf.WriteString(fmt.Sprintf("<h2>%s</h2>\n", nd.content)); err != nil {
				return nil, err
			}
		}
	}

	return buf, nil
}

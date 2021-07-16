package markdown

import "errors"

var (
	ErrorTokenize = errors.New("unexpected toeknize error")
	ErrorParse    = errors.New("unexpected parse error")
	ErrorGenerate = errors.New("unexpected generate error")
)

package markdown

type nodeKind int

const (
	root nodeKind = iota
	paragraph
	heading1
	heading2
)

type node struct {
	kind     nodeKind
	children []*node
	cur      *node
	content  string
}

func (nd *node) parse(b *makdownElement) error {
	switch b.kind {
	case hash:
		nd.children = append(nd.cur.children, &node{
			kind:    heading1,
			content: b.v,
		})

		nd.cur = nd.children[len(nd.children)-1]

		return nil

	case text:
		nd.cur.children = append(nd.cur.children, &node{
			kind:    paragraph,
			content: b.v,
		})

		nd.cur = nd.children[len(nd.children)-1]

		return nil
	default:
		return ErrorParse
	}

}

func parse(doc *markdown) (*node, error) {
	nd := &node{kind: root}
	nd.cur = nd
	for _, elm := range doc.makdownElements {
		if err := nd.parse(elm); err != nil {
			return nil, err
		}
	}

	nd.cur = nil

	return nd, nil
}

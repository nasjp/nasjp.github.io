package markdown

type node struct {
	kind     nodeKind
	next     *node
	children *node
	content  string
}

type nodeKind int

const (
	nodeP nodeKind = iota
)

func (nd *node) p(str string) *node {
	nd.next = &node{
		kind:    nodeP,
		content: str,
	}

	return nd.next
}

func parse(tk *token) (*node, error) {
	head := &node{}
	nd := head
	for ; tk != nil; tk = tk.next {
		nd = nd.p(tk.val)
	}

	return head.next, nil
}

package markdown

type node struct {
	kind nodeKind
	next *node
	// children *node
	content string
}

type nodeKind int

const (
	nodeP nodeKind = iota
	nodeH1
	nodeH2
	nodeH3
	nodeH4
	nodeH5
	nodeH6
)

func (kind tokenKind) nodeKindH() nodeKind {
	return map[tokenKind]nodeKind{
		tokenH1: nodeH1,
		tokenH2: nodeH2,
		tokenH3: nodeH3,
		tokenH4: nodeH4,
		tokenH5: nodeH5,
		tokenH6: nodeH6,
	}[kind]
}

func (nd *node) parse(tk *token) *node {
	curNodeKind := nodeP

	if tk.kind.isH() {
		curNodeKind = tk.kind.nodeKindH()
	}

	nd.next = &node{
		kind:    curNodeKind,
		content: tk.val,
	}

	return nd.next
}

func parse(tk *token) (*node, error) {
	head := &node{}
	nd := head
	for ; tk != nil; tk = tk.next {
		nd = nd.parse(tk)
	}

	return head.next, nil
}

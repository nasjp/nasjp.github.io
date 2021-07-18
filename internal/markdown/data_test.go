package markdown

type test struct {
	name     string
	markdown string
	token    *markdown
	node     *node
	html     string
}

var tests = []test{
	{
		"Empty",
		"",
		&markdown{},
		&node{kind: root},
		"",
	},
	{
		"Text",
		"foo  bar",
		&markdown{makdownElements: []*makdownElement{{kind: text, v: "foo  bar"}}},
		&node{kind: root, children: []*node{{kind: paragraph, content: "foo  bar"}}},
		"<p>foo  bar</p>",
	},
	{
		"GreaterThanSign",
		"two &gt; one",
		&markdown{makdownElements: []*makdownElement{{kind: text, v: "two &gt; one"}}},
		&node{kind: root, children: []*node{{kind: paragraph, content: "two &gt; one"}}},
		"<p>two &gt; one</p>",
	},
	{
		"Heading1",
		"# hoge",
		&markdown{makdownElements: []*makdownElement{{kind: hash, v: "#"}, {kind: text, v: "hoge"}}},
		&node{kind: root, children: []*node{{kind: heading1, content: "#", children: []*node{{kind: paragraph, content: "hoge"}}}}},
		"<h1><p>hoge</p></h1>",
	},
	{
		"Heading2",
		"## hoge",
		&markdown{makdownElements: []*makdownElement{{kind: hash, v: "##"}, {kind: text, v: "hoge"}}},
		&node{kind: root, children: []*node{{kind: heading2, content: "", children: []*node{{kind: paragraph, content: "hoge"}}}}},
		"<h2><p>hoge</p></h2>",
	},
}

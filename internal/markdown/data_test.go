package markdown

type test struct {
	name     string
	markdown string
	block    *block
	html     string
}

var tests = []test{
	{
		"Empty",
		"",
		&block{},
		"",
	},
	{
		"Paragraph",
		"foo  bar",
		&block{blocks: []*block{{kind: paragraph, content: "foo  bar"}}},
		"<p>foo  bar</p>",
	},
	{
		"GreaterThanSign",
		"two &gt; one",
		&block{blocks: []*block{{kind: paragraph, content: "two &gt; one"}}},
		"<p>two &gt; one</p>",
	},
	{
		"Heading1",
		"# hoge",
		&block{blocks: []*block{{kind: heading, num: 1, blocks: []*block{{kind: paragraph, content: "hoge"}}}}},
		"<h1><p>hoge</p></h1>",
	},
	{
		"Heading2",
		"## hoge",
		&block{blocks: []*block{{kind: heading, num: 2, blocks: []*block{{kind: paragraph, content: "hoge"}}}}},
		"<h2><p>hoge</p></h2>",
	},
}

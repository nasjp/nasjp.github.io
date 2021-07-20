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
		&block{blocks: []*block{{kind: paragraph, inlines: []*inline{{kind: str, content: "foo  bar"}}}}},
		"<p>foo  bar</p>",
	},
	{
		"GreaterThanSign",
		"two &gt; one",
		&block{blocks: []*block{{kind: paragraph, inlines: []*inline{{kind: str, content: "two &gt; one"}}}}},
		"<p>two &gt; one</p>",
	},
	{
		"Heading1",
		"# hoge",
		&block{blocks: []*block{{kind: heading, num: 1, inlines: []*inline{{kind: str, content: "hoge"}}}}},
		"<h1>hoge</h1>",
	},
	{
		"Heading2",
		"## hoge",
		&block{blocks: []*block{{kind: heading, num: 2, inlines: []*inline{{kind: str, content: "hoge"}}}}},
		"<h2>hoge</h2>",
	},
	{
		"Emphasis",
		"*hoge*",
		&block{blocks: []*block{{kind: paragraph, inlines: []*inline{{kind: emphasis, content: "hoge"}}}}},
		"<p><em>hoge</em></h2>",
	},
}

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
		"<p><em>hoge</em></p>",
	},
	{
		"Strong",
		"**hoge**",
		&block{blocks: []*block{{kind: paragraph, inlines: []*inline{{kind: strong, content: "hoge"}}}}},
		"<p><strong>hoge</strong></p>",
	},
	{
		"Blockquote",
		"> hoge",
		&block{blocks: []*block{{kind: blockquote, blocks: []*block{{kind: paragraph, inlines: []*inline{{kind: str, content: "hoge"}}}}}}},
		"<blockquote><p>hoge</p></blockquote>",
	},
	{
		"Link",
		"[hoge](http://example.com)",
		&block{blocks: []*block{{kind: paragraph, inlines: []*inline{{kind: inlineLink, content: "hoge", attributes: map[string]string{"href": "http://example.com"}}}}}},
		"<p><a href=\"http://example.com\">hoge</a></p>",
	},
}

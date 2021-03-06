package bbgo_test

import (
	"strings"
	"testing"

	"github.com/namreg/bbgo"
)

func TestBBGO_Parse(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty string", ``, ``},
		{"new line", "hello\n\nworld", `hello<br><br>world`},
		{"check escape", `<script src="'>`, `&lt;script src=&#34;&#39;&gt;`},
		{"not defined tag", `[foo][b]hello[/b][/foo]`, `[foo]<b>hello</b>[/foo]`},
		{"b", `[b][b]hello[/b]`, `<b><b>hello</b>`},
		{"img", `[img]http://example.com/logo.png[/img]`, `<img src="http://example.com/logo.png" />`},
		{"img with title", `[img="bla bla bla"]http://example.com/logo.png[/img]`, `<img title="bla bla bla" src="http://example.com/logo.png" />`},
		{"img without src", `[img][/img]`, `<img src="" />`},
		{"quote", "[quote]hello[/quote]", `<blockquote>hello</blockquote>`},
		{"nested quote", "[quote=][quote=][/quote][/quote]", `<blockquote><blockquote></blockquote></blockquote>`},
		{"quote with attr", `[quote name=Someguy]hello[/quote]`, `<blockquote><cite>Someguy said:</cite>hello</blockquote>`},
		{"url", `[url]https://en.wikipedia.org[/url]`, `<a href="https://en.wikipedia.org">https://en.wikipedia.org</a>`},
		{"url with value", `[url=https://en.wikipedia.org]English Wikipedia[/url]`, `<a href="https://en.wikipedia.org">English Wikipedia</a>`},
		{"code", `[code][b]some[/b]\n[i]stuff[/i]\n[/quote][/code][b]more[/b]`, `<pre>[b]some[/b]\n[i]stuff[/i]\n[/quote]</pre><b>more</b>`},
		{"list", `[list][*] item 1[*] item 2[*] item 3[/list]`, `<ul><li> item 1</li><li> item 2</li><li> item 3</li></ul>`},
		{"color", "[color=#00BFFF]hello[/color]", `<span style="color: #00BFFF;">hello</span>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbgo.New()
			got := b.Parse(tt.input)
			if tt.want != got {
				t.Fatalf("want = %s, got = %s", tt.want, got)
			}
		})
	}
}

func BenchmarkBBGO_Parse(b *testing.B) {
	b.StopTimer()

	var sb strings.Builder
	for i := 0; i < 20000; i++ {
		sb.WriteString("[quote=]hello")
	}

	sb.WriteString("hello")

	for i := 0; i < 20000; i++ {
		sb.WriteString("world[/quote]")
	}

	input := sb.String()

	p := bbgo.New()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(input)
	}
}

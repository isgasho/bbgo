package bbgo_test

import (
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
		{"not defined tag", `[foo][b]hello[/b][/foo]`, `[foo]<strong>hello</strong>[/foo]`},
		{"b", `[b][b]hello[/b]`, `<strong><strong>hello</strong>`},
		{"img", `[img]http://example.com/logo.png[/img]`, `<img src="http://example.com/logo.png" />`},
		{"img with title", `[img="bla bla bla"]http://example.com/logo.png[/img]`, `<img title="bla bla bla" src="http://example.com/logo.png" />`},
		{"img without src", `[img][/img]`, `<img src="" />`},
		{"quote", "[quote]hello[/quote]", `<blockquote>hello</blockquote>`},
		{"quote with attr", "[quote name=Someguy]hello[/quote]", `<blockquote><cite>Someguy said:</cite>hello</blockquote>`},
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

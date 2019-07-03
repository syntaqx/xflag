package xflag

import (
	"flag"
	"testing"
)

func TestStringSlice(t *testing.T) {
	var (
		p    *[]string
		in   = s("-f", "foo", "-f=bar")
		want = s("foo", "bar")
	)

	p = testParse(nil, in)
	if !cs(*p, want) {
		t.Errorf("Default nil %v: Got %v, Want %v", in, *p, want)
	}

	p = testParse(s(), in)
	if !cs(*p, want) {
		t.Errorf("No default %v: Got %v, Want %v", in, *p, want)
	}

	p = testParse(s("tridge", "baaz"), in)
	if !cs(*p, want) {
		t.Errorf("With default %v: Got %v, Want %v", in, *p, want)
	}

	p = testParse(nil, s())
	if !cs(*p, s()) {
		t.Errorf("Default nil and no arguments: Got %v", *p)
	}
}

func testParse(def, argv []string) *[]string {
	fs := flag.NewFlagSet("", flag.PanicOnError)
	s := StringSliceFS(fs, "f", def, "ignored")
	_ = fs.Parse(argv)
	return s
}

func s(a ...string) []string { return a }

func cs(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

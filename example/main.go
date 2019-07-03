package main

import (
	"flag"
	"fmt"

	"github.com/syntaqx/xflag"
)

func main() {
	var a = xflag.StringSlice("a", nil, "repeatable string flag")

	var b []string
	xflag.StringSliceVar(&b, "b", []string{"foo", "bar", "baz"}, "repeatable string flag (+default)")

	flag.Parse()

	fmt.Printf("a == %+s\n", a)
	fmt.Printf("b == %+s\n", b)
}

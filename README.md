# xflag

[![GoDoc](https://godoc.org/github.com/syntaqx/xflag?status.svg)](https://godoc.org/github.com/syntaqx/xflag)
[![CircleCI](https://circleci.com/gh/syntaqx/xflag.svg?style=shield)](https://circleci.com/gh/syntaqx/xflag)
[![Go Report Card](https://goreportcard.com/badge/syntaqx/xflag)](https://goreportcard.com/report/syntaqx/xflag)
[![Coverage Status](https://coveralls.io/repos/github/syntaqx/xflag/badge.svg?branch=master)](https://coveralls.io/github/syntaqx/xflag?branch=master)

Package xflag defines custom flags for since stdlib didn't.

## Usage

Define flags using `xflag.StringSlice()`

```go
var a = xflag.StringSlice("a", nil, "repeatable string flag")
```

Or alternatively, you can bind to a variable with `xflag.StringSliceVar()`

```go
var b []string
xflag.StringSliceVar(&b, "b", []string{"foo", "bar", "baz"}, "repeatable string flag (+default)")
```

Once all flags have been defined it's time to parse the command line:

```go
flag.Parse()
```

The flags may then be used directly:

```go
fmt.Printf("a == %+s\n", a) // a == &[]
fmt.Printf("b == %+s\n", b) // b == [foo bar baz]
```

- `a` - When referencing the flag directly, it's always a pointer
- `b` - Binding to a variable the value is stored directly

## Yay! 🎉

So what's going to happen from this?

```sh
go run ./example/main.go -a 1 -b 2 -a 3
```

Go on, [give it a try](https://play.golang.org/p/ngT-DnJPB-V)

## License

[mit]: https://opensource.org/licenses/MIT

This project is open source software released under the [MIT license][mit].

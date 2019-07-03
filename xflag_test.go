package xflag

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCasesStringSlice = []struct {
	args     []string
	defaults []string
	expected []string
	usage    string
}{
	{
		args:     []string{"-f", "foo", "-f=bar"},
		defaults: nil,
		expected: []string{"foo", "bar"},
	},
	{
		args:     []string{"-f", "foo", "-f=bar"},
		defaults: []string{},
		expected: []string{"foo", "bar"},
	},
	{
		args:     []string{"-f", "foo", "-f=bar"},
		defaults: []string{"tridge", "baaz"},
		expected: []string{"foo", "bar"},
	},
	{
		args:     []string{},
		defaults: nil,
		expected: []string{},
	},
}

func TestStringSlice(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	for _, tt := range testCasesStringSlice {
		tt := tt
		t.Run("", func(st *testing.T) {
			os.Args = append([]string{"cmd"}, tt.args...)
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)

			actual := StringSlice("f", tt.defaults, tt.usage)

			assert.NotPanics(st, func() { flag.Parse() })
			assert.Equal(st, &tt.expected, actual)
		})
	}
}

func TestStringSliceFS(t *testing.T) {
	for _, tt := range testCasesStringSlice {
		tt := tt
		t.Run("", func(st *testing.T) {
			fs := flag.NewFlagSet("cmd", flag.ContinueOnError)
			actual := StringSliceFS(fs, "f", tt.defaults, tt.usage)

			assert.NoError(st, fs.Parse(tt.args))
			assert.Equal(st, &tt.expected, actual)
		})
	}
}

func TestStringSliceVar(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	for _, tt := range testCasesStringSlice {
		tt := tt
		t.Run("", func(st *testing.T) {
			os.Args = append([]string{"cmd"}, tt.args...)
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)

			var actual []string
			StringSliceVar(&actual, "f", tt.defaults, tt.usage)

			assert.NotPanics(t, func() { flag.Parse() })
			assert.Equal(st, &tt.expected, &actual)
		})
	}
}

func TestStringSliceVarFS(t *testing.T) {
	for _, tt := range testCasesStringSlice {
		tt := tt
		t.Run("", func(st *testing.T) {
			fs := flag.NewFlagSet("cmd", flag.ContinueOnError)

			var actual []string
			StringSliceVarFS(fs, &actual, "f", tt.defaults, tt.usage)

			assert.NoError(st, fs.Parse(tt.args))
			assert.Equal(st, &tt.expected, &actual)
		})
	}
}

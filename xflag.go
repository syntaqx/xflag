// Package xflag defines custom flags for since stdlib didn't.
package xflag

import (
	"flag"
	"fmt"
	"sync"
)

type stringSlice struct {
	s *[]string
	d bool
	l *sync.Mutex
}

func (s *stringSlice) String() string {
	if s.s == nil || len(*s.s) == 0 {
		return ""
	}
	return fmt.Sprintf("%v", *s.s)
}

func (s *stringSlice) Set(v string) error {
	s.l.Lock()
	defer s.l.Unlock()
	if !s.d {
		*s.s = []string{}
		s.d = true
	}
	*s.s = append(*s.s, v)
	return nil
}

// StringSlice defines a stringSlice flag with specified name, default value,
// and usage string. The return value is the address of a slice variable that
// stores the value of the flag.
func StringSlice(name string, value []string, usage string) *[]string {
	return StringSliceFS(flag.CommandLine, name, value, usage)
}

// StringSliceFS defines a stringSlice flag with specified name, default value,
// and usage string. The return value is the address of a slice that stores the
// value of the flag.
func StringSliceFS(f *flag.FlagSet, name string, value []string, usage string) *[]string {
	ss := &[]string{}
	StringSliceVarFS(f, ss, name, value, usage)
	return ss
}

// StringSliceVar defines a stringSlice flag with specified name, default value,
// and usage string. The argument p points to a string variable in which to
// store the value of the flag.
func StringSliceVar(p *[]string, name string, value []string, usage string) {
	StringSliceVarFS(flag.CommandLine, p, name, value, usage)
}

// StringSliceVar defines a stringSlice flag with specified name, default value,
// and usage string. The argument p points to a string variable in which to
// store the value of the flag.
func StringSliceVarFS(f *flag.FlagSet, p *[]string, name string, value []string, usage string) {
	*p = append([]string{}, *p...)
	s := &stringSlice{s: p, l: &sync.Mutex{}}
	if value != nil {
		*s.s = append(*s.s, value...)
	}
	f.Var(s, name, usage)
}

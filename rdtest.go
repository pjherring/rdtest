package rdtest

import (
	"fmt"
	"reflect"
	"runtime"
	. "testing"
)

type Tester struct {
	t    *T
	Done func()
}

func NewTester(t *T) Tester {
	return Tester{t: t}
}

func MaybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func (t *Tester) AssertEqual(one, two interface{}) {
	if one != two {
		fatal(t.t, "%v != %v", one, two)
	}
}

func (t *Tester) AssertNil(i interface{}) {
	if i != nil {
		fatal(t.t, "%v != nil", i)
	}
}

func (t *Tester) AssertNotNil(i interface{}) {
	if i == nil {
		fatal(t.t, "%v == nil", i)
	}
}

func (t *Tester) Assert(b bool) {
	if !b {
		fatal(t.t, "assertion failed")
	}
}

func Assert(t *T, assertion bool, args ...interface{}) {
	if !assertion {
		if cnt := len(args); cnt == 0 {
			fatal(t, "Assertion failed")
		} else {
			var msg string
			var ok bool

			if msg, ok = args[0].(string); !ok {
				msg = fmt.Sprintf("%s", args[0])
			}

			if len(args) > 1 {
				fatal(t, msg, args[1:]...)
			} else {
				fatal(t, msg)
			}
		}
	}
}

func AssertNotNil(t *T, args ...interface{}) {
}

func AssertNil(t *T, args ...interface{}) {
	Assert(t, args[0] == nil, args[1:])
}

func Equal(t *T, one, two interface{}) {
	if one != two {
		fatal(t, "%v != %v", one, two)
	}
}

func IsNil(t *T, i interface{}) {
	if i == nil {
		return
	}

	if i != nil || !reflect.ValueOf(i).IsNil() {
		fatal(t, "%v != nil", i)
	}
}

func NotNil(t *T, i interface{}) {
	if i == nil || reflect.ValueOf(i).IsNil() {
		fatal(t, "%v == nil", i)
	}
}

func True(t *T, b bool) {
	if !b {
		fatal(t, "not true")
	}
}

func fatal(t *T, fmt string, i ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	t.Logf("%v (line %d)", file, line)
	t.Fatalf(fmt, i...)
}

package rdtest

import (
	"fmt"
	"reflect"
)

type CaptureMatcher struct {
	Captured interface{}
}

func NewCaptureMatcher() *CaptureMatcher {
	return &CaptureMatcher{}
}

func (c *CaptureMatcher) Matches(i interface{}) bool {
	c.Captured = i
	return true
}

func (c *CaptureMatcher) String() string {
	return fmt.Sprintf("captured %v", c.Captured)
}

type deepEqualMatcher struct {
	expected interface{}
	obtained interface{}
}

func DeepEqualMatcher(e interface{}) *deepEqualMatcher {
	return &deepEqualMatcher{
		expected: e,
	}
}

func (d *deepEqualMatcher) Matches(obtained interface{}) bool {
	d.obtained = obtained
	return reflect.DeepEqual(d.expected, d.obtained)
}

func (d *deepEqualMatcher) String() string {
	return fmt.Sprintf("should deep equal %v", d.expected)
}

type kindMatcher struct {
	expectedValue interface{}
	expectedType  reflect.Type
	obtainedValue interface{}
	obtainedType  reflect.Type
}

func KindMatcher(i interface{}) *kindMatcher {
	return &kindMatcher{
		expectedValue: i,
		expectedType:  reflect.TypeOf(i),
	}
}

func (k *kindMatcher) Matches(i interface{}) bool {
	k.obtainedValue = i

	k.obtainedType = reflect.TypeOf(k.obtainedValue)

	return k.expectedType.String() == k.obtainedType.String()
}

func (k *kindMatcher) String() string {
	return fmt.Sprintf("is same type as %v", k.expectedType)
}

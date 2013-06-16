package assert

import "fmt"

type nilValueMatcher struct {
}

func NilValue() *nilValueMatcher {
	matcher := new(nilValueMatcher)
	return matcher
}

func (m *nilValueMatcher) Matches(value interface{}) bool {
	return value == nil
}

func (m *nilValueMatcher) DescribeMismatch(value interface{}) error {
	return fmt.Errorf("Expected: is nil but: was <%v>", value)
}

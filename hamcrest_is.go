package assert

import "fmt"

type isMatcher struct {
	expectedValue interface{}
}

func Is(expectedValue interface{}) *isMatcher {
	matcher := new(isMatcher)
	matcher.expectedValue = expectedValue
	return matcher
}

func (m *isMatcher) Matches(value interface{}) bool {
	return m.expectedValue == value
}

func (m *isMatcher) DescribeMismatch(value interface{}) error {
	return fmt.Errorf("Expected: is <%v> but: was <%v>", m.expectedValue, value)
}

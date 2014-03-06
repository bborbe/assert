package assert

type Matcher interface {
	Matches(value interface{}) bool
	DescribeMismatch(value interface{}) error
	Message(message string) Matcher
}

func AssertThat(value interface{}, matcher Matcher) error {
	return That(value, matcher)
}

func That(value interface{}, matcher Matcher) error {
	if matcher.Matches(value) {
		return nil
	}
	return matcher.DescribeMismatch(value)
}

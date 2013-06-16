package assert

type Matcher interface {
	Matches(value interface{}) bool
	DescribeMismatch(value interface{}) error
}

func assertThat(value interface{}, matcher Matcher) error {
	return That(value, matcher)
}

func That(value interface{}, matcher Matcher) error {
	if matcher.Matches(value) {
		return nil
	}
	return matcher.DescribeMismatch(value)
}

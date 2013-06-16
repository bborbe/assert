package assert

import (
	"testing"
)

func TestHamcrestEqualsInt(t *testing.T) {
	{
		err := assertThat(1, Is(1))
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := assertThat(1, Is(2))
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := assertThat(1, Is(2))
		if err.Error() != "Expected: is <2> but: was <1>" {
			t.Fatal("error message missmatch")
		}
	}
}

func TestHamcrestEqualsString(t *testing.T) {
	{
		err := assertThat("a", Is("a"))
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := assertThat("a", Is("b"))
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := assertThat("a", Is("b"))
		if err.Error() != "Expected: is <b> but: was <a>" {
			t.Fatal("error message missmatch: %s", err.Error())
		}
	}
}

func TestNilValue(t *testing.T) {
	{
		err := assertThat(nil, NilValue())
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := assertThat(t, NilValue())
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := assertThat(make([]byte, 0), NilValue())
		if err.Error() != "Expected: is nil but: was <[]>" {
			t.Fatal("error message missmatch: %s", err.Error())
		}
	}
}

func TestNotNilValue(t *testing.T) {
	{
		err := assertThat(t, NotNilValue())
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := assertThat(nil, NotNilValue())
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := assertThat(nil, NotNilValue())
		if err.Error() != "Expected: is not nil but: was nil" {
			t.Fatal("error message missmatch: %s", err.Error())
		}
	}
}


package assert

import (
	"testing"
)

func TestHamcrestEqualsInt(t *testing.T) {
	{
		err := AssertThat(1, Is(1))
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat(1, Is(2))
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat(1, Is(2))
		if err.Error() != "Expected: is <2> but: was <1>" {
			t.Fatal("error message missmatch")
		}
	}
}

func TestHamcrestEqualsString(t *testing.T) {
	{
		err := AssertThat("a", Is("a"))
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat("a", Is("b"))
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat("a", Is("b"))
		if err.Error() != "Expected: is <b> but: was <a>" {
			t.Fatal("error message missmatch: %s", err.Error())
		}
	}
}

func TestNilValue(t *testing.T) {
	{
		err := AssertThat(nil, NilValue())
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat(t, NilValue())
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat(make([]byte, 0), NilValue())
		if err.Error() != "Expected: is nil but: was <[]>" {
			t.Fatal("error message missmatch: %s", err.Error())
		}
	}
}

func TestNotNilValue(t *testing.T) {
	{
		err := AssertThat(t, NotNilValue())
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat(nil, NotNilValue())
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat(nil, NotNilValue())
		if err.Error() != "Expected: is not nil but: was nil" {
			t.Fatal("error message missmatch: %s", err.Error())
		}
	}
}

type TestInterface interface {
	Print()
}

type TestObjectWithoutPrint struct{}

type TestObjectWithPrint struct{}

func (*TestObjectWithPrint) Print() {}

func TestImplements(t *testing.T) {
	{
		var i *TestInterface = nil
		value := new(TestObjectWithPrint)
		err := AssertThat(value, Implements(i))
		if err != nil {
			t.Fatal("shouldn't return error")
		}
	}
	{
		var i *TestInterface = nil
		value := new(TestObjectWithoutPrint)
		err := AssertThat(value, Implements(i))
		if err == nil {
			t.Fatal("should return error")
		}
	}
	{
		var i *TestInterface = nil
		value := new(TestObjectWithoutPrint)
		err := AssertThat(value, Implements(i))
		if err.Error() != "expected type 'assert.TestInterface' but got '*assert.TestObjectWithoutPrint'" {
			t.Fatal("errormessage is incorrect")
		}
	}
}

func TestBuildError(t *testing.T) {
	{
		err := buildError("expected type '%s' but got '%s'", "", "foo", "bar")
		if err == nil {
			t.Fatal("err is nil")
		}
		if err.Error() != "expected type 'foo' but got 'bar'" {
			t.Fatal("errormessage is incorrect")
		}
	}
	{
		err := buildError("expected type '%s' but got '%s'", "message", "foo", "bar")
		if err == nil {
			t.Fatal("err is nil")
		}
		if err.Error() != "message, expected type 'foo' but got 'bar'" {
			t.Fatal("errormessage is incorrect")
		}
	}
}

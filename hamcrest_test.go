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
		expectedValue := "expected <2> but got <1>"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
		}
	}
	{
		err := AssertThat(1, Is(2).Message("msg"))
		expectedValue := "msg, expected <2> but got <1>"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
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
		expectedValue := "expected <b> but got <a>"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
		}
	}
	{
		err := AssertThat("a", Is("b").Message("msg"))
		expectedValue := "msg, expected <b> but got <a>"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
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
		expectedValue := "expected nil but: was <[]>"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
		}
	}
	{
		err := AssertThat(make([]byte, 0), NilValue().Message("msg"))
		expectedValue := "msg, expected nil but: was <[]>"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
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
		expectedValue := "expected not nil value"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
		}
	}
	{
		err := AssertThat(nil, NotNilValue().Message("msg"))
		expectedValue := "msg, expected not nil value"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
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

func TestGt(t *testing.T) {
	{
		err := AssertThat(3, Gt(2))
		if err != nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat(2, Gt(3))
		if err == nil {
			t.Fatal("expect nil")
		}
	}
	{
		err := AssertThat(2, Gt(3))
		expectedValue := "expected <3> is greater than <2>"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
		}
	}
	{
		err := AssertThat(2, Gt(3).Message("msg"))
		expectedValue := "msg, expected <3> is greater than <2>"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
		}
	}
	{
		err := AssertThat(2.1, Gt(3))
		expectedValue := "expected type int but got float64"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
		}
	}
	{
		err := AssertThat(2.1, Gt(3).Message("msg"))
		expectedValue := "msg, expected type int but got float64"
		if err.Error() != expectedValue {
			t.Fatalf("error message missmatch, expected '%v' but was '%v'", expectedValue, err.Error())
		}
	}

}

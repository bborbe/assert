package assert

import (
	"testing"
)

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

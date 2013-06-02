package assert

import (
	"testing"
)

func TestNotNull(t *testing.T) {
	if NotNull("foo", t) != nil {
		t.Errorf("shouldn't return error")
	}
	if NotNull("foo", nil) == nil {
		t.Errorf("should return error")
	}
	if NotNull("foo", nil).Error() != "foo, expect not null value" {
		t.Errorf("errormessage is incorrect")
	}
}

func TestNull(t *testing.T) {
	if Null("foo", nil) != nil {
		t.Errorf("shouldn't return error")
	}
	if Null("foo", t) == nil {
		t.Errorf("should return error")
	}
	if Null("foo", t).Error() != "foo, expect null value" {
		t.Errorf("errormessage is incorrect")
	}
}

func TestEqualsString(t *testing.T) {
	if EqualsString("foo", "a", "a") != nil {
		t.Errorf("shouldn't return error")
	}
	if EqualsString("foo", "a", "b") == nil {
		t.Errorf("should return error")
	}
	if EqualsString("foo", "a", "b").Error() != "foo, expected 'a' but got 'b'" {
		t.Errorf("errormessage is incorrect")
	}
}

func TestEqualsInt(t *testing.T) {
	if EqualsInt("foo", 1, 1) != nil {
		t.Errorf("shouldn't return error")
	}
	if EqualsInt("foo", 1, 2) == nil {
		t.Errorf("should return error")
	}
	if EqualsInt("foo", 1, 2).Error() != "foo, expected 1 but got 2" {
		t.Errorf("errormessage is incorrect")
	}
}

func TestTrue(t *testing.T) {
	if True("foo", true) != nil {
		t.Errorf("shouldn't return error")
	}
	if True("foo", false) == nil {
		t.Errorf("should return error")
	}
	if True("foo", false).Error() != "foo, expected true but got false" {
		t.Errorf("errormessage is incorrect")
	}
}

func TestFalse(t *testing.T) {
	if False("foo", false) != nil {
		t.Errorf("shouldn't return error")
	}
	if False("foo", true) == nil {
		t.Errorf("should return error")
	}
	if False("foo", true).Error() != "foo, expected false but got true" {
		t.Errorf("errormessage is incorrect")
	}
}

type TestInterface interface {
	Print()
}

type TestObjectWithoutPrint struct{}

type TestObjectWithPrint struct{}

func (*TestObjectWithPrint) Print() {}

func TestImplements(t *testing.T) {
	var i *TestInterface = nil
	if Implements("foo", i, new(TestObjectWithPrint)) != nil {
		t.Errorf("shouldn't return error")
	}
	if Implements("foo", i, new(TestObjectWithoutPrint)) == nil {
		t.Errorf("should return error")
	}
	if Implements("foo", i, new(TestObjectWithoutPrint)).Error() != "foo, expected type 'TestInterface' but got 'TestObjectWithoutPrint'" {
		t.Errorf("errormessage is incorrect")
	}
}

func TestLength(t *testing.T) {
	if Length("foo", 4, len([]byte("test"))) != nil {
		t.Errorf("shouldn't return error")
	}
	if Length("foo", 3, len([]byte("test"))) == nil {
		t.Errorf("should return error")
	}
	if Length("foo", 3, len([]byte("test"))).Error() != "foo, expected array length 3 but got 4" {
		t.Errorf("errormessage is incorrect")
	}
}

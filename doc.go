// Copyright 2013 Benjamin Borbe. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package provide assertions for testing

	import (
		"github.com/bborbe/assert"
		"testing"
	)

	func TestEquals(t *testing.T) {
		value := "a"
		expectedValue := "a"
		err := assert.Implements("should be equals", expectedValue, value)
		if err != nil {
			t.Error(err)
		}
	}

	func TestImplements(t *testing.T) {
		var value := ...
		var expected *InterfaceName
		err := assert.Implements("should of type InterfaceName", expected, value)
		if err != nil {
			t.Error(err)
		}
	}

*/
package assert

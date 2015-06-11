// Copyright (c) 2015, Benjamin Borbe <bborbe@rocketnews.de>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package provide assertions for testing

	import (
		. "github.com/bborbe/assert"
		"testing"
	)

	func TestNotNilValue(t *testing.T) {
		value := ...
		err := assertThat(value, NotNilValue())
		if err != nil {
			t.Fatal(err)
		}
	}

	func TestNilNull(t *testing.T) {
		value := ...
		err := assertThat(value, NilValue())
		if err != nil {
			t.Fatal(err)
		}
	}

	func TestEquals(t *testing.T) {
		value := ...
		err := assertThat(value, Is("foo"))
		if err != nil {
			t.Fatal(err)
		}
	}

*/
package assert

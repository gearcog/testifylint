// Code generated by testifylint/internal/testgen. DO NOT EDIT.

package len

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenChecker(t *testing.T) {
	var arr [3]int

	// Invalid.
	{
		assert.Equal(t, len(arr), 42)                                    // want "len: use assert\\.Len"
		assert.Equalf(t, len(arr), 42, "msg with args %d %s", 42, "42")  // want "len: use assert\\.Lenf"
		assert.Equal(t, 42, len(arr))                                    // want "len: use assert\\.Len"
		assert.Equalf(t, 42, len(arr), "msg with args %d %s", 42, "42")  // want "len: use assert\\.Lenf"
		assert.True(t, len(arr) == 42)                                   // want "len: use assert\\.Len"
		assert.Truef(t, len(arr) == 42, "msg with args %d %s", 42, "42") // want "len: use assert\\.Lenf"
		assert.True(t, 42 == len(arr))                                   // want "len: use assert\\.Len"
		assert.Truef(t, 42 == len(arr), "msg with args %d %s", 42, "42") // want "len: use assert\\.Lenf"
	}

	// Valid.
	{
		assert.Len(t, arr, 42)
		assert.Lenf(t, arr, 42, "msg with args %d %s", 42, "42")
		assert.Len(t, arr, len(arr))
		assert.Lenf(t, arr, len(arr), "msg with args %d %s", 42, "42")
	}

	// Ignored.
	{
		assert.Equal(t, len(arr), len(arr))
		assert.Equalf(t, len(arr), len(arr), "msg with args %d %s", 42, "42")
		assert.True(t, len(arr) == len(arr))
		assert.Truef(t, len(arr) == len(arr), "msg with args %d %s", 42, "42")
		assert.NotEqual(t, 42, len(arr))
		assert.NotEqualf(t, 42, len(arr), "msg with args %d %s", 42, "42")
		assert.NotEqual(t, len(arr), 42)
		assert.NotEqualf(t, len(arr), 42, "msg with args %d %s", 42, "42")
		assert.Greater(t, len(arr), 42)
		assert.Greaterf(t, len(arr), 42, "msg with args %d %s", 42, "42")
		assert.Greater(t, 42, len(arr))
		assert.Greaterf(t, 42, len(arr), "msg with args %d %s", 42, "42")
		assert.GreaterOrEqual(t, len(arr), 42)
		assert.GreaterOrEqualf(t, len(arr), 42, "msg with args %d %s", 42, "42")
		assert.GreaterOrEqual(t, 42, len(arr))
		assert.GreaterOrEqualf(t, 42, len(arr), "msg with args %d %s", 42, "42")
		assert.Less(t, len(arr), 42)
		assert.Lessf(t, len(arr), 42, "msg with args %d %s", 42, "42")
		assert.Less(t, 42, len(arr))
		assert.Lessf(t, 42, len(arr), "msg with args %d %s", 42, "42")
		assert.LessOrEqual(t, len(arr), 42)
		assert.LessOrEqualf(t, len(arr), 42, "msg with args %d %s", 42, "42")
		assert.LessOrEqual(t, 42, len(arr))
		assert.LessOrEqualf(t, 42, len(arr), "msg with args %d %s", 42, "42")
		assert.True(t, 42 != len(arr))
		assert.Truef(t, 42 != len(arr), "msg with args %d %s", 42, "42")
		assert.True(t, len(arr) != 42)
		assert.Truef(t, len(arr) != 42, "msg with args %d %s", 42, "42")
		assert.True(t, 42 > len(arr))
		assert.Truef(t, 42 > len(arr), "msg with args %d %s", 42, "42")
		assert.True(t, len(arr) > 42)
		assert.Truef(t, len(arr) > 42, "msg with args %d %s", 42, "42")
		assert.True(t, 42 >= len(arr))
		assert.Truef(t, 42 >= len(arr), "msg with args %d %s", 42, "42")
		assert.True(t, len(arr) >= 42)
		assert.Truef(t, len(arr) >= 42, "msg with args %d %s", 42, "42")
		assert.True(t, 42 < len(arr))
		assert.Truef(t, 42 < len(arr), "msg with args %d %s", 42, "42")
		assert.True(t, len(arr) < 42)
		assert.Truef(t, len(arr) < 42, "msg with args %d %s", 42, "42")
		assert.True(t, 42 <= len(arr))
		assert.Truef(t, 42 <= len(arr), "msg with args %d %s", 42, "42")
		assert.True(t, len(arr) >= 42)
		assert.Truef(t, len(arr) >= 42, "msg with args %d %s", 42, "42")
		assert.False(t, 42 == len(arr))
		assert.Falsef(t, 42 == len(arr), "msg with args %d %s", 42, "42")
		assert.False(t, len(arr) == 42)
		assert.Falsef(t, len(arr) == 42, "msg with args %d %s", 42, "42")
		assert.False(t, 42 != len(arr))
		assert.Falsef(t, 42 != len(arr), "msg with args %d %s", 42, "42")
		assert.False(t, len(arr) != 42)
		assert.Falsef(t, len(arr) != 42, "msg with args %d %s", 42, "42")
		assert.False(t, 42 > len(arr))
		assert.Falsef(t, 42 > len(arr), "msg with args %d %s", 42, "42")
		assert.False(t, len(arr) > 42)
		assert.Falsef(t, len(arr) > 42, "msg with args %d %s", 42, "42")
		assert.False(t, 42 >= len(arr))
		assert.Falsef(t, 42 >= len(arr), "msg with args %d %s", 42, "42")
		assert.False(t, len(arr) >= 42)
		assert.Falsef(t, len(arr) >= 42, "msg with args %d %s", 42, "42")
		assert.False(t, 42 < len(arr))
		assert.Falsef(t, 42 < len(arr), "msg with args %d %s", 42, "42")
		assert.False(t, len(arr) < 42)
		assert.Falsef(t, len(arr) < 42, "msg with args %d %s", 42, "42")
		assert.False(t, 42 <= len(arr))
		assert.Falsef(t, 42 <= len(arr), "msg with args %d %s", 42, "42")
		assert.False(t, len(arr) <= 42)
		assert.Falsef(t, len(arr) <= 42, "msg with args %d %s", 42, "42")
	}
}

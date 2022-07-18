package testcasetesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Do(t *testing.T) {
	var err error

	outstr, _ := Do("a", 1, true)
	assert.Equal(t, "[1A]", outstr)
	t.Log(outstr)

	outstr, _ = Do("a", 1, false)
	t.Log(outstr)
	assert.Equal(t, "[1a]", outstr)

	outstr, _ = Do("a", 13, false)
	assert.Equal(t, "a", outstr)
	t.Log(outstr)

	outstr, _ = Do("a", 13, true)
	assert.Equal(t, "A", outstr)
	t.Log(outstr)

	_, err = Do("funny", 1, true)
	t.Log(err)
	assert.EqualError(t, err, "invalid s")

	_, err = Do("a", 100, false)
	t.Log(err)
	assert.EqualError(t, err, "invalid i")

}

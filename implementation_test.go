package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixEvaluate(t *testing.T) {
	assert := assert.New(t)
	res, err := PrefixEvaluate("+ 5 * - 4 2 3")
	if assert.Nil(err) {
		assert.Equal(11.0, res)
	}
	res, err = PrefixEvaluate("+ 2 2")
	if assert.Nil(err) {
		assert.Equal(4.0, res)
	}
	res, err = PrefixEvaluate("^ 3 + 31 * 9 -3")
	if assert.Nil(err) {
		assert.Equal(81.0, res)
	}
	res, err = PrefixEvaluate("+ + / 45 + 7 8 / 10 ^ 4 0.5 23")
	if assert.Nil(err) {
		assert.Equal(31.0, res)
	}
	res, err = PrefixEvaluate("- + 0 + ^ 3 3 - 10 ^ 2 3 - ^ 2 4 + 3 + 5")
	if assert.Nil(err) {
		assert.Equal(15.0, res)
	}
	res, err = PrefixEvaluate("* 2 + 12 3 *")
	assert.Equal(res, 0.0)
	assert.Error(err)
	res, err = PrefixEvaluate("")
	assert.Equal(res, 0.0)
	assert.Error(err)
	res, err = PrefixEvaluate("* 2 + 12 3 % 43")
	assert.Equal(res, 0.0)
	assert.Error(err)
	res, err = PrefixEvaluate("-2 + 12 3 % 43")
	assert.Equal(res, 0.0)
	assert.Error(err)
}

func ExamplePrefixEvaluate() {
	res, _ := PrefixEvaluate("+ 2 5")
	fmt.Println(res)

	// Output: 7
}

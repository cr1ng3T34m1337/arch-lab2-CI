package lab2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var buf []byte

type Writer struct{}

func (w *Writer) Write(data []byte) (n int, err error) {
	buf = append(buf, data...)
	return len(buf), nil
}

func TestHandlerCompute(t *testing.T) {
	assert := assert.New(t)
	reader := strings.NewReader("* 4 / - 6 1 2")
	writer := &Writer{}
	handler := ComputeHandler{
		Input:  reader,
		Output: writer,
	}
	err := handler.Compute()
	res := string(buf)
	fmt.Println("res", res)
	if assert.Nil(err) {
		assert.Equal("10", res)
	}
	buf = []byte{}
	reader = strings.NewReader("* 2 - 6 1 2")
	handler = ComputeHandler{
		Input:  reader,
		Output: writer,
	}
	err = handler.Compute()
	assert.Error(err)
}

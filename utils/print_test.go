package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintTest(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		const msg = "test"
		want := []string{msg}
		got, err := GetPrinted(func() {
			fmt.Println(msg)
		})
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("no return value", func(t *testing.T) {
		got, err := GetPrinted(func() {})
		assert.NoError(t, err)
		assert.Empty(t, got)
	})

	t.Run("readAll error handling", func(t *testing.T) {
		wantErr := errors.New("this is sparta")
		readAll = func(r io.Reader) ([]byte, error) {
			assert.NotNil(t, r)
			return nil, wantErr
		}
		defer func() { readAll = ioutil.ReadAll }()
		got, err := GetPrinted(func() {
			fmt.Println("somewhat")
		})
		assert.Equal(t, wantErr, err)
		assert.Empty(t, got)
	})
}
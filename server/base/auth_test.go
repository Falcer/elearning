package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	res := Verify("123")
	assert.Nil(t, res, "They should be Nil")
}

func TestVerifyError(t *testing.T) {
	res := Verify("123")
	assert.NotNil(t, res, "They should not be Nil")
}

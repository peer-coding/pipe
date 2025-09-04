package ulid_test

import (
	"testing"

	"github.com/peer-coding/pipe/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	ulidLen := 26

	result := ulid.New()

	assert.NotEmpty(t, result)
	assert.Equal(t, len(result), ulidLen)
}

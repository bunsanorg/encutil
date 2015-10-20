package encutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestMarshaler string

func (m *TestMarshaler) Marshal() ([]byte, error) {
	return []byte(*m), nil
}

func (m *TestMarshaler) Unmarshal(data []byte) error {
	*m = TestMarshaler(data)
	return nil
}

func TestString(t *testing.T) {
	var m TestMarshaler
	err := UnmarshalFromString(&m, "hello, world")
	assert.NoError(t, err)
	assert.Equal(t, TestMarshaler("hello, world"), m)
	data, err := MarshalToString(&m)
	assert.NoError(t, err)
	assert.Equal(t, "hello, world", data)
}

package encutil

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	err = tmpfile.Close()
	require.NoError(t, err)

	var m TestMarshaler
	err = ioutil.WriteFile(tmpfile.Name(), []byte("hello, world"), 0666)
	require.NoError(t, err)
	err = UnmarshalFromFile(&m, tmpfile.Name())
	assert.NoError(t, err)
	assert.Equal(t, TestMarshaler("hello, world"), m)
	err = MarshalToFile(&m, tmpfile.Name())
	assert.NoError(t, err)
	data, err := ioutil.ReadFile(tmpfile.Name())
	require.NoError(t, err)
	assert.Equal(t, "hello, world", string(data))
}

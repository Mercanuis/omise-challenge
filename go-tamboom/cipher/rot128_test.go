package cipher

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	TestBuffer        = []byte{128, 129, 130}
	ReverseTestBuffer = []byte{0, 1, 2}
)

func TestRot128Reader_Read(t *testing.T) {
	reader, err := NewRot128Reader(bytes.NewBuffer(TestBuffer))
	require.NoError(t, err)
	require.NotNil(t, reader)

	buf := make([]byte, 3, 3)
	n, err := reader.Read(buf)
	require.NoError(t, err)
	require.Equal(t, 3, n)
	require.Equal(t, ReverseTestBuffer, buf)
}

func TestRot128Reader_Reversible(t *testing.T) {
	reader, err := NewRot128Reader(bytes.NewBuffer(TestBuffer))
	require.NoError(t, err)
	require.NotNil(t, reader)

	reader, err = NewRot128Reader(reader)
	require.NoError(t, err)
	require.NotNil(t, reader)

	buf := make([]byte, 3, 3)
	n, err := reader.Read(buf)
	require.NoError(t, err)
	require.Equal(t, 3, n)
	require.Equal(t, TestBuffer, buf)
}

func TestRot128Writer_Write(t *testing.T) {
	buf := &bytes.Buffer{}
	writer, err := NewRot128Writer(buf)
	require.NoError(t, err)
	require.NotNil(t, writer)

	n, err := writer.Write(TestBuffer)
	require.NoError(t, err)
	require.Equal(t, 3, n)
	require.Equal(t, ReverseTestBuffer, buf.Bytes())
}

func TestRot128Writer_Reversible(t *testing.T) {
	buf := &bytes.Buffer{}
	writer, err := NewRot128Writer(buf)
	require.NoError(t, err)
	require.NotNil(t, writer)

	writer, err = NewRot128Writer(writer)
	require.NoError(t, err)
	require.NotNil(t, writer)

	n, err := writer.Write(TestBuffer)
	require.NoError(t, err)
	require.Equal(t, 3, n)
	require.Equal(t, TestBuffer, buf.Bytes())
}

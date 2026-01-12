package tests

import (
	"code/code"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata"
	res, _ := code.GetPathSize(fixturePath)
	require.Equal(t, "1.0 MB", res)
}

func TestGetSizeDirectory(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_directory_1,2mb"
	res, _ := code.GetPathSize(fixturePath)
	require.Equal(t, "1.2 MB", res)
}

func TestGetSizeFile(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_1kb.png"
	res, _ := code.GetPathSize(fixturePath)
	require.Equal(t, "1.0 kB", res)
}

func TestGetSizeError(t *testing.T) {
	fixturePath := ""
	_, err := code.GetPathSize(fixturePath)
	require.EqualError(t, err, "Empty path")
}

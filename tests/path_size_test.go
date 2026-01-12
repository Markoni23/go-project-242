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
	res, _ := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath})
	require.Equal(t, int64(1049625), res)
}

func TestGetSizeDirectory(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_directory_1,2mb"
	res, _ := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath})
	require.Equal(t, int64(1150974), res)
}

func TestGetSizeFile(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_1kb.png"
	res, _ := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath})
	require.Equal(t, int64(1049), res)
}

func TestGetSizeError(t *testing.T) {
	fixturePath := ""
	_, err := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath})
	require.EqualError(t, err, "Empty path")
}

func TestHumanSize(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_1kb.png"
	res, _ := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath})
	humanFormat := code.FormatSize(res, true)

	require.Equal(t, "1.0kB", humanFormat)
}

func TestWithoutHiddens(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_directory_with_hidden"
	res, _ := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath})
	require.Equal(t, int64(1048575), res)
}

func TestIncludeHiddens(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_directory_with_hidden"
	res, _ := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath, IncludeHiddens: true})
	require.Equal(t, int64(2097150), res)
}

func TestWithoutRecursive(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata"
	res, _ := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath})
	require.Equal(t, int64(1049625), res)
}

func TestWithRecursive(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata"
	res, _ := code.GetPathSize(&code.GetPathSizeDTO{Path: fixturePath, Recursive: true})
	require.Equal(t, int64(3257366), res)
}

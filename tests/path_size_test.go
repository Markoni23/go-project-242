package tests

import (
	"code"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata"
	res, _ := code.GetPathSize(fixturePath, false, false, false)
	require.Equal(t, "1049625B", res)
}

func TestGetSizeDirectory(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_directory_1,2mb"
	res, _ := code.GetPathSize(fixturePath, false, false, false)
	require.Equal(t, "1150974B", res)
}

func TestGetSizeFile(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_1kb.png"
	res, _ := code.GetPathSize(fixturePath, false, false, false)
	require.Equal(t, "1049B", res)
}

func TestGetSizeError(t *testing.T) {
	fixturePath := ""
	_, err := code.GetPathSize(fixturePath, false, false, false)
	require.EqualError(t, err, "Empty path")
}

func TestHumanSize(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_1kb.png"
	res, _ := code.GetPathSize(fixturePath, false, true, false)
	require.Equal(t, "1.0kB", res)
}

func TestWithoutHiddens(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_directory_with_hidden"
	res, _ := code.GetPathSize(fixturePath, false, false, false)
	require.Equal(t, "1048575B", res)
}

func TestIncludeHiddens(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_directory_with_hidden"
	res, _ := code.GetPathSize(fixturePath, false, false, true)
	require.Equal(t, "2097150B", res)
}

func TestWithoutRecursive(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata"
	res, _ := code.GetPathSize(fixturePath, false, false, false)
	require.Equal(t, "1049625B", res)
}

func TestWithRecursive(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata"
	res, _ := code.GetPathSize(fixturePath, true, false, false)
	require.Equal(t, "3257366B", res)
}

func TestFolderWithdNested(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_with_hidden_and_nested"
	res, _ := code.GetPathSize(fixturePath, true, false, false)
	require.Equal(t, "18B", res)
}

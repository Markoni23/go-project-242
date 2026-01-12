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
	require.Equal(t, int64(1049625), res)
}

func TestGetSizeDirectory(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_directory_1,2mb"
	res, _ := code.GetPathSize(fixturePath)
	require.Equal(t, int64(1150974), res)
}

func TestGetSizeFile(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_1kb.png"
	res, _ := code.GetPathSize(fixturePath)
	require.Equal(t, int64(1049), res)
}

func TestGetSizeError(t *testing.T) {
	fixturePath := ""
	_, err := code.GetPathSize(fixturePath)
	require.EqualError(t, err, "Empty path")
}

func TestHumanSize(t *testing.T) {
	cwd, _ := os.Getwd()
	fixturePath := cwd + "/testdata/fixture_1kb.png"
	res, _ := code.GetPathSize(fixturePath)
	humanFormat := code.FormatSize(res, true)

	require.Equal(t, "1.0kB", humanFormat)
}

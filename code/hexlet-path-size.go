package code

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
)

func GetPathSize(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("Empty path")
	}

	//absPath, err := filepath.Abs(path)
	//if err != nil {
	//	return "", err
	//}

	info, err := os.Lstat(path)

	if err != nil {
		return "", err
	}

	var size int64
	if info.IsDir() {

		size = getDirectorySize(path)
	} else {
		size = info.Size()
	}

	return getHumanFormat(size), nil
}

func getDirectorySize(path string) int64 {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0
	}
	var sum int64
	for _, entry := range entries {
		if entry.IsDir() {
			continue
			//sum += getDirectorySize(filepath.Join(path, entry.Name()))
		}
		info, err := entry.Info()
		if err != nil {
			return 0
		}
		sum += info.Size()
	}

	return sum
}

func getHumanFormat(size int64) string {
	return humanize.Bytes(uint64(size))
}

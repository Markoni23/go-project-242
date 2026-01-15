package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	if path == "" {
		return "", fmt.Errorf("empty path")
	}

	info, err := os.Lstat(path)

	if err != nil {
		return "", err
	}

	var size int64
	if info.IsDir() {
		size, err = getDirectorySize(path, recursive, human, all)
		if err != nil {
			return "", err
		}
	} else {
		if !all && strings.HasPrefix(info.Name(), ".") {
			size = 0
		} else {
			size = info.Size()
		}
	}

	return fmt.Sprint(FormatSize(size, human)), nil
}

func getDirectorySize(path string, recursive, human, all bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return -1, err
	}
	var sum int64
	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if entry.IsDir() {
			if recursive {
				nestedDirSize, err := getDirectorySize(filepath.Join(path, entry.Name()), recursive, human, all)
				if err != nil {
					return -1, err
				}
				sum += nestedDirSize
			}
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return -1, err
		}

		sum += info.Size()
	}

	return sum, nil
}

func FormatSize(size int64, isHumanFormat bool) string {
	if isHumanFormat {
		return humanizeBytes(size)
	}
	return fmt.Sprintf("%dB", size)
}

func humanizeBytes(bytes int64) string {
	const base = 1000
	if bytes < base {
		return fmt.Sprintf("%dB", bytes)
	}
	suffixes := " KMGTPE"

	div, exp := 1, 0
	for n := bytes ; n >= base; n /= base {
		div *= base
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(bytes)/float64(div), suffixes[exp])
}

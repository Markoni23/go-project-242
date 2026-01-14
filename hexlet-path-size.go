package code

import (
	"fmt"
	"math"
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
		size = getDirectorySize(path, recursive, human, all)
	} else {
		if !all && strings.HasPrefix(info.Name(), ".") {
			size = 0
		} else {
			size = info.Size()
		}
	}

	return fmt.Sprint(FormatSize(size, human)), nil
}

func getDirectorySize(path string, recursive, human, all bool) int64 {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0
	}
	var sum int64
	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if entry.IsDir() {
			if recursive {
				sum += getDirectorySize(filepath.Join(path, entry.Name()), recursive, human, all)
			}
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return 0
		}

		sum += info.Size()
	}

	return sum
}

func FormatSize(size int64, isHumanFormat bool) string {
	if isHumanFormat {
		return humanateBytes(uint64(size), 1000, []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"})
	}
	return fmt.Sprintf("%dB", size)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f%s"
	if val < 10 {
		f = "%.1f%s"
	}

	return fmt.Sprintf(f, val, suffix)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

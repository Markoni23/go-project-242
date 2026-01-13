package code

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type GetPathSizeDTO struct {
	Path           string
	IncludeHiddens bool
	Recursive      bool
}

func GetPathSize(dto *GetPathSizeDTO) (int64, error) {
	if dto.Path == "" {
		return -1, fmt.Errorf("Empty path")
	}

	info, err := os.Lstat(dto.Path)

	if err != nil {
		return -1, err
	}

	var size int64
	if info.IsDir() {
		size = getDirectorySize(dto)
	} else {
		size = info.Size()
	}

	return size, nil
}

func getDirectorySize(dto *GetPathSizeDTO) int64 {
	entries, err := os.ReadDir(dto.Path)
	if err != nil {
		return 0
	}
	var sum int64
	for _, entry := range entries {
		if entry.IsDir() {
			if dto.Recursive {
				sum += getDirectorySize(
					&GetPathSizeDTO{
						Path:           filepath.Join(dto.Path, entry.Name()),
						IncludeHiddens: dto.IncludeHiddens,
						Recursive:      dto.Recursive,
					},
				)
			} else {
				continue
			}
		}

		if !dto.IncludeHiddens && strings.HasPrefix(entry.Name(), ".") {
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
		return humanateBytes(uint64(size), 1000, []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"})
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

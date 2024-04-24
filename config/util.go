package config

import (
	"path/filepath"
	"strings"
)

func getFileExtension(fileName string) string {
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex == -1 || dotIndex == len(fileName)-1 {
		return ""
	}
	return fileName[dotIndex+1:]
}

func getFileFullName(fileName string, dirPath string) string {
	matches, err := filepath.Glob(filepath.Join(dirPath, fileName+"*"))
	if err != nil {
		return ""
	}

	if len(matches) > 0 {
		return matches[0]
	} else {
		return ""
	}
}

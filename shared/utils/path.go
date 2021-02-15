package utils

import (
	"path"
	"runtime"
)

// GetSourcePath returns source path
func GetSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

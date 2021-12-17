package main

import (
	"os"
	"path/filepath"
	"strings"
)

func isGoFilename(filename string) bool {
	return strings.HasSuffix(filename, ".go") ||
		strings.HasSuffix(filename, ".go2")
}

func envVarOrDefault(envKey, defaultValue string) string {
	if v := os.Getenv(envKey); v != "" {
		return v
	}
	return defaultValue
}

// filepathAbs is a faster and error-free version of filepath.Abs.
// If workdir is already available, there is no need to do a os.Getwd for
// every filepath.Abs call.
func filepathAbs(wd, filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(wd, filename)
}

package utils

import "runtime"

func GetTempDir() string {
	GOOS := runtime.GOOS
	dir := ""

	if GOOS == "darwin" || GOOS == "linux" {
		dir = "/tmp"
	} else if GOOS == "windows" {
		dir = "%TEMP%"
	}
	return dir
}

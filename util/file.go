package util

import "os"

func IsExist(path string) {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

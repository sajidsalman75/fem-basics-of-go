package fileutils

import "os"

func WriteTextFile(fileName string, content string) error {
	return os.WriteFile(fileName, []byte(content), 0644)
}

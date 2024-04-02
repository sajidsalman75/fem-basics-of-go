package main

import (
	"fem/go/files/fileutils"
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	c, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\nDouble Original: %v%v", c, c, c)
		fileutils.WriteTextFile(rootPath+"/data/output.txt", newContent)
	} else {
		fmt.Printf("ERROR PANIC!! %v", err)
	}
}

package main

import (
	"fmt"
	"strings"
)

// typemode
func getPath(path string, typemode string) (string, error) {
	// get last /
	lastIndex := strings.LastIndex(path, "/")
	if lastIndex != -1 {
		secondLastIndex := strings.LastIndex(path[:lastIndex], "/")
		if secondLastIndex != -1 {
			result := path[:secondLastIndex] + "/" + typemode
			return result, nil
		} else {
			fmt.Println("Only one / found")
			return "", fmt.Errorf("only one / found")
		}
	} else {
		return "", fmt.Errorf("No / found")
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAscii(input string, banner string) string {
	file := "banners/" + banner + ".txt"

	data, err := os.ReadFile(file)
	if err != nil {
		return "Error loading banner"
	}

	lines := strings.Split(string(data), "\n")
	var result strings.Builder

	for i := range 8 {
		for _, c := range input {
			if c < ' ' || c > '~' {
				continue
			}
			index := int(c - ' ')
			start := index * 9
			result.WriteString(lines[start+i])
		}
		result.WriteString("\n")
	}

	fmt.Println(result.String())
	return result.String()
}

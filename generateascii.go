package main

import (
	"os"
	"strings"
)

func GenerateAscii(input string, banner string) (string, error) {
	file := "banners/" + banner + ".txt"

	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	content := string(data)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")

	var result strings.Builder

	words := strings.Split(input, "\n")
	for wi, word := range words {
		if wi > 0 {
			result.WriteString("\n")
		}
		for i := range 8 {
			for _, c := range word {
				if c < ' ' || c > '~' {
					continue
				}
				index := int(c-' ') * 9
				result.WriteString(lines[index+i])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

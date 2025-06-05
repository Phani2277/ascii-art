package proccesor

import (
	"fmt"
	"strings"

	"ascii-art/banners"
)

func Processing(text string) {
	charMap := banners.Parser("banners/standard.txt")

	text = strings.ReplaceAll(text, "\n", "\\n")
	inputLines := strings.Split(text, "\\n")
	words := false
	for i := 0; i < len(inputLines); i++ {
		if len(inputLines[i]) != 0 {
			words = true
			break
		}
	}
	for i, line := range inputLines {
		if line == "" {
			if i > 0 || words {
				fmt.Print("\n")
			}
			continue
		}

		outputLines := make([]string, 8)
		for i := 0; i < 8; i++ {
			var outputLine strings.Builder
			for _, r := range line {
				if art, ok := charMap[r]; ok && i < len(art) {
					outputLine.WriteString(art[i])
				}
			}
			outputLines[i] = outputLine.String()
		}

		for j, l := range outputLines {
			fmt.Print(l)
			if j < len(outputLines)-1 {
				fmt.Print("\n")
			}
		}

		if i < len(inputLines) {
			fmt.Print("\n")
		}
	}
}

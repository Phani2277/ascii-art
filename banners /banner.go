package banners

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Parser(bannerfile string) map[rune][]string {

	content, err := os.ReadFile(bannerfile)
	if !IsValidFile(bannerfile) {
		fmt.Println("standard.txt does not match the expected hash. Please check the file.")
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	charMap := make(map[rune][]string)
	currentRune := ' '
	index := 1

	for index+8 <= len(lines) {
		charLines := lines[index : index+8]
		charMap[currentRune] = charLines
		currentRune++
		index += 9
	}

	return charMap
}

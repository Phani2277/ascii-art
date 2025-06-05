package banners

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func IsValidFile(str string) bool {
	hasher := sha256.New()
	s, err := os.ReadFile(str)
	if err != nil {
		log.Fatal(err)
	}
	hasher.Write(s)
	hashDeclaredMap := map[string]bool{"e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf": true}
	hashCalculated := hex.EncodeToString(hasher.Sum(nil))
	return hashDeclaredMap[hashCalculated]
}

func CheckASCII(input string) error {
	for _, char := range input {
		if char > 127 {
			return fmt.Errorf("invalid character found: '%s' is not an ASCII symbol", input)
		}
	}
	return nil
}

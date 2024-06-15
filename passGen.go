package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

var (
	lengthPtr = flag.Int("length", 16, "Length of the password to generate (default 16)")
	lowercasePtr = flag.Bool("lowercase", true, "Include lowercase letters in the password")
	uppercasePtr = flag.Bool("uppercase", true, "Include uppercase letters in the password")
	numbersPtr = flag.Bool("numbers", true, "Include numbers in the password")
	symbolsPtr = flag.Bool("symbols", true, "Include symbols in the password")
)

const (
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers        = "0123456789"
	symbols        = "!@#$%^&*()-_=+[]{};':\",./<>?"
)

func main() {
	flag.Parse()

	if *lengthPtr < 8 {
		fmt.Println("Error: password length must be at least 8 characters")
		return
	}

	charPool := ""
	if *lowercasePtr {
		charPool += lowercaseChars
	}
	if *uppercasePtr {
		charPool += uppercaseChars
	}
	if *numbersPtr {
		charPool += numbers
	}
	if *symbolsPtr {
		charPool += symbols
	}

	if charPool == "" {
		fmt.Println("Error: at least one character set must be included (lowercase, uppercase, numbers, or symbols)")
		return
	}

	password := generatePassword(*lengthPtr, charPool)
	fmt.Println("Generated password:", password)
}

func generatePassword(length int, charPool string) string {
	byteArray := make([]byte, length)
	maxIndex := big.NewInt(int64(len(charPool) - 1))

	for i := range byteArray {
		randInt, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			panic(err) // This should not happen in normal usage
		}
		byteArray[i] = charPool[randInt.Int64()]
	}

	return string(byteArray)	
}

package main

import (
	"strings"
	"unicode"

	"github.com/dm1trypon/mks-win/keyboard"
)

func typeText() {
	text := "Hello, World"
	runes := strings.Split(text, "")

	for _, letter := range runes {
		symb := []rune(letter)[0]

		if unicode.IsUpper(symb) {
			keyboard.ComboKeyPress([]string{"VK_SHIFT", string(unicode.ToLower(symb))})
			continue
		}

		keyboard.KeyPress(letter)
	}
}

func main() {
	typeText()
}

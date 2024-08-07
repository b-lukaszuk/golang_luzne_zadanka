// simple solution,
// not entirely idiot/error-proof, not optimized for speed
// actually it's a tweaked Caesar Cipher solution
package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func getAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func codeLetter(letter rune, rotateBy int) rune {
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	if unicode.IsUpper(letter) {
		alphabet = strings.ToUpper(alphabet)
	}
	var rot int = getAbs(rotateBy) % len(alphabet)
	var rotAlphabet string = alphabet[rot:] + alphabet[:rot]
	if rotateBy < 0 { // decoding
		alphabet, rotAlphabet = rotAlphabet, alphabet
	}
	var indOfLetter int = strings.IndexRune(alphabet, letter)
	if indOfLetter < 0 { // rune not found in alphabet
		return letter
	}
	return []rune(rotAlphabet)[indOfLetter]
}

func getRotation(letter rune) int {
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	return strings.IndexRune(alphabet, letter)
}

func codeMsg(msg string, keyword string, decodeMode bool) string {
	var rotation int = 0
	// multiply keyword, so len(keyword) >= len(msg)
	if len(keyword) < len(msg) {
		ratio := float64(len(msg)/len(keyword)) + 1
		keyword = strings.Repeat(keyword, int(math.Ceil(ratio)))
	}
	var codedMsg strings.Builder

	for i, c := range msg {
		rotation = getRotation([]rune(keyword)[i])
		if decodeMode {
			rotation *= -1
		}
		codedMsg.WriteRune(codeLetter(c, rotation))
	}

	return codedMsg.String()
}

func main() {

	fmt.Printf("Toy program.\n\n")

	var msg string = "abc"
	var keyword string = "bcd"
	var codedMsg string = codeMsg(msg, keyword, false)
	fmt.Printf("%q coded with %q keyword is: %q\n", msg, keyword, codedMsg)

	msg = "I came, I saw, I conquered"
	keyword = "golang"
	codedMsg = codeMsg(msg, keyword, false)
	fmt.Printf("%q coded with %q keyword is: %q\n", msg, keyword, codedMsg)
	fmt.Printf("%q decoded with %q keyword is: %q\n", codedMsg, keyword,
		codeMsg(codedMsg, keyword, true))

	fmt.Printf("\nThat's all. Goodbye!\n")
}

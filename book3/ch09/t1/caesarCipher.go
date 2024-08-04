// simple solution,
// not entirely idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
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

func codeMsg(msg string, rotateBy int) string {
	var codedMsg strings.Builder

	for _, c := range msg {
		codedMsg.WriteRune(codeLetter(c, rotateBy))
	}

	return codedMsg.String()
}

func main() {

	fmt.Printf("Toy program.\n\n")

	const msg string = "L fdph, L vdz, L frqtxhuhg."
	const rot int = -3
	fmt.Printf("%q decoded with rot %d is:\n", msg, rot)
	fmt.Printf("%q\n", codeMsg(msg, rot))

	fmt.Printf("\nThat's all. Goodbye!\n")
}

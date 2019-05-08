/**
 * Link of the problem
 * https://leetcode.com/problems/unique-morse-code-words
 */
package leetcode

var morse = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
}

func uniqueMorseRepresentations(words []string) int {
	bucket := map[string]string{}
	var transformed string
	for _, word := range words {
		transformed = ""
		for _, l := range word {
			transformed += morse[l]
		}

		if bucket[transformed] == "" {
			bucket[transformed] = word
		}
	}

	return len(bucket)
}

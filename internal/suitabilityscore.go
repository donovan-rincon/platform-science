package internal

import "math"

func baseSuitabilityScore(address, name string) float64 {
	lenAddress := len(address)
	lenName := len(name)
	base := 0.0
	if lenAddress%2 == 0 {
		base = float64(countVowels(name)) * 1.5
	} else {
		base = float64(countConsonants(name))
	}

	if hasCommonFactor(lenAddress, lenName) {
		base = base * 1.5
	}
	return base
}

func countVowels(s string) int {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	count := 0
	for _, c := range s {
		if vowels[c] {
			count++
		}
	}
	return count
}

func countConsonants(s string) int {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	count := 0
	for _, c := range s {
		if !vowels[c] && c != ' ' {
			count++
		}
	}
	return count
}

func hasCommonFactor(a int, b int) bool {
	for i := 2; i <= int(math.Min(float64(a), float64(b))); i++ {
		if a%i == 0 && b&i == 0 {
			return true
		}
	}
	return false
}

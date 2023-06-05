package src

import "math/rand"

const defaultLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberLetters = "1234567890"
const characterLetters = "!@#$%^&*+?"

func RandCharacter(letter []rune, l *[]rune) []rune {
	for n := range *l {
		(*l)[n] = letter[rand.Intn(len(letter))]
	}
	return *l
}

func ShuffleArray(r *[]rune) []rune {
	for i := len(*r) - 1; i >= 0; i-- {
		num := rand.Intn(len(*r))
		(*r)[i], (*r)[num] = (*r)[num], (*r)[i]
	}
	return *r
}

func RandomString(Len uint64, numLen uint64, charLen uint64) string {
	defl := make([]rune, Len-numLen-charLen)
	numl := make([]rune, numLen)
	charl := make([]rune, charLen)

	defLetters := []rune(defaultLetters)
	numLetters := []rune(numberLetters)
	charLetters := []rune(characterLetters)

	RandCharacter(defLetters, &defl)
	RandCharacter(numLetters, &numl)
	RandCharacter(charLetters, &charl)

	finalLetters := []rune(string(defl) + string(numl) + string(charl))
	return string(ShuffleArray(&finalLetters))
}

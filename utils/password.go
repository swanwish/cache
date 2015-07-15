package utils

import (
	"math/rand"
	"time"
)

type Password struct {
	Source string `json:"-"`
	Pwd    string `json:"pwd"`
	Length int64  `json:"length"`
	Date   int64  `json:"date"`
}

var (
	character = [][]string{
		{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
		{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
		{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "+", "_"}}
	baseSeed int64 = 1436489870
)

func GeneratePassword(source string, date, length int64) Password {
	//	character := [][]string{{"A", "B"}, {"a", "b"}, {"1", "2"}, {"!", "@"}}
	pwd := ""
	typeLength := len(character)
	if date == 0 {
		date = baseSeed
	}
	seed := date + getSourceOffset(source)
	if seed == 0 {
		seed = time.Now().Unix()
	}
	if length == 0 {
		length = 8
	}
	rand.Seed(seed)
	for i := int64(0); i < length; i++ {
		randomType := rand.Intn(typeLength)
		randomIndex := rand.Intn(len(character[randomType]))
		pwd += character[randomType][randomIndex]
	}
	return Password{Source: source, Pwd: pwd, Length: length, Date: date}
}

func getSourceOffset(source string) int64 {
	var offset int64
	for i := 0; i < len(source); i++ {
		offset += int64(source[i])
	}
	return offset
}

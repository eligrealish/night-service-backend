package utils

import (
	"strings"
	"unicode/utf8"
)

type CharUtils struct {
}

func NewCharUtils() *CharUtils {
	instance := &CharUtils{}
	return instance
}

func (u CharUtils) UppercaseFirst(str string) string {
	if str == "" {
		return ""
	}
	// Get the first rune
	r, size := utf8.DecodeRuneInString(str)
	// Convert the first rune to upper case and concatenate with the rest of the string
	return strings.ToUpper(string(r)) + str[size:]
}

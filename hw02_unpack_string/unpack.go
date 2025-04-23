package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builder strings.Builder
	var lastChar rune

	for i, r := range s {
		if unicode.IsDigit(r) { // Если текущий символ это цифра
			if i == 0 { // Если цифра в начале строки, строка некорректна
				return "", ErrInvalidString
			}
			if lastChar == 0 { // Если предыдущий символ был не буквой, строка некорректна
				return "", ErrInvalidString
			}
			count := int(r - '0')                                        // Конвертируем символ в цифру
			builder.WriteString(strings.Repeat(string(lastChar), count)) // Добавляем повторяющийся символ в builder
			lastChar = 0
		} else {
			if lastChar != 0 {
				builder.WriteRune(lastChar) // Добавляем предыдущий символ в builder, есди это не цифра
			}
			lastChar = r
		}
	}

	if lastChar != 0 {
		builder.WriteRune(lastChar) // Если последний символ был буквой, добавляем его в builder
	}

	return builder.String(), nil
}

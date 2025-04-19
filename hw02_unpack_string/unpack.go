package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builder strings.Builder
	lastChar := rune(0)
	count := 0
	for i, r := range s {
		if unicode.IsDigit(r) { // Если текущий символ это цифра
			// Если цифра идет в начале строки или после другой цифры, это ошибка
			if i == 0 || i > 0 && unicode.IsDigit(rune(s[i-1])) {
				return "", ErrInvalidString
			}

			if r == '0' {
				continue // если 0 ничего не делаем
			}
			count := int(r - '0') // Конвертируем символ в цифру
			builder.WriteString(strings.Repeat(string(lastChar), count))
		} else {
			// Проверка на наличие предыдущего символа для создания результата
			if lastChar != 0 && count != 0 {
				return "", ErrInvalidString
			}
			// Если это не первый символ и не цифра
			if i > 0 && !unicode.IsDigit(rune(s[i-1])) {
				builder.WriteRune(lastChar)
			}
			lastChar = r
		}
	}
	// Если последний символ не 0 то его добавляем в строку
	if lastChar != 0 && !unicode.IsDigit(rune(s[len(s)-1])) {
		builder.WriteRune(lastChar)
	}
	return builder.String(), nil
}

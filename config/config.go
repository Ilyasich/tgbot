package config

import (
	"fmt"
	"strings"

	"github.com/dbzyuzin/tgbot"
)

func CountLetters(msg tgbot.Message) (int, int) {
	vowels := "аеёиоуыюяАЕЁИОУЫЮЯ"
	consonant := "бгджвзйклмнпрстфхцчшщБГДЖВЗЙКЛМНПРСТФХЦЧШЩ"
	vowelCount := 0
	consonantCount := 0

	for _, value := range msg.Text {
		if strings.ContainsRune(vowels, value) {
			vowelCount++
		} else if strings.ContainsRune(consonant, value) {
			consonantCount++
		}

	}
	fmt.Printf("Word: %s\nГласные: %d\nСогласные: %d\n", msg.Text, vowelCount, consonantCount)
	return vowelCount, consonantCount
}

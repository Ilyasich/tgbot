package main

import (
	"fmt"
	"tgbot/config"

	"github.com/dbzyuzin/tgbot"
)

func main() {

	tgbot.RegisterHandler(func(msg tgbot.Message) {
		vowelCount, consonantCount := config.CountLetters(msg)
		msgToSend := fmt.Sprintf("Слово: %s\nГласных букв: %d\nСогласных букв: %d\n", msg.Text, vowelCount, consonantCount)
		tgbot.SendMessage(msg.ChatID, msgToSend)
	})

	tgbot.Start()
}

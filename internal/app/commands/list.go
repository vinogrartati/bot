package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "Here all ths products: \n\n"
	for _, p := range c.productService.List() {
		outputMsg += p.Title
		outputMsg += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("ошибка отправки %v", err)
	}
}

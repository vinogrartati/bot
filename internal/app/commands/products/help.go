package products

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *ProductCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - print list of commands\n"+
			"/get - get an entity\n"+
			"/list - get a list of your entity\n"+
			"/delete - delete an existing entity\n"+
			"/create - create a new entity\n"+
			"/edit - edit an entity",
	)
	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("ошибка отправки %v", err)
	}
}

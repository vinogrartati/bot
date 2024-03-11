package products

import (
	"fmt"
	"log"

	"bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *ProductCommander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	p := &product.Product{
		Title: args,
	}
	c.productService.Create(p)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("product %s created", p.Title))

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("ошибка отправки %v", err)
	}
}

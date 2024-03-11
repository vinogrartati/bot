package products

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *ProductCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	p, err := c.productService.Get(arg)
	if err != nil {
		log.Println("wrong index to delete", args)
		return
	}

	c.productService.Remove(arg)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("product %s removed", p.Title))

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("ошибка отправки %v", err)
	}
}

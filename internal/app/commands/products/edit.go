package products

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *ProductCommander) Edit(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.CommandArguments(), " ")
	if len(args) > 2 {
		log.Println("wrong args", args)
		return
	}

	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	oldP, err := c.productService.Get(index)
	if err != nil {
		log.Println("wrong product index", index)
		return
	}

	p := &product.Product{
		Title: args[1],
	}
	err = c.productService.Update(index, p)
	if err != nil {
		log.Printf("fail to update product with id %d to %v: %v", index, p, err)
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("product %s updated to %s", oldP.Title, p.Title))

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("ошибка отправки %v", err)
	}
}

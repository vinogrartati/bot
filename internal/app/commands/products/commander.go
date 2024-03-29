package products

import (
	"encoding/json"
	"fmt"
	"log"

	"bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ProductCommander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

type CommandData struct {
	Offset int `json:"offset"`
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *ProductCommander {
	return &ProductCommander{
		bot:            bot,
		productService: productService,
	}
}

func (c *ProductCommander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsedData),
		)
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("ошибка отправки %v", err)
		}
		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	case "delete":
		c.Delete(update.Message)
	case "create":
		c.Create(update.Message)
	case "edit":
		c.Edit(update.Message)
	default:
		c.Default(update.Message)
	}
}

package sender

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// TgSender is a struct that implements Sender interface
// and sends messages to Telegram chatbot
type TgSender struct {
	bot *tgbotapi.BotAPI

	receiverId int64
}

func NewTgSender(token string) *TgSender {
	bot, _ := tgbotapi.NewBotAPI(token)

	return &TgSender{
		bot:        bot,
		receiverId: getReceiverId(bot),
	}
}

// Send sends a photo and its meta data to the receiver
func (s *TgSender) Send(msg string, photo []byte) {
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo,
	}
	photoMsg := tgbotapi.NewPhotoUpload(int64(s.receiverId), photoFileBytes)
	photoMsg.Caption = msg

	_, err := s.bot.Send(photoMsg)
	if err != nil {
		s.bot.Send(tgbotapi.NewMessage(int64(s.receiverId), fmt.Sprintf("Error sending message: %v\n", err)))
	}
}

func getReceiverId(bot *tgbotapi.BotAPI) int64 {
	updates, _ := bot.GetUpdatesChan(tgbotapi.UpdateConfig{})

	for update := range updates {
		if update.Message != nil {
			// If message is command '/admin', then return chat id
			if update.Message.IsCommand() && update.Message.Command() == "admin" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("You login as admin. Chat ID: %d\n", update.Message.Chat.ID))
				bot.Send(msg)
				return update.Message.Chat.ID
			}
		}
	}

	return 0
}

package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	tg        *tgbotapi.BotAPI
	chatIDs   []int64
	parseMode string
}

type OptionFunc func(*Bot) error

func NewTelegramBot(botAuthToken string, opts ...OptionFunc) (*Bot, error) {

	tb, err := tgbotapi.NewBotAPI(botAuthToken)
	if err != nil {
		return nil, err
	}

	bot := &Bot{tg: tb}

	for _, opt := range opts {
		err := opt(bot)
		if err != nil {
			return nil, err
		}
	}
	return bot, nil
}

func (b *Bot) GetUpdates(timeout int) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(1000)
	u.Timeout = timeout

	updates, err := b.tg.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}
	return updates, nil
}

func (b *Bot) SendMessage(text string) {
	for _, id := range b.chatIDs {
		go b.sendParseMessage(id, text)
	}
}

func (b *Bot) sendTxtMessage(chatID int64, text string) {
	messager := tgbotapi.NewMessage(chatID, text)
	b.tg.Send(messager)
}

func (b *Bot) sendParseMessage(chatID int64, text string) {
	messager := tgbotapi.NewMessage(chatID, text)
	messager.ParseMode = b.parseMode
	b.tg.Send(messager)
}

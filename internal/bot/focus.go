package bot

import (
	"focusMate/internal/utils"
	"time"

	"gopkg.in/telebot.v4"
)

func StartFocus(chatID int64,msg string,bot *telebot.Bot) {
	dur, err := utils.ParseDuration(msg)
	if err != nil {
		return
	}

	chat := telebot.Chat{
		ID: chatID,
	}

	go func() {
		ticker := time.NewTicker(2 * time.Minute)
		defer ticker.Stop()

		timer := time.NewTimer(dur)
		defer timer.Stop()

		bot.Send(&chat, "Режим фокуса включён! Длительность: "+dur.String())

		select {
		case <-ticker.C:
				bot.Send(&chat,"Не отвлекайся! Тебя ждут дела!")
		case <- timer.C:
			bot.Send(&chat, "✅ Время вышло! Отличная работа!")
			return
		}
	}()
}
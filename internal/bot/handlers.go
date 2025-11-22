package bot

import (
	"focusMate/internal/utils"

	"gopkg.in/telebot.v4"
)

func HandleRemind(c telebot.Context, b *telebot.Bot) {
    m := c.Message()
    chatID := m.Chat.ID
    text, dur, err := utils.ParseMessage(m.Text)
    if err != nil {
        c.Send("Неверный формат. Пример: /remind 10m выпить воды")
        return
    }

    newReminder := NewReminder(chatID, text, dur)
    AddReminder(newReminder)
    StartReminder(newReminder, b)
    c.Send("Окей! Я напомню через " + dur.String() + ": " + text)
}

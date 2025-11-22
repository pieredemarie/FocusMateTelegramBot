package bot

import (
	"fmt"
	"time"

	"gopkg.in/telebot.v4"
)

type Reminder struct {
	ChatID   int64
	Text     string
	Duration time.Duration
}

var reminders = make(map[int64][]Reminder)

func NewReminder(chatID int64, text string, dur time.Duration) Reminder {
	return Reminder{
		ChatID:   chatID,
		Text:     text,
		Duration: dur,
	}
}

func AddReminder(rem Reminder) {
	reminders[rem.ChatID] = append(reminders[rem.ChatID], rem)
}

func RemoveReminder(slice []Reminder,rem Reminder) []Reminder {
	for i, r := range slice {
		if r == rem {
			return append(slice[:i],slice[i+1:]...)
		}
	}
	return slice
}

func StartReminder(rem Reminder,bot *telebot.Bot) {
	go func() {
		time.Sleep(rem.Duration)
		
		chat := &telebot.Chat{ID: int64(rem.ChatID)}
		_, err := bot.Send(chat,"Напоминание!" + rem.Text)
		if err != nil {
			fmt.Println(err)
		}

		reminders[rem.ChatID] = RemoveReminder(reminders[rem.ChatID], rem)
	}()
}


package bot

import (
    "gopkg.in/telebot.v4"
    "time"
)

type PomodoroSession struct {
    ChatID int64
    Active bool
    StopCh chan struct{}
}

var pomodoroSessions = make(map[int64]*PomodoroSession)

func StartPomodoro(chatID int64, bot *telebot.Bot) {
    if session, ok := pomodoroSessions[chatID]; ok && session.Active {
        bot.Send(&telebot.Chat{ID: chatID}, "У тебя уже запущен цикл")
        return
    }

    stopCh := make(chan struct{})
    pomodoroSessions[chatID] = &PomodoroSession{ChatID: chatID, Active: true, StopCh: stopCh}
    chat := &telebot.Chat{ID: chatID}

    go func() {
        bot.Send(chat, "Pomodoro старт! 25 минут работы")

        for {
            workTimer := time.NewTimer(25 * time.Minute)
            select {
            case <-workTimer.C:
                bot.Send(chat, "Время отдыха 5 минут!")
            case <-stopCh:
                workTimer.Stop()
                bot.Send(chat, "Pomodoro остановлен")
                pomodoroSessions[chatID].Active = false
                return
            }

            restTimer := time.NewTimer(5 * time.Minute)
            select {
            case <-restTimer.C:
                bot.Send(chat, "Рабочая сессия снова, вперед!")
            case <-stopCh:
                restTimer.Stop()
                bot.Send(chat, "⏹Pomodoro остановлен")
                pomodoroSessions[chatID].Active = false
                return
            }
        }
    }()
}

func HandlePomodoro(c telebot.Context, b *telebot.Bot) {
    chatID := c.Message().Chat.ID
    StartPomodoro(chatID, b)
}

func StopPomodoro(c telebot.Context, b *telebot.Bot) {
    chatID := c.Message().Chat.ID
    if session, ok := pomodoroSessions[chatID]; ok && session.Active {
        session.StopCh <- struct{}{}
    } else {
        b.Send(c.Message().Chat, "У тебя нет активного цикла")
    }
}

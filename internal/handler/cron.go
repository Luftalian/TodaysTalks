package handler

import (
	"context"
	"log"
	"time"

	"github.com/Luftalian/TodaysTalks/internal/repository"
	"github.com/traPtitech/go-traq"
)

func (h *Handler) OnCronHandler() {
	log.Printf("Cron job: %v", time.Now())
	// すべてのこのBotが参加しているチャンネルのメッセージで今日送信されたもののリストをそのチャンネルに送信
	channels, err := h.repo.GetChannels(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	getMessageParams := repository.MessageLimitationParams{
		Since: time.Now().Add(-24 * time.Hour),
		Until: time.Now(),
		Order: "desc",
	}
	for _, channel := range channels {
		messages, err := h.repo2.GetMessage(channel.ID, &getMessageParams)
		if err != nil {
			log.Println(err)
			continue
		}
		connectedText := connectMessage(messages)
		if connectedText == "" {
			continue
		}
		err = h.repo2.PostMessage(connectedText, false, channel.ID)
		if err != nil {
			log.Println(err)
		}
	}
}

func connectMessage(messages []traq.Message) string {
	var connectedText string
	for _, message := range messages {
		connectedText += "- " + message.Content + "\n"
	}
	return connectedText
}

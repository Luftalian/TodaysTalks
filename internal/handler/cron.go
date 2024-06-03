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
		Order: "asc",
	}
	for _, channel := range channels {
		messages, err := h.repo2.GetMessages(channel.ID, &getMessageParams)
		if err != nil {
			log.Println(err)
			continue
		}
		channelInfo, err := h.repo.GetChannelByID(context.Background(), channel.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		connectedText := connectOnlyFromUserIDMessage(messages, channelInfo.UserID)
		if connectedText == "" {
			continue
		}
		err = h.repo2.PostMessage(connectedText, false, channel.ID)
		if err != nil {
			log.Println(err)
		}
	}
}

func connectOnlyFromUserIDMessage(messages []traq.Message, userId string) string {
	var connectedText string
	for _, message := range messages {
		log.Println(message.UserId)
		if message.UserId == userId {
			connectedText += "- " + message.Content + "\n"
		}
	}
	return connectedText
}

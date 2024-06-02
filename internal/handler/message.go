package handler

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/Luftalian/TodaysTalks/internal/repository"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func (h *Handler) OnMessageCreatedHandler(p *payload.MessageCreated) {
	log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
	log.Println(p.Message.Embedded[0].ID)
	if ok := checkMention(p.Message.Embedded); ok {
		message := p.Message.Text
		if strings.Contains(message, "/help") {
			resp, err := h.repo2.GetBot(os.Getenv("BOT_ID"), true)
			if err != nil {
				log.Println("Failed to get bot: " + err.Error())
				return
			}
			err = h.repo2.PostMessage(resp.(string), false, p.Message.ChannelID)
			if err != nil {
				log.Println("Failed to post message: " + err.Error())
				return
			}
		} else if strings.Contains(message, "/join") {
			err := h.repo2.LetBotJoinChannel(p.Message.ChannelID)
			if err != nil {
				log.Println("Failed to let bot join channel: " + err.Error())
				return
			}
			err = h.repo2.PostMessage("joined", false, p.Message.ChannelID)
			if err != nil {
				log.Println("Failed to post message: " + err.Error())
				return
			}
			subscribeChannelParams := repository.SubscribeChannelParams{
				ID:   p.Message.ChannelID,
				Name: p.Message.ChannelID,
			}
			err = h.repo.SubscribeChannel(context.Background(), subscribeChannelParams)
			if err != nil {
				log.Println("Failed to subscribe channel: " + err.Error())
				return
			}
		} else if strings.Contains(message, "/leave") {
			err := h.repo2.LetBotLeaveChannel(p.Message.ChannelID)
			if err != nil {
				log.Println("Failed to let bot leave channel: " + err.Error())
				return
			}
			err = h.repo2.PostMessage("left", false, p.Message.ChannelID)
			if err != nil {
				log.Println("Failed to post message: " + err.Error())
				return
			}
			err = h.repo.UnsubscribeChannel(context.Background(), p.Message.ChannelID)
			if err != nil {
				log.Println("Failed to unsubscribe channel: " + err.Error())
				return
			}
		}
	}
}

func checkMention(embedded []payload.EmbeddedInfo) bool {
	for _, embed := range embedded {
		if embed.Raw == os.Getenv("BOT_NAME") && embed.Type == "user" {
			return true
		}
	}
	return false
}

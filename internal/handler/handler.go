package handler

import (
	"github.com/Luftalian/TodaysTalks/internal/repository"
	"github.com/robfig/cron/v3"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

type Handler struct {
	repo  *repository.Repository
	repo2 *repository.Repository2
}

func New(repo *repository.Repository, repo2 *repository.Repository2) *Handler {
	return &Handler{
		repo:  repo,
		repo2: repo2,
	}
}

func (h *Handler) SetupSubscriptionEvent(bot *traqwsbot.Bot) {
	bot.OnPing(h.OnPingHandler)
	bot.OnError(h.OnErrorHandler)
	bot.OnMessageCreated(h.OnMessageCreatedHandler)
}

func (h *Handler) SetUpCron(c *cron.Cron) {
	c.AddFunc("20 * * * * *", h.OnCronHandler)
}

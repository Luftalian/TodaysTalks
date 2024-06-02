package main

import (
	"os"
	"time"

	"github.com/Luftalian/TodaysTalks/internal/handler"
	"github.com/Luftalian/TodaysTalks/internal/repository"
	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron/v3"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

func main() {
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: os.Getenv("ACCESS_TOKEN"), // Required
		Origin:      "wss://q.trap.jp",         // Optional (default: wss://q.trap.jp)
	})
	if err != nil {
		panic(err)
	}

	// // connect to database
	// db, err := sqlx.Connect("mysql", config.MySQL().FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// // migrate tables
	// if err := migration.MigrateTables(db.DB); err != nil {
	// 	log.Fatal(err)
	// }
	db := &sqlx.DB{}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	c := cron.New(cron.WithLocation(loc))

	// setup repository
	repo := repository.New(db)
	repo2 := repository.New2(bot.API())

	// setup routes
	h := handler.New(repo, repo2)
	h.SetupSubscriptionEvent(bot)
	h.SetUpCron(c)

	c.Start()

	if err := bot.Start(); err != nil {
		panic(err)
	}
}

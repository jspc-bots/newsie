package main

import (
	"time"

	"github.com/jspc/bottom"
	"github.com/lrstanley/girc"
)

type Bot struct {
	bottom bottom.Bottom
}

func New(user, password, server string, verify bool, tz string, feeds Feeds) (b Bot, err error) {
	timezone, err := time.LoadLocation(tz)
	if err != nil {
		return
	}

	b.bottom, err = bottom.New(user, password, server, verify)
	if err != nil {
		return
	}

	b.bottom.Client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
		c.Cmd.Join(Chan)
	})

	router := bottom.NewRouter()
	router.AddRoute(`get\s+headlines`, func(_, channel string, _ []string) (err error) {
		headlines, err := feeds.Headlines()
		if err != nil {
			return
		}

		for _, h := range headlines {
			b.bottom.Client.Cmd.Messagef(channel, "%s: %s (Read: %s)", h.Published.In(timezone).Format("Jan 2 15:04"), h.Title, h.Url)
		}

		return
	})

	b.bottom.Middlewares.Push(router)

	return
}

package main

import (
	"sort"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type Headline struct {
	Title     string
	Published time.Time
	Url       string
}

type Feeds []string

func NewFeeds(s string) Feeds {
	return Feeds(strings.Split(s, ","))
}

func (f Feeds) Headlines() (h []Headline, err error) {
	h = make([]Headline, 0)
	p := gofeed.NewParser()

	var feed *gofeed.Feed
	for _, url := range f {
		feed, err = p.ParseURL(url)
		if err != nil {
			return
		}

		for _, i := range feed.Items {
			h = append(h, Headline{
				Title:     i.Title,
				Published: *i.PublishedParsed,
				Url:       i.Link,
			})
		}

	}

	sort.Slice(h, func(i, j int) bool {
		return h[j].Published.Before(h[i].Published)
	})

	if len(h) > 25 {
		h = h[:25]
	}

	return
}

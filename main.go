package main

import (
	"os"
)

const (
	Nick = "newsie"
	Chan = "#dashboard"
)

var (
	Username  = os.Getenv("SASL_USER")
	Password  = os.Getenv("SASL_PASSWORD")
	Server    = os.Getenv("SERVER")
	VerifyTLS = os.Getenv("VERIFY_TLS") == "true"
	RssFeeds  = os.Getenv("RSS_FEEDS")
	Timezone  = os.Getenv("TZ")
)

func main() {
	c, err := New(Username, Password, Server, VerifyTLS, Timezone, NewFeeds(RssFeeds))
	if err != nil {
		panic(err)
	}

	panic(c.bottom.Client.Connect())
}

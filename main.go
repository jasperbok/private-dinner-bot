package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle(tele.OnText, handleRecipe)

	b.Start()
}

func handleRecipe(c tele.Context) error {
	user := c.Sender()
	text := c.Text()

	if strings.ToLower(text) == "recept" {
		recipes := GetRecipes()
		choice := recipes[rand.Intn(len(recipes))]
		return c.Send(fmt.Sprintf("Wat dacht je van %s?", strings.ToLower(choice.Title)))
	}

	return c.Send(fmt.Sprintf("%s zei '%s'", user.Username, text))
}

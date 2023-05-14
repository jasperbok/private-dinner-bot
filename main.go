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

func handleYesBtn(c tele.Context) error {
	responses := []string{
		"%s zegt in ieder geval 'ja'",
		"Da's een 'ja' van %s",
		"%s is akkoord",
		"%s zegt 'ja'",
	}
	return c.Send(fmt.Sprintf(responses[rand.Intn(len(responses))], c.Sender().Username))
}

func handleNoBtn(c tele.Context) error {
	responses := []string{
		"%s wilt het niet",
		"%s zegt 'neen'",
		"Geen akkoord van %s",
		"Van %s moet ik wat anders bedenken",
	}
	return c.Send(fmt.Sprintf(responses[rand.Intn(len(responses))], c.Sender().Username))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/recept", func(c tele.Context) error {
		// Universal markup builders.
		selector := &tele.ReplyMarkup{}

		// Inline buttons.
		inYes := selector.Data("Aw yis", "yes")
		inNo := selector.Data("Nooooz", "no")

		selector.Inline(selector.Row(inYes, inNo))

		b.Handle(&inYes, handleYesBtn)
		b.Handle(&inNo, handleNoBtn)

		recipes := GetRecipes()
		choice := recipes[rand.Intn(len(recipes))]
		return c.Send(fmt.Sprintf("Wat dacht je van %s?", strings.ToLower(choice.Title)), selector)
	})

	b.Handle(tele.OnText, handleText)

	b.Start()
}

func handleText(c tele.Context) error {
	user := c.Sender()
	text := c.Text()

	return c.Send(fmt.Sprintf("%s zei '%s'", user.Username, text))
}

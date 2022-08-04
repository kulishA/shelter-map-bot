package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"log"
	m "shelter-map-bot/internal/bot/message"
	"shelter-map-bot/internal/config"
	"shelter-map-bot/internal/transport/grpc"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Error while initialise config: %s", err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(conf.Bot.Key)
	if err != nil {
		log.Panic(err)
	}

	api, err := grpc.CreateAndConnect(fmt.Sprintf("%s:%s", conf.Api.Host, conf.Api.Port))
	if err != nil {
		logrus.Fatalf("Error while initialise api client: %s", err.Error())
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				message := tgbotapi.NewMessage(update.Message.Chat.ID, m.GetHelpMessage())
				if _, err := bot.Send(message); err != nil {
					logrus.Errorf("Error while send message: %s", err.Error())
				}
				break
			case "start":
				message := tgbotapi.NewMessage(update.Message.Chat.ID, m.GetStartMessage())
				btn := tgbotapi.KeyboardButton{
					RequestLocation: true,
					Text:            "Надіслати мою локацію 📍",
				}
				message.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn})
				if _, err := bot.Send(message); err != nil {
					logrus.Errorf("Error while send message: %s", err.Error())
				}
				break
			}
		} else {
			location := update.Message.Location
			if location == nil {
				message := tgbotapi.NewMessage(update.Message.Chat.ID, m.GetStartMessage())
				if _, err := bot.Send(message); err != nil {
					logrus.Errorf("Error while send message: %s", err.Error())
				}
			} else {
				shelter, err := api.GetFirstNearbyShelter(float32(location.Latitude), float32(location.Longitude))
				if err != nil {
					message := tgbotapi.NewMessage(update.Message.Chat.ID, m.GetErrorMessage())
					if _, err := bot.Send(message); err != nil {
						logrus.Errorf("Error while send message: %s", err.Error())
					}
				}
				shelterLocation := tgbotapi.NewLocation(update.Message.Chat.ID, float64(shelter.Latitude), float64(shelter.Longitude))
				if _, err := bot.Send(shelterLocation); err != nil {
					logrus.Errorf("Error while send message: %s", err.Error())
				}
				shelterInfo := tgbotapi.NewMessage(update.Message.Chat.ID, m.GetShelterDescriptionMessage(shelter))
				if _, err := bot.Send(shelterInfo); err != nil {
					logrus.Errorf("Error while send message: %s", err.Error())
				}
			}
		}
	}
}

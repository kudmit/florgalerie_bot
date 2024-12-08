// ("")
// https://github.com/kudmit/florgalerie_bot.git
package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("Token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Переменные для хранения выбранного языка пользователей
	userLanguage := make(map[int64]string)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				// Выбор языка
				buttons := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("DEU"),
						tgbotapi.NewKeyboardButton("EN"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("UK"),
						tgbotapi.NewKeyboardButton("RU"),
					),
				)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Select a language:")
				msg.ReplyMarkup = buttons
				bot.Send(msg)
			case "DEU":
				userLanguage[update.Message.Chat.ID] = "DEU"
				sendGreeting(bot, update.Message.Chat.ID, "Willkommen in unserem Geschäft 💐Florgalerie!\n\nIch helfe Ihnen, einen Strauß nach Ihrem Geschmack auszuwählen oder aus den verfügbaren Optionen in unserem Telegram-Kanal auszuwählen. Außerdem können Sie eine Bestellung zu einer bestimmten Zeit aufgeben.⏱️\n\nAktuelles Angebot:\nVon August bis Oktober ist die Saison der zarten Tulpen⚘️!", "Weiter!")
			case "EN":
				userLanguage[update.Message.Chat.ID] = "EN"
				sendGreeting(bot, update.Message.Chat.ID, "Welcome to our store 💐Florgalerie!\n\nI will help you choose a bouquet to your liking or pick from the available options in our Telegram channel, as well as place an order for a specific time.⏱️\n\nCurrent offer:\nFrom August to October, enjoy the season of delicate tulips⚘️!", "Next!")
			case "UK":
				userLanguage[update.Message.Chat.ID] = "UK"
				sendGreeting(bot, update.Message.Chat.ID, "Ласкаво просимо до нашого магазину 💐Florgalerie!\n\nЯ допоможу вам підібрати букет на ваш смак або вибрати з уже наявних у нашому Telegram-каналі, а також оформити замовлення на певний час.⏱️\n\nАктуальна пропозиція:\nЗ серпня по жовтень триває сезон ніжних тюльпанів⚘️!", "Далі!")
			case "RU":
				userLanguage[update.Message.Chat.ID] = "RU"
				sendGreeting(bot, update.Message.Chat.ID, "Приветствуем Вас в нашем магазине 💐Florgalerie!\n\nЯ помогу Вам подобрать букет на Ваш вкус или выбрать из из уже имеющихся в нашем Telegram канале, а так же оформить заказ к определенному времени.⏱️\n\nАктуальное предложение:\nС августа по октябрь проходит сезон нежных тюльпанов⚘️!", "Далее!")
			case "Weiter!", "Next!", "Далі!", "Далее!":
				// Следующий вопрос: "Хотите выбрать букет или создать свой?"
				lang := userLanguage[update.Message.Chat.ID]
				var question, chooseButton, createButton string
				switch lang {
				case "DEU":
					question = "Möchten Sie einen Blumenstrauß aus unserem Telegram-Kanal auswählen oder Ihren eigenen zusammenstellen?💐"
					chooseButton = "Auswählen!"
					createButton = "Mein eigener Strauß!"
				case "EN":
					question = "Would you like to choose a bouquet from our Telegram channel or create your own?💐"
					chooseButton = "Choose!"
					createButton = "Create my own!"
				case "UK":
					question = "Хотіли б ви вибрати букет із нашого Telegram-каналу чи скомпонувати свій власний?💐"
					chooseButton = "Вибрати!"
					createButton = "Створити свій!"
				case "RU":
					question = "Хотели бы вы выбрать букет из множества уже имеющихся в нашем Telegram канале или скомпоновать свой💐?"
					chooseButton = "Выбрать!"
					createButton = "Создать свой!"
				}

				sendQuestion(bot, update.Message.Chat.ID, question, chooseButton, createButton)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please select a valid option.")
				bot.Send(msg)
			}
		}
	}
}

// sendGreeting отправляет приветственное сообщение и кнопку "Далее!"
func sendGreeting(bot *tgbotapi.BotAPI, chatID int64, message string, nextButton string) {
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(nextButton),
		),
	)
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}

// sendQuestion отправляет следующий вопрос с кнопками "Выбрать!" и "Создать свой!"
func sendQuestion(bot *tgbotapi.BotAPI, chatID int64, question, chooseButton, createButton string) {
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(chooseButton),
			tgbotapi.NewKeyboardButton(createButton),
		),
	)
	msg := tgbotapi.NewMessage(chatID, question)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}

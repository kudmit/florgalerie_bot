// ("7605031210:AAGTiIboCT3mxxLO6egJ3Zhkr8LAVcdu6yo")
// https://github.com/kudmit/florgalerie_bot.git
package main

import (
	"log"
	"regexp"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Локализованное сообщение о сохранении букета
func sendBouquetSavedMessage(bot *tgbotapi.BotAPI, chatID int64, lang string, details string) {
	var message string
	switch lang {
	case "DEU":
		message = "Ihr Strauß wurde gespeichert: " + details
	case "EN":
		message = "Your bouquet has been saved: " + details
	case "UK":
		message = "Ваш букет збережено: " + details
	case "RU":
		message = "Ваш букет сохранен: " + details
	}
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

// Приветствие
func sendGreeting(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var greeting, nextButton string
	switch lang {
	case "DEU":
		greeting = "Willkommen in unserem Geschäft 💐Florgalerie!"
		nextButton = "Weiter!"
	case "EN":
		greeting = "Welcome to our store 💐Florgalerie!"
		nextButton = "Next!"
	case "UK":
		greeting = "Ласкаво просимо до нашого магазину 💐Florgalerie!"
		nextButton = "Далі!"
	case "RU":
		greeting = "Приветствуем Вас в нашем магазине 💐Florgalerie!"
		nextButton = "Далее!"
	}

	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(nextButton),
		),
	)
	msg := tgbotapi.NewMessage(chatID, greeting)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}

// Вопрос о букете или создании
func sendQuestion(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var question, button1, button2 string
	switch lang {
	case "DEU":
		question = "Möchten Sie einen Blumenstrauß auswählen oder Ihren eigenen zusammenstellen?"
		button1 = "Auswählen!"
		button2 = "Mein eigener Strauß!"
	case "EN":
		question = "Would you like to choose a bouquet or create your own?"
		button1 = "Choose!"
		button2 = "Create my own!"
	case "UK":
		question = "Хотіли б ви вибрати букет або створити власний?"
		button1 = "Вибрати!"
		button2 = "Створити свій!"
	case "RU":
		question = "Вы хотите выбрать букет или создать свой собственный?"
		button1 = "Выбрать!"
		button2 = "Создать свой!"
	}

	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(button1),
			tgbotapi.NewKeyboardButton(button2),
		),
	)
	msg := tgbotapi.NewMessage(chatID, question)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}

// Сообщение для выбора букета
func sendBouquetChoiceMessage(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message string
	switch lang {
	case "DEU":
		message = "Bitte wählen Sie einen Blumenstrauß aus unserem Telegram-Kanal aus (z. B.: #123)."
	case "EN":
		message = "Please choose a bouquet from our Telegram channel (e.g., #123)."
	case "UK":
		message = "Будь ласка, виберіть букет з нашого Telegram-каналу (наприклад: #123)."
	case "RU":
		message = "Пожалуйста, выберите букет из нашего Telegram-канала (например: #123)."
	}
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

// Сообщение для составления собственного букета
func sendCustomBouquetMessage(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message string
	switch lang {
	case "DEU":
		message = "Bitte beschreiben Sie die Zusammensetzung Ihres Straußes - Namen der Blumen und ihre Anzahl."
	case "EN":
		message = "Please describe the composition of your bouquet - flower names and their quantities."
	case "UK":
		message = "Будь ласка, опишіть склад вашого букета - назви квітів та їх кількість."
	case "RU":
		message = "Пожалуйста, опишите состав букета — названия цветов и их количество."
	}
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

// Случай нескольких букетов
func sendSingleOrMultipleQuestion(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var question, singleButton, multipleButton string
	switch lang {
	case "DEU":
		question = "Möchten Sie einen oder mehrere Sträuße bestellen?"
		singleButton = "Nur einen"
		multipleButton = "Mehrere Sträuße"
	case "EN":
		question = "Would you like to order one or multiple bouquets?"
		singleButton = "Just one"
		multipleButton = "Multiple bouquets"
	case "UK":
		question = "Хотіли б ви замовити один чи кілька букетів?"
		singleButton = "Тільки один"
		multipleButton = "Кілька букетів"
	case "RU":
		question = "Хотели бы вы заказать один или несколько букетов?"
		singleButton = "Только один"
		multipleButton = "Несколько букетов"
	}

	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(singleButton),
			tgbotapi.NewKeyboardButton(multipleButton),
		),
	)
	msg := tgbotapi.NewMessage(chatID, question)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}
func sendPackagingQuestion(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var question, craftPaper, coloredWrap, noWrap string
	switch lang {
	case "DEU":
		question = "Bitte wählen Sie eine Verpackung:"
		craftPaper = "Kraftpapier"
		coloredWrap = "Bunte Verpackung"
		noWrap = "Ohne Verpackung"
	case "EN":
		question = "Please choose a packaging:"
		craftPaper = "Craft paper"
		coloredWrap = "Colored wrap"
		noWrap = "No packaging"
	case "UK":
		question = "Оберіть, будь ласка, упаковку:"
		craftPaper = "Крафтовий папір"
		coloredWrap = "Кольорова упаковка"
		noWrap = "Упаковка не потрiбна"
	case "RU":
		question = "Выберите пожалуйста упаковку:"
		craftPaper = "Крафтовая бумага"
		coloredWrap = "Цветная упаковка"
		noWrap = "Без упаковки"
	}

	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(craftPaper),
			tgbotapi.NewKeyboardButton(coloredWrap),
			tgbotapi.NewKeyboardButton(noWrap),
		),
	)
	msg := tgbotapi.NewMessage(chatID, question)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}

// Запрос времени
func sendOrderTimeRequest(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message string
	switch lang {
	case "DEU":
		message = "Bitte geben Sie Datum und Uhrzeit ein, zu der Ihre Bestellung fertig sein soll (z. B. 2023-12-31 15:30)."
	case "EN":
		message = "Please enter the date and time by which your order should be ready (e.g., 2023-12-31 15:30)."
	case "UK":
		message = "Будь ласка, введіть дату та час, до якого ваше замовлення має бути готове (наприклад: 2023-12-31 15:30)."
	case "RU":
		message = "Введите дату и время срока, к которому должен быть готов Ваш заказ (например: 2023-12-31 15:30)."
	}
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}
func sendStoreClosedOptions(bot *tgbotapi.BotAPI, chatID int64, lang string, nextDay time.Time) {
	var message, returnButton, nextDayButton string
	switch lang {
	case "DEU":
		message = "Leider ist das Geschäft zu dieser Zeit geschlossen. Sie können Ihre Bestellung um " + nextDay.Format("2006-01-02 08:00") + " abholen oder eine andere Zeit eingeben."
		returnButton = "Zurück zur Zeitauswahl"
		nextDayButton = "Möglichst früh morgen abholen"
	case "EN":
		message = "The store is closed at this time. You can pick up your order at " + nextDay.Format("2006-01-02 08:00") + " or enter a new time."
		returnButton = "Return to time selection"
		nextDayButton = "Get as soon as possible tomorrow"
	case "UK":
		message = "Магазин закритий у цей час. Ви можете забрати замовлення о " + nextDay.Format("2006-01-02 08:00") + " або ввести інший час."
		returnButton = "Повернутися до вибору часу"
		nextDayButton = "Забрати якнайшвидше завтра"
	case "RU":
		message = "Магазин закрыт в это время. Вы можете забрать заказ в " + nextDay.Format("2006-01-02 08:00") + " или выбрать другое время."
		returnButton = "Вернуться к выбору времени"
		nextDayButton = "Получить как можно скорее завтра"
	}

	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(returnButton),
			tgbotapi.NewKeyboardButton(nextDayButton),
		),
	)
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}

// Логика обработки времени
func handleOrderTime(bot *tgbotapi.BotAPI, chatID int64, input string, lang string, userData map[int64]map[string]string) {
	loc, _ := time.LoadLocation("Europe/Vienna")
	currentTime := time.Now().In(loc)

	parsedTime, err := time.ParseInLocation("2006-01-02 15:04", input, loc)
	if err != nil {
		sendInvalidTimeMessage(bot, chatID, lang)
		return
	}

	if parsedTime.Before(currentTime) {
		sendInvalidTimeMessage(bot, chatID, lang)
		return
	}

	if !isWithinWorkingHours(parsedTime) {
		// Если магазин закрыт, предложить варианты
		nextDay := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day()+1, 8, 0, 0, 0, loc)
		sendStoreClosedOptions(bot, chatID, lang, nextDay)
		return
	}

	// Сохраняем корректное время
	userData[chatID]["time"] = parsedTime.Format("2006-01-02 15:04")
	sendOrderTimeSavedMessage(bot, chatID, lang)
}

// Сообщение о некорректном времени
func sendInvalidTimeMessage(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message string
	switch lang {
	case "DEU":
		message = "Ungültige Eingabezeit."
	case "EN":
		message = "Invalid time input."
	case "UK":
		message = "Некоректний час."
	case "RU":
		message = "Некорректный ввод времени."
	}
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

// Сообщение о закрытом магазине
func sendClosedMessage(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message string
	switch lang {
	case "DEU":
		message = "Leider ist das Geschäft zu dieser Zeit geschlossen. Bitte wählen Sie eine Zeit während der Öffnungszeiten."
	case "EN":
		message = "The store is closed at that time. Please choose a time during business hours."
	case "UK":
		message = "Магазин зачинений у цей час. Будь ласка, оберіть час у межах робочих годин."
	case "RU":
		message = "Магазин закрыт в это время. Пожалуйста, выберите время в рамках рабочего времени."
	}
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

// Сообщение о сохранении времени
func sendOrderTimeSavedMessage(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message string
	switch lang {
	case "DEU":
		message = "Ihre Bestellzeit wurde gespeichert."
	case "EN":
		message = "Your order time has been saved."
	case "UK":
		message = "Час вашого замовлення збережено."
	case "RU":
		message = "Время вашего заказа сохранено."
	}
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

// Проверка рабочего времени
func isWithinWorkingHours(t time.Time) bool {
	weekday := t.Weekday()
	hour := t.Hour()

	if weekday >= time.Monday && weekday <= time.Friday {
		return hour >= 9 && hour <= 21
	}

	if weekday == time.Saturday || weekday == time.Sunday {
		return hour >= 9 && hour <= 15
	}

	return false
}
func sendOrderConfirmation(bot *tgbotapi.BotAPI, chatID int64, lang string, userData map[int64]map[string]string) {
	order := userData[chatID]

	// Локализованные тексты для сообщения
	var confirmationMessage, nextButton string
	switch lang {
	case "DEU":
		confirmationMessage = "Bitte bestätigen Sie Ihre Bestellung:\n"
		nextButton = "Weiter!"
	case "EN":
		confirmationMessage = "Please confirm your order:\n"
		nextButton = "Next!"
	case "UK":
		confirmationMessage = "Будь ласка, підтвердіть ваше замовлення:\n"
		nextButton = "Далі!"
	case "RU":
		confirmationMessage = "Пожалуйста, подтвердите ваш заказ:\n"
		nextButton = "Далее!"
	}

	// Формируем сообщение с деталями заказа
	confirmationMessage += "🕒 Время: " + order["time"] + "\n"
	confirmationMessage += "💐 Букет: " + order["bouquet"] + "\n"
	confirmationMessage += "📦 Упаковка: " + order["packaging"] + "\n"
	confirmationMessage += "📊 Количество: " + order["quantity"] + "\n"

	// Кнопка "Далее"
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(nextButton),
		),
	)

	// Отправка сообщения
	msg := tgbotapi.NewMessage(chatID, confirmationMessage)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}
func sendPaymentMethodQuestion(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message, prepaidButton, nonPrepaidButton string
	switch lang {
	case "DEU":
		message = "Möchten Sie den Strauß mit Vorauszahlung oder ohne Vorauszahlung kaufen?"
		prepaidButton = "Mit Vorauszahlung"
		nonPrepaidButton = "Ohne Vorauszahlung"
	case "EN":
		message = "Would you like to buy the bouquet with prepayment or without prepayment?"
		prepaidButton = "With prepayment"
		nonPrepaidButton = "Without prepayment"
	case "UK":
		message = "Хочете купити букет з передоплатою чи без передоплати?"
		prepaidButton = "З передоплатою"
		nonPrepaidButton = "Без передоплати"
	case "RU":
		message = "Хотите купить букет с предоплатой или без предоплаты?"
		prepaidButton = "С предоплатой"
		nonPrepaidButton = "Без предоплаты"
	}

	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(prepaidButton),
			tgbotapi.NewKeyboardButton(nonPrepaidButton),
		),
	)

	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyMarkup = buttons
	bot.Send(msg)
}

func getApproximateDateTime(orderTime string) string {
	parsedTime, err := time.Parse("2006-01-02 15:04", orderTime)
	if err != nil {
		return orderTime // Если ошибка, возвращаем исходное время
	}
	approxTime := parsedTime.Add(3 * time.Minute) // Добавляем 3 минуты
	return approxTime.Format("2006-01-02 15:04")  // Возвращаем дату и время в формате YYYY-MM-DD HH:MM
}

func sendPrepaymentDetails(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message string
	switch lang {
	case "DEU":
		message = "Sie können die Bestellung mit diesen Angaben bezahlen: AT 1234567890. Bei Vorauszahlung wird Ihr Strauß pünktlich fertig!"
	case "EN":
		message = "You can pay for the order using these details: AT 1234567890. With prepayment, your bouquet will be ready on time!"
	case "UK":
		message = "Ви можете оплатити замовлення за цими реквізитами: AT 1234567890. При внесенні передоплати ваш букет буде готовий вчасно!"
	case "RU":
		message = "Вы можете оплатить заказ по этим реквизитам: AT 1234567890. При внесении предоплаты ваш букет будет готов точно в срок!"
	}
	bot.Send(tgbotapi.NewMessage(chatID, message))
}

func sendNoPrepaymentDetails(bot *tgbotapi.BotAPI, chatID int64, lang string, userData map[int64]map[string]string) {
	approxDateTime := getApproximateDateTime(userData[chatID]["time"])
	var message string
	switch lang {
	case "DEU":
		message = "Die endgültige Komposition des Straußes dauert noch einige Minuten nach der Zahlung in unserem Geschäft. Vielen Dank für Ihre Bestellung!\n"
		message += "⏳ Ungefähre Fertigstellungszeit nach der Zahlung: " + approxDateTime
	case "EN":
		message = "The final bouquet arrangement will take a few more minutes after payment in our store. Thank you for your order!\n"
		message += "⏳ Approximate order readiness time after payment: " + approxDateTime
	case "UK":
		message = "Остаточне складання букета займе ще кілька хвилин після оплати в нашому магазині. Дякуємо за ваше замовлення!\n"
		message += "⏳ Приблизний час готовності замовлення після оплати: " + approxDateTime
	case "RU":
		message = "Финальная компоновка букета займет еще несколько минут после оплаты в нашем магазине. Благодарим вас за сделанный заказ!\n"
		message += "⏳ Примерное время готовности заказа после оплаты: " + approxDateTime
	}
	bot.Send(tgbotapi.NewMessage(chatID, message))
}

func sendThankYouMessage(bot *tgbotapi.BotAPI, chatID int64, lang string) {
	var message string
	switch lang {
	case "DEU":
		message = "Vielen Dank, dass Sie uns gewählt haben, Ihre Bestellung wird bereits bearbeitet! ☺️ Verwenden Sie @florgalerie_bot 🤖 für eine erneute Bestellung!"
	case "EN":
		message = "Thank you for choosing us, your order is already in progress! ☺️ Use @florgalerie_bot 🤖 to reorder!"
	case "UK":
		message = "Дякуємо, що Ви обрали нас, Ваше замовлення вже в процесі приготування! ☺️ Використовуйте @florgalerie_bot 🤖 для повторного замовлення!"
	case "RU":
		message = "Благодарим, что Вы выбрали нас, Ваш заказ уже в процессе приготовления! ☺️ Используйте @florgalerie_bot 🤖 для повторного заказа!"
	}
	bot.Send(tgbotapi.NewMessage(chatID, message))
}

func sendOrderDetailsToAdmin(bot *tgbotapi.BotAPI, adminChatID int64, userChatID int64, userData map[int64]map[string]string) {
	order, exists := userData[userChatID]
	if !exists {
		log.Printf("No order data found for user: %d", userChatID)
		return
	}

	// Формируем сообщение с информацией о заказе
	message := "Новый заказ от пользователя:\n"
	message += "📞 Chat ID: " + strconv.FormatInt(userChatID, 10) + "\n"
	message += "💐 Букет: " + order["bouquet"] + "\n"
	message += "📦 Упаковка: " + order["packaging"] + "\n"
	message += "📊 Количество: " + order["quantity"] + "\n"
	message += "⏳ Время: " + order["time"] + "\n"

	// Отправляем сообщение администратору
	msg := tgbotapi.NewMessage(adminChatID, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send order details to admin: %v", err)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("7605031210:AAGTiIboCT3mxxLO6egJ3Zhkr8LAVcdu6yo")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	userLanguage := make(map[int64]string)
	userState := make(map[int64]string)
	userData := make(map[int64]map[string]string)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			// Инициализация вложенной карты для userData
			if userData[chatID] == nil {
				userData[chatID] = make(map[string]string)
			}

			switch update.Message.Text {
			case "/start":
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
				msg := tgbotapi.NewMessage(chatID, "Select a language:")
				msg.ReplyMarkup = buttons
				bot.Send(msg)
			case "DEU", "EN", "UK", "RU":
				userLanguage[chatID] = update.Message.Text
				sendGreeting(bot, chatID, update.Message.Text)
			case "Next!", "Weiter!", "Далі!", "Далее!":
				if userState[chatID] == "confirming_order" {
					userState[chatID] = "choosing_payment"
					sendPaymentMethodQuestion(bot, chatID, userLanguage[chatID])
				} else {
					sendQuestion(bot, chatID, userLanguage[chatID])
				}
			case "Choose!", "Auswählen!", "Вибрати!", "Выбрать!":
				userState[chatID] = "choosing_bouquet"
				sendBouquetChoiceMessage(bot, chatID, userLanguage[chatID])

				removeKeyboard := tgbotapi.NewRemoveKeyboard(true)
				msg := tgbotapi.NewMessage(chatID, "")
				msg.ReplyMarkup = removeKeyboard
				bot.Send(msg)
			case "Create my own!", "Mein eigener Strauß!", "Створити свій!", "Создать свой!":
				userState[chatID] = "creating_bouquet"
				sendCustomBouquetMessage(bot, chatID, userLanguage[chatID])

				// Удаляем кнопки после выбора
				removeKeyboard := tgbotapi.NewRemoveKeyboard(true)
				msg := tgbotapi.NewMessage(chatID, " ") // Здесь текст не важен, можно оставить пустую строку или пробел
				msg.ReplyMarkup = removeKeyboard
				bot.Send(msg)

			case "Just one", "Nur einen", "Тільки один", "Только один":
				userState[chatID] = "choosing_packaging"
				userData[chatID]["quantity"] = "1"
				msg := tgbotapi.NewMessage(chatID, " ")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				bot.Send(msg)
				sendPackagingQuestion(bot, chatID, userLanguage[chatID])
			case "Multiple bouquets", "Mehrere Sträuße", "Кілька букетів", "Несколько букетов":
				userState[chatID] = "choosing_packaging"
				userData[chatID]["quantity"] = "multiple"
				msg := tgbotapi.NewMessage(chatID, "Допилить случай для выбора нескольких.")
				bot.Send(msg)
				sendPackagingQuestion(bot, chatID, userLanguage[chatID])
			case "Kraftpapier", "Bunte Verpackung", "Ohne Verpackung", "Craft paper", "Colored wrap", "No packaging",
				"Крафтовий папір", "Кольорова упаковка", "Упаковка не потрiбна", "Крафтовая бумага", "Цветная упаковка", "Без упаковки":
				userState[chatID] = "waiting_for_time"
				userData[chatID]["packaging"] = update.Message.Text

				removeKeyboard := tgbotapi.NewRemoveKeyboard(true)
				msg := tgbotapi.NewMessage(chatID, " "+update.Message.Text)
				msg.ReplyMarkup = removeKeyboard
				bot.Send(msg)

				sendOrderTimeRequest(bot, chatID, userLanguage[chatID])
			default:
				switch userState[chatID] {
				case "choosing_bouquet":
					matched, _ := regexp.MatchString(`^#\d+$`, update.Message.Text)
					if matched {
						userData[chatID]["bouquet"] = update.Message.Text
						sendBouquetSavedMessage(bot, chatID, userLanguage[chatID], update.Message.Text)
						sendSingleOrMultipleQuestion(bot, chatID, userLanguage[chatID])
						userState[chatID] = "choosing_single_or_multiple"
					} else {
						bot.Send(tgbotapi.NewMessage(chatID, "Please enter a valid bouquet number (e.g., #123)."))
					}
				case "creating_bouquet":
					userData[chatID]["bouquet"] = update.Message.Text
					sendBouquetSavedMessage(bot, chatID, userLanguage[chatID], update.Message.Text)
					sendSingleOrMultipleQuestion(bot, chatID, userLanguage[chatID])
					userState[chatID] = "choosing_single_or_multiple"
				case "waiting_for_time":
					handleOrderTime(bot, chatID, update.Message.Text, userLanguage[chatID], userData)
					userData[chatID]["time"] = update.Message.Text
					sendOrderConfirmation(bot, chatID, userLanguage[chatID], userData)
					userState[chatID] = "confirming_order"
				case "Return to time selection", "Zurück zur Zeitauswahl", "Повернутися до вибору часу", "Вернуться к выбору времени":
					userState[chatID] = "waiting_for_time"
					sendOrderTimeRequest(bot, chatID, userLanguage[chatID])
				case "Get as soon as possible tomorrow", "Möglichst früh morgen abholen", "Забрати якнайшвидше завтра", "Получить как можно скорее завтра":
					nextDay := time.Now().Add(24 * time.Hour).Format("2006-01-02 08:00")
					userData[chatID]["time"] = nextDay
					sendOrderTimeSavedMessage(bot, chatID, userLanguage[chatID])

				case "choosing_payment":
					switch update.Message.Text {
					case "С предоплатой", "Mit Vorauszahlung", "З передоплатою", "With prepayment":
						sendPrepaymentDetails(bot, chatID, userLanguage[chatID])
						sendThankYouMessage(bot, chatID, userLanguage[chatID])
						adminChatID := int64(4367763577084) // Chat ID администратора
						sendOrderDetailsToAdmin(bot, adminChatID, chatID, userData)
						// Возвращаем пользователя к началу с кнопкой /start
						startButton := tgbotapi.NewReplyKeyboard(
							tgbotapi.NewKeyboardButtonRow(
								tgbotapi.NewKeyboardButton("/start"),
							),
						)
						msg := tgbotapi.NewMessage(chatID, " ")
						msg.ReplyMarkup = startButton
						bot.Send(msg)
					case "Без предоплаты", "Ohne Vorauszahlung", "Без передоплати", "Without prepayment":
						sendNoPrepaymentDetails(bot, chatID, userLanguage[chatID], userData)
						sendThankYouMessage(bot, chatID, userLanguage[chatID])
						adminChatID := int64(999999999) // Chat ID администратора
						sendOrderDetailsToAdmin(bot, adminChatID, chatID, userData)
						// Возвращаем пользователя к началу с кнопкой /start
						startButton := tgbotapi.NewReplyKeyboard(
							tgbotapi.NewKeyboardButtonRow(
								tgbotapi.NewKeyboardButton("/start"),
							),
						)
						msg := tgbotapi.NewMessage(chatID, " ")
						msg.ReplyMarkup = startButton
						bot.Send(msg)
					}
				default:
					bot.Send(tgbotapi.NewMessage(chatID, "Please select a valid option."))
				}
			}
		}
	}
}

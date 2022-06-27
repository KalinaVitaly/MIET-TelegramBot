package main

import (
	"MIET-TelegramBot/internal/app/telegramserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	configsPath string
)

func init() {
	flag.StringVar(&configsPath, "config-path", "configs/telegrambot.toml", "Path to configs")
}

func main() {
	flag.Parse()
	config := telegramserver.NewConfig()
	if _, err := toml.DecodeFile(configsPath, config); err != nil {
		log.Panic(err)
	}

	// var mapNames map[string]string = make(map[string]string)

	// mapNames["БТС"] = "BTS"
	// mapNames["Д"] = "D"
	// mapNames["ИВТ"] = "IVT"
	// mapNames["ИKT"] = "IKT"
	// mapNames["ИТD"] = "ITD"
	// mapNames["КТ"] = "KT"
	// mapNames["Л"] = "L"
	// mapNames["М"] = "M"
	// mapNames["МТ"] = "MT"
	// mapNames["НБ"] = "NB"
	// mapNames["НМ"] = "NM"
	// mapNames["П"] = "P"
	// mapNames["PИН"] = "PIN"
	// mapNames["РТ"] = "RT"
	// mapNames["ТБ"] = "TB"
	// mapNames["УК"] = "UK"
	// mapNames["УТС"] = "UTS"
	// mapNames["ЭН"] = "EN"
	// mapNames["ИБ"] = "IB"
	// files, err := ioutil.ReadDir("./resources/")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	for key, value := range mapNames {
	// 		if strings.Contains(file.Name(), key) {
	// 			fileName := file.Name()
	// 			fileName = strings.Replace(fileName, key, value, -1)
	// 			fmt.Println("Result", file.Name(), fileName, "/resources/"+file.Name())
	// 			os.Chmod("./resources/"+file.Name(), 0777)
	// 			os.Rename("./resources/"+file.Name(), "./resources/"+fileName)
	// 		}
	// 	}
	// }
	// schedule, err := datareader.ReadDataFromFile("resources\\BTS-11.json")

	// if err != nil {
	// 	fmt.Println("Error read schedule")
	// }

	// for i := range schedule.Data {
	// 	fmt.Println(schedule.Data[i].Class)
	// }

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

}

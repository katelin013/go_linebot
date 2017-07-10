package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var messages []linebot.Message
var imgary [5]linebot.ImageMessage

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if strings.Contains(message.Text, "test") {

					s := linebot.NewImageMessage("test", "test")
					fmt.Println(s)

				} else if strings.Contains(message.Text, "妹子") {
					imgAry := getPrettyUrl()

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(imgAry[0], imgAry[0]), linebot.NewImageMessage(imgAry[1], imgAry[1]), linebot.NewImageMessage(imgAry[2], imgAry[2]), linebot.NewImageMessage(imgAry[3], imgAry[3]), linebot.NewImageMessage(imgAry[4], imgAry[4])).Do(); err != nil {
						log.Print(err)
					}
				} else {

				}
			}
		}
	}
}

func botReplyMsg(s string) string {
	result := ""
	switch s {
	case "1":
		result = "a"
	case "2":
		result = "b"
	default:
		result = "我聽不懂你在說什麼小朋友!!"
	}

	return result
}

func getPrettyUrl() []string {
	var strAry []string
	var imgAry []string

	rand.Seed(time.Now().UnixNano())

	baseUrl := "https://ck101.com/forum-3581-" + strconv.Itoa(rand.Intn(1000)) + ".html"

	doc, err := goquery.NewDocument(baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("td.thumb_preview").Find("a").Each(func(i int, s *goquery.Selection) {
		text2, _ := s.Attr("href")
		strAry = append(strAry, text2)
	})
	prettyUrl := strAry[rand.Intn(len(strAry))]

	doc2, err := goquery.NewDocument(prettyUrl)
	if err != nil {
		log.Fatal(err)
	}

	doc2.Find("img").Each(func(i int, s *goquery.Selection) {
		imgText, _ := s.Attr("file")
		if len(imgText) > 0 && strings.Contains(imgText, "https") && strings.Contains(imgText, "jpg") {
			imgAry = append(imgAry, imgText)
		}
	})
	return imgAry
}

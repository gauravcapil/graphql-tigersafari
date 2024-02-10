package emailserver

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"gaurav.kapil/graphql-tigersafari/dbutils"
	"gaurav.kapil/graphql-tigersafari/graph/model"
	"gaurav.kapil/graphql-tigersafari/models"
)

var Notifier notifier
var Sender sender

func Initialize() {
	Notifier = new(TigerNotifier)
	Sender = new(TigerSender)
}

type notifier interface {
	Notify(sighting *models.MailData)
}

type TigerNotifier struct {
}

type sender interface {
	SendNotificationTo(sighting *models.MailData)
}

type TigerSender struct {
}

func getEmail(username string) string {
	result := []*model.UserData{}
	dbutils.DbConn.Where(&model.UserData{UserName: username}).First(&result)

	return result[0].Email
}

func (TigerSender) SendNotificationTo(sighting *models.MailData) {
	auth := smtp.PlainAuth("", "graphql-tigersafarikittens@gmail.com", "uhhmsaswcwpccsdg", "smtp.gmail.com")

	emailAddress := getEmail(sighting.User)

	to := []string{emailAddress}
	msg := []byte(fmt.Sprintf("To: %s\r\n"+

		"Subject: New Sighting for tiger id: %d\r\n"+

		"\r\n"+

		"The tiger was again seen at :%s,\r\n Location (lat, long): %f, %f, \r\nthe picture proof can be found at http://localhost:%s/%s", emailAddress,
		sighting.Sighting.TigerID,
		sighting.Sighting.SeenAt,
		sighting.Sighting.SeenAtLat,
		sighting.Sighting.SeenAtLon,
		os.Getenv("PORT"),
		sighting.Sighting.PhotoLocation) + "\r\n")
	log.Printf("mail body: %s", msg)
	err := smtp.SendMail("smtp.gmail.com:587", auth, "graphql-tigersafarikittens@gmail.com", to, msg)

	if err != nil {

		log.Printf("the email was not sent due to error:%s", err.Error())

	}

}

var Mailchan chan *models.MailData
var IsReady chan bool

func StartEmailServer() {
	Mailchan = make(chan *models.MailData, 100)
	IsReady = make(chan bool)
	go func() {
		for {
			IsReady <- true
			go Sender.SendNotificationTo(<-Mailchan)
		}
	}()
}

func (TigerNotifier) Notify(sighting *models.MailData) {
	Mailchan <- sighting
}

package emailserver

import (
	"fmt"
	"log"
	"net/smtp"

	"gaurav.kapil/tigerhall/dbutils"
	"gaurav.kapil/tigerhall/graph/model"
	"gaurav.kapil/tigerhall/models"
)

func getEmail(username string) string {
	result := model.UserData{}
	dbutils.DbConn.Where(&model.UserData{UserName: username}).First(&result)
	return result.Email
}

func SendNotificationTo(sighting *models.MailData) {
	auth := smtp.PlainAuth("", "tigerhallkittens@gmail.com", "uhhmsaswcwpccsdg", "smtp.gmail.com")

	emailAddress := getEmail(sighting.User)

	to := []string{emailAddress}

	msg := []byte("To: " + emailAddress + "\r\n" +
		"Subject: New Sighting Alert\r\n" +
		"\r\n" +
		fmt.Sprintf("%v", sighting.Sighting) +
		"\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "john.doe@gmail.com", to, msg)

	if err != nil {

		log.Fatal(err)

	}

}

var mailchan chan *models.MailData

func StartEmailServer() {
	mailchan = make(chan *models.MailData, 100)
	go func() {
		for {
			go SendNotificationTo(<-mailchan)
		}
	}()
}

func Notify(sighting *models.MailData) {
	mailchan <- sighting
}

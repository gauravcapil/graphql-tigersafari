package emailserver

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"gaurav.kapil/tigerhall/dbutils"
	"gaurav.kapil/tigerhall/graph/model"
	"gaurav.kapil/tigerhall/models"
)

func getEmail(username string) string {
	result := []*model.UserData{}
	dbutils.DbConn.Where(&model.UserData{UserName: username}).First(&result)

	return result[0].Email
}

func SendNotificationTo(sighting *models.MailData) {
	auth := smtp.PlainAuth("", "tigerhallkittens@gmail.com", "uhhmsaswcwpccsdg", "smtp.gmail.com")

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
	err := smtp.SendMail("smtp.gmail.com:587", auth, "tigerhallkittens@gmail.com", to, msg)

	if err != nil {

		log.Printf("the email was not sent due to error:%s", err.Error())

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

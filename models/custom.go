package models

import "gaurav.kapil/graphql-tigersafari/graph/model"

type MailData struct {
	User     string
	Sighting model.Sighting
}

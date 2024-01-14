package emailserver_test

import (
	"testing"
	"time"

	"gaurav.kapil/tigerhall/emailserver"
	"gaurav.kapil/tigerhall/models"
	"github.com/stretchr/testify/require"
)

type NotifierTest struct {
}

func (NotifierTest) Notify(sighting *models.MailData) {
	// let it be empty
}

type SenderTest struct {
}

func (SenderTest) SendNotificationTo(sighting *models.MailData) {
}

func TestServerRun(t *testing.T) {
	emailserver.Notifier = new(NotifierTest)
	emailserver.StartEmailServer()
	time.Sleep(time.Second)
	require.True(t, <-emailserver.IsReady)
}

func TestServerSend(t *testing.T) {
	emailserver.Notifier = new(emailserver.TigerNotifier)
	emailserver.Sender = new(SenderTest)
	emailserver.StartEmailServer()
	emailserver.Notifier.Notify(&models.MailData{User: "testu"})
	require.Equal(t, "testu", (<-emailserver.Mailchan).User)
}

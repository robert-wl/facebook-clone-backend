package mail

import (
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils"
)

const defaultName = "Robber"
const defaultAddress = "robber.tpaweb@gmail.com"
const defaultPassword = "-"

var name = utils.GetDotENVVariable("EMAIL_SENDER_NAME", defaultName)
var address = utils.GetDotENVVariable("EMAIL_SENDER_ADDRESS", defaultAddress)
var password = utils.GetDotENVVariable("EMAIL_SENDER_PASSWORD", defaultPassword)

func SendVerification(email string, subject string, content string) (bool, error) {

	sender := NewGmailSender(name, address, password)

	to := []string{email}

	if err := sender.SendEmail(subject, content, to, nil, nil); err != nil {
		return false, err
	}

	return true, nil
}

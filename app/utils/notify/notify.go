package notify

import (
	"errors"
	"fmt"
	"github.com/edwardzhanged/novel-go/app/model"
	"strings"
)

type Notifier interface {
	Compose() error
	Notify() error
}

type Email struct {
	To      string
	Subject string
}

func (email *Email) Compose() error {
	return nil
}

func (email *Email) Notify() error {
	return nil
}

func SendNotification(notifier Notifier, user *model.UserInfo) error {
	fmt.Printf("Sending Email %s\n", user.Username)
	err := notifier.Notify()
	if err != nil {
		return err
	}
	if strings.Contains(user.Username, "lxf") {
		return errors.New("user is not allowed to send notification")
	}
	return nil
}

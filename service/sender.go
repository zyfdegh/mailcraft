package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/zyfdegh/mailcraft/entity"
)

const (
	senderJSONPath = "./sender.json"
)

var (
	// ErrNilSender returned when struct sender is nil pointer
	ErrNilSender = errors.New("sender is nil pointer")
)

// SaveSender save sender to file
func SaveSender(sender *entity.Sender) (path string, err error) {
	if sender == nil {
		return "", ErrNilSender
	}
	data, err := json.Marshal(sender)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(senderJSONPath, data, 0644)
	if err != nil {
		return "", err
	}
	return senderJSONPath, nil
}

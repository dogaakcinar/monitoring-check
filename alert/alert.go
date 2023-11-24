package alert

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var MattermostBaseUrl = "http://localhost:8065/hooks/"

type AlertSender interface {
	SendAlert(message string) error
}

type AlertMessage struct {
	Message string `json:"message"`
}

type MattermostSender struct {
	HookID string
}

func (m *MattermostSender) SendAlert(message string) error {
	msg := AlertMessage{
		Message: message,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal Mattermost message: %v", err)
		return err
	}
	resp, err := http.Post(MattermostBaseUrl+m.HookID, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to send Mattermost alert: %v", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Printf("Failed to send Mattermost alert, status code: %d err: %v", resp.StatusCode, err)
		return err
	}
	return nil
}

func InitializeMattermostSender(HookID string) AlertSender {
	return &MattermostSender{
		HookID: HookID,
	}
}

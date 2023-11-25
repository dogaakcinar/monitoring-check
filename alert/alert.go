package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AlertSender interface {
	SendAlert(message string) error
}

type AlertMessage struct {
	Message string `json:"text"`
}

type MattermostSender struct {
	HookID    string
	ServerUrl string
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
	fmt.Println(m.ServerUrl + m.HookID)
	resp, err := http.Post(m.ServerUrl+m.HookID, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to send Mattermost alert: %v", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to send Mattermost alert, status code: %d err: %v", resp.StatusCode, err)
		return err
	}

	log.Println("Mattermost alert sent successfully")

	return nil
}

func InitializeMattermostSender(HookID string, ServerUrl string) AlertSender {
	return &MattermostSender{
		HookID:    HookID,
		ServerUrl: ServerUrl,
	}
}

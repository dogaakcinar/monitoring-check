package alert

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type MattermostMessage struct {
	ChannelID string `json:"channel_id"`
	Message   string `json:"message"`
}

func SendMattermostAlert(duration time.Duration) {
	msg := MattermostMessage{
		ChannelID: "your-channel-id",
		Message:   "No requests in the last " + duration.String(),
	}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal Mattermost message: %v", err)
		return
	}
	resp, err := http.Post("https://your-mattermost-server.com/api/v4/posts", "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Failed to send Mattermost alert: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Printf("Failed to send Mattermost alert, status code: %d", resp.StatusCode)
	}
}

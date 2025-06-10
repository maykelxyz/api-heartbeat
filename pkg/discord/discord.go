package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordEmbed struct {
	Title string `json:"title,omitempty"`
	Color int    `json:"color,omitempty"`
}

type DiscordMessage struct {
	Content string         `json:"content,omitempty"`
	Embeds  []DiscordEmbed `json:"embeds,omitempty"`
}

func GetHealthEmbed(isOK bool) DiscordEmbed {
	var color int
	var title string

	if isOK {
		color = 0x00ff00 // Green
		title = "API Health Status ✅"
	} else {
		color = 0xff0000 // Red
		title = "API Health Status ❌"
	}

	return DiscordEmbed{
		Title: title,
		Color: color,
	}
}

func GetServiceStatusEmbed() DiscordEmbed {
	return DiscordEmbed{
		Title: "Service Status 🟢",
		Color: 0x00ff00, // Green
	}
}

func SendDiscordHealthEmbedMessage(isOK bool, discordWebhookURL string) error {
	healthEmbed := GetHealthEmbed(isOK)

	payload := DiscordMessage{
		Embeds: []DiscordEmbed{healthEmbed},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
func SendDiscordServiceStatusEmbedMessage(discordWebhookURL string) error {
	serviceEmbed := GetServiceStatusEmbed()

	payload := DiscordMessage{
		Embeds: []DiscordEmbed{serviceEmbed},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

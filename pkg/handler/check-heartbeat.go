package handler

import (
	"api-heartbeat/pkg/api"
	"api-heartbeat/pkg/discord"
)

func CheckAPIHeartbeat(url string, discordWebhookURL string) error {
	isAlive := api.CheckAPIHeartbeat(url)
	if !isAlive {
		discord.SendDiscordHealthEmbedMessage(isAlive, discordWebhookURL)
	}
	return nil
}

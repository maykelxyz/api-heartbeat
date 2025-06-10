package handler

import (
	"api-heartbeat/pkg/api"
	"api-heartbeat/pkg/discord"
)

func DiscordNotify(apiURL, discordWebhookURL string) error {
	isAlive := api.CheckAPIHeartbeat(apiURL)
	discord.SendDiscordHealthEmbedMessage(isAlive, discordWebhookURL)
	discord.SendDiscordServiceStatusEmbedMessage(discordWebhookURL)
	return nil
}

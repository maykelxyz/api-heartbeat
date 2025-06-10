package executor

import (
	"api-heartbeat/pkg/config"
	"api-heartbeat/pkg/handler"
	"fmt"
)

func ExecuteJob(handlerName string, cfg config.Config) error {
	switch handlerName {
	case "CheckAPIHeartbeat":
		return handler.CheckAPIHeartbeat(cfg.APIURL, cfg.DiscordWebhookURL)
	case "DiscordNotify":
		return handler.DiscordNotify(cfg.APIURL, cfg.DiscordWebhookURL)
	default:
		return fmt.Errorf("handler not found: %s", handlerName)
	}
}

# API Heartbeat Monitor with Discord Integration

A robust Go-based monitoring system that checks the health of your APIs and sends notifications to Discord when issues are detected. The system uses a flexible job scheduling system to monitor multiple APIs with configurable frequencies.

## Features

- 🔄 Automated API health checks with configurable schedules
- 🔔 Discord notifications for API status changes
- ⏱️ Flexible job scheduling using cron expressions
- 📊 YAML-based job configuration
- 📝 Structured logging
- 🔐 Environment-based configuration
- 🏗️ Modular architecture with clean separation of concerns

## Prerequisites

- Go 1.16 or higher
- Discord webhook URL
- APIs to monitor

## Installation

1. Clone the repository:
```bash
git clone https://github.com/michaeljohnpediglorio/api-heartbeat.git
cd api-heartbeat
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory:
```env
DISCORD_WEBHOOK_URL=your_discord_webhook_url
TIMEZONE=UTC
```

## Project Structure

```
.
├── main.go           # Application entry point
├── pkg/
│   ├── api/         # API client implementations
│   ├── config/      # Configuration management
│   ├── discord/     # Discord notification handling
│   ├── executor/    # Job execution logic
│   ├── handler/     # API endpoint handlers
│   ├── jobs/        # Job configuration files
│   └── scheduler/   # Job scheduling system
```

## Configuration

Jobs are configured using YAML files in the `pkg/jobs` directory. Example job configuration:

```yaml
name: "Example API Check"
enabled: true
frequency: "*/5 * * * *"  # Every 5 minutes
handler: "http_check"
description: "Checks the health of example.com API"
```

## Usage

Run the application:

```bash
go run main.go
```

Build and run:

```bash
go build
./api-heartbeat
```

## Discord Notifications

The system will send notifications to your Discord channel when:
- An API becomes unresponsive
- An API returns an unexpected status code
- Response time exceeds the threshold
- An API recovers from a failure
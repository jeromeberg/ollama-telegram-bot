# ollama-telegram-bot

ollama-telegram-bot is a Telegram bot built in Go that interfaces with a large language model (LLM) using Ollama.

## Features

- [X] Preprompt configuration
- [X] Log history
- [X] Persistent chat context until restart

## Getting Started

### Prerequisites

#### Golang

Make sure Golang is installed on your machine. You can install it from the official [Go website](https://golang.org/dl/) or use a package manager like Homebrew: `brew install go`

#### Ollama

Make sure Ollama is installed on your machine. You can download it from the official [Ollama website](https://ollama.com/download) or use a package manager like Homebrew: `brew install ollama`

### Build the app

Clone the repository to your local machine:

```bash
git clone https://github.com/jeromeberg/ollama-telegram-bot.git
```

Build executable from sources:

```bash
go build -o bot ./src
```

### Create a Telegram Bot

1.	Open Telegram and search for the user BotFather.
2.	Start a chat with BotFather and use the `/start` command.
3.	Create a new bot by typing `/newbot` and follow the instructions to set up your bot name and username.
4.	After creation, BotFather will provide you with a token. Save this token as it will be needed to configure your bot.

### Configure the Bot

Edit the `config.json` file in the project root directory:

```json
{
    "botToken": "YOUR_TELEGRAM_BOT_TOKEN",
    "prePrompt": "You are a friendly bot. Please respond to the following queries in a concise and helpful manner:\n",
    "model": "MODEL_NAME",
    "serverUrl": "http://localhost:11434/api/generate",
    "enableLog": true,
    "startMessage": "Hello I am a friendly bot! How can I help you?"
}
```

- Replace YOUR_TELEGRAM_BOT_TOKEN with the token you received from BotFather.
- Replace MODEL_NAME with the name of the model you are using. (e.g.: `gemma2`)
- Replace `http://localhost:11434/api/generate` with the URL of your Ollama server if necessary.

### Download ollama model

Ensure that the model you specified in `config.json` is downloaded and available on your Ollama server.

```bash
ollama download <model-name>
```

Also make sure that ollama is running.

### Run the Bot

Once everything is configured, run your bot using: `./bot`

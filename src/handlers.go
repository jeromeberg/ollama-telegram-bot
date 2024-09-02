package main

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"net/http"
	"bytes"
	"encoding/json"
	"bufio"
	"log"
)

type RequestPayload struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ResponsePayload struct {
	Response   string `json:"response"`
	Done       bool   `json:"done"`
	DoneReason string `json:"done_reason,omitempty"`
	Context    []int  `json:"context,omitempty"`
}

func handlers(bot *telebot.Bot, config *Config, userContexts map[int64]*UserContext) {
	// start handler
	bot.Handle("/start", func(c telebot.Context) error {
		userID := c.Sender().ID
		if _, exists := userContexts[userID]; !exists {
			userContexts[userID] = &UserContext{UserID: userID, History: []string{}}
		}
		return c.Send(config.StartMessage)
	})

	// about handler
	bot.Handle("/about", func(c telebot.Context) error {
		about := fmt.Sprintf("Model: %s", config.Model)
		return c.Send(about)
	})

	/*
	// reset handler
	bot.Handle("/reset", func(c telebot.Context) error {
		userID := c.Sender().ID
	
		if _, exists := userContexts[userID]; exists {
			userContexts[userID].History = []string{}
		}
		return c.Send("History and context deleted.")
	})

	// help handler
	bot.Handle("/help", func(c telebot.Context) error {
		userID := c.Sender().ID
		if _, exists := userContexts[userID]; !exists {
			userContexts[userID] = &UserContext{UserID: userID, History: []string{}}
		}
		return c.Send("/reset : reset history and context\n/about : model")
	})	
	
	*/

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		userID := c.Sender().ID
		message := c.Text()

		//Log
		if config.EnableLog {
			log.Printf("[%d] %s", userID, message)
		}

		c.Notify(telebot.Typing)

		userContext, exists := userContexts[userID]
		if !exists {
			userContext = &UserContext{UserID: userID, History: []string{}}
			userContexts[userID] = userContext
		}

		userContext.History = append(userContext.History, "User: "+message)

		prompt := config.PrePrompt
		for _, msg := range userContext.History {
			prompt += msg + "\n"
		}

		payload := RequestPayload{
			Model:  config.Model,
			Prompt: prompt,
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		req, err := http.NewRequest("POST", config.ServerURL, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)

		var res string

		for scanner.Scan() {
			line := scanner.Text()

			var responsePayload ResponsePayload
			err := json.Unmarshal([]byte(line), &responsePayload)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}

			res += responsePayload.Response

			if responsePayload.Done {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
			return err
		}

		userContext.History = append(userContext.History, "Bot: "+res)

		// Log
		if config.EnableLog {
			log.Printf("[bot] %s", res)
		}

		return c.Send(res)
	})
}


package handler

import "fmt"

type CheerUpHandler struct {
}

func NewCheerUpHandler() *CheerUpHandler {
	return &CheerUpHandler{}
}

func (c *CheerUpHandler) HandleMention() error {
	fmt.Println("I have been mentioned")
	return nil
}

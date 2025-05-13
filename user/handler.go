package user

import (
	"fmt"
	"strings"
)

type CommandHandler struct {
	userService *Service
}

func NewCommandHandler(userService *Service) *CommandHandler {
	return &CommandHandler{userService: userService}
}

func (h *CommandHandler) Handle(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		fmt.Println("Please enter a command.")
		return
	}

	switch parts[0] {
	case "register":
		if len(parts) < 2 {
			fmt.Println("Usage: register <user_name>")
			return
		}
		username := parts[1]
		err := h.userService.Register(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("registered successfully\n")
		}
	default:
		fmt.Println("Unknown command:", parts[0])
	}
}

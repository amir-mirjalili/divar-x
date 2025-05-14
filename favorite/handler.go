package favorite

import (
	"fmt"
	"strings"
)

type CommandHandler struct {
	service *Service
}

func NewCommandHandler(service *Service) *CommandHandler {
	return &CommandHandler{service: service}
}

func (h *CommandHandler) Handle(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		fmt.Println("Please enter a command.")
		return
	}

	switch parts[0] {
	case "add_favorite":
		if len(parts) < 2 {
			fmt.Println("Usage: add_favorite <username> <title>")
			return
		}
		username := parts[1]
		title := parts[2]
		err := h.service.Insert(username, title)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("added successfully\n")
		}
	case "rem_favorite":
		if len(parts) < 3 {
			fmt.Println("Usage: rem_favorite <username> <title>")
			return
		}
		username := parts[1]
		title := strings.Join(parts[2:], " ")
		_, err := h.service.Delete(username, title)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Ad deleted successfully.")
		}
	case "list_favorite_advertises":
		if len(parts) < 2 {
			fmt.Println("Usage: list_favorite_advertises <username>")
			return
		}
		username := parts[1]
		data, err := h.service.GetListByUserName(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
	default:
		fmt.Println("Unknown command:", parts[0])
	}
}

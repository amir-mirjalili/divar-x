package advertise

import (
	"fmt"
	"strings"
)

type CommandHandler struct {
	adsService *Service
}

func NewCommandHandler(service *Service) *CommandHandler {
	return &CommandHandler{adsService: service}
}

func (h *CommandHandler) Handle(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		fmt.Println("Please enter a command.")
		return
	}

	switch parts[0] {
	case "add_advertise":
		if len(parts) < 2 {
			fmt.Println("Usage: add_advertise <username> <title>")
			return
		}
		username := parts[1]
		title := parts[2]
		err := h.adsService.Insert(username, title)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("posted successfully\n")
		}
	case "rem_advertise":
		if len(parts) < 3 {
			fmt.Println("Usage: rem_advertise <username> <title>")
			return
		}
		username := parts[1]
		title := strings.Join(parts[2:], " ")
		_, err := h.adsService.Delete(username, title)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Ad deleted successfully.")
		}
	case "list_my_advertises":
		if len(parts) < 2 {
			fmt.Println("Usage: list_my_advertises <username>")
			return
		}
		username := parts[1]
		data, err := h.adsService.GetListByUserName(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
	default:
		fmt.Println("Unknown command:", parts[0])
	}
}

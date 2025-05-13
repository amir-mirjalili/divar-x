package main

import (
	"bufio"
	"fmt"
	"github.com/amir-mirjalili/divar-x/advertise"
	"github.com/amir-mirjalili/divar-x/user"
	"os"
	"strings"
)

func main() {
	userRepo := user.NewInMemoryUserRepository()
	userService := user.NewUserService(userRepo)
	userCommandHandler := user.NewCommandHandler(userService)

	userChecker := user.NewUserChecker(userRepo)
	adsRepo := advertise.NewInMemoryUserRepository()
	adsService := advertise.NewAdsService(adsRepo, userChecker)
	adsCommandHandler := advertise.NewCommandHandler(adsService)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome! Type your command:")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		// Dispatch command
		if strings.HasPrefix(input, "register") {
			userCommandHandler.Handle(input)
		} else if strings.HasPrefix(input, "add_advertise") {
			adsCommandHandler.Handle(input)
		} else {
			fmt.Println("Unknown command.")
		}
	}
}

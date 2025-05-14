package main

import (
	"bufio"
	"fmt"
	"github.com/amir-mirjalili/divar-x/advertise"
	"github.com/amir-mirjalili/divar-x/favorite"
	"github.com/amir-mirjalili/divar-x/user"
	"os"
	"strings"
)

func main() {
	userRepo := user.NewInMemoryUserRepository()
	userService := user.NewUserService(userRepo)
	userCommandHandler := user.NewCommandHandler(userService)
	userChecker := user.NewUserChecker(userRepo)

	adsRepo := advertise.NewInMemoryAdRepository()
	adsService := advertise.NewAdsService(adsRepo, userChecker)
	adsCommandHandler := advertise.NewCommandHandler(adsService)
	adsChecker := advertise.NewAdsChecker(adsRepo)

	favoriteRepo := favorite.NewInMemoryFavoriteRepository()
	favoriteService := favorite.NewFavoriteService(favoriteRepo, userChecker, adsChecker)
	favoriteCommandHandler := favorite.NewCommandHandler(favoriteService)

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
		} else if strings.HasPrefix(input, "add_advertise") || strings.HasPrefix(input, "rem_advertise") || strings.HasPrefix(input, "list_my_advertises") {
			adsCommandHandler.Handle(input)
		} else if strings.HasPrefix(input, "add_favorite") || strings.HasPrefix(input, "rem_favorite") || strings.HasPrefix(input, "list_favorite_advertises") {
			favoriteCommandHandler.Handle(input)
		} else {
			fmt.Println("Unknown command.")
		}
	}
}

package main

import (
	"fmt"
	"torngo/internal/api"
	"torngo/internal/api/user"
)

func main() {
	session := api.NewSession("cl8xxMu2j7GkLdKT")

	params := &user.UserBasicIdParams{
		Id:      12,
		Comment: "Fetching basic user info",
	}
	basicResponse, err := user.GetBasicId(session, params)
	fmt.Println(basicResponse.Profile.Name, err)
}

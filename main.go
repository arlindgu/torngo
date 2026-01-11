package main

import (
	"fmt"
	"time"
	"torngo/internal/api"
	"torngo/internal/api/user"
)

func main() {
	now := time.Now().Unix()
	from := now - 60*60*24*14 // 2 weeks ago
	to := now                 // 24 hours ago

	session := api.NewSession("cl8xxMu2j7GkLdKT")
	params := user.UserAttacksParams{
		Limit:     200,
		Filter:    user.UserAttacksIncoming,
		Timestamp: api.ApiTimestamp(now),
		Sort:      api.Ascending,
		From:      api.ApiFrom(from),
		To:        api.ApiTo(to),
		Comment:   "Getting user attacks",
	}

	response, err := user.GetAttacks(session, &params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User Attacks:", response.Attacks)
}

package main

import (
	"fmt"
	"time"
	"torngo/internal/api"
	"torngo/internal/api/user"
)

func main() {

	var ts int64 = time.Now().Unix()

	//TornGO Wrapper
	// Create URL
	respurl, urlerr := user.CreateAmmoURL(&user.UserAmmoParams{
		BaseParams: api.BaseParams{
			APIKey:    "▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓",
			Timestamp: int(ts),
			Comment:   "TornGO AmmoRequest",
		},
	})
	if urlerr != nil {
		panic(urlerr)
	}
	fmt.Println("Generated URL:", respurl)

	//Skip URLCreation and do GET Request in one function
	ammoResp, resperr := user.GetAmmo(&user.UserAmmoParams{
		BaseParams: api.BaseParams{
			APIKey:    "▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓",
			Timestamp: int(ts),
			Comment:   "TornGO AmmoRequest",
		},
	})
	if resperr != nil {
		panic(resperr)
	}
	fmt.Printf("Ammo Response: %+v\n", ammoResp.Ammo)
}

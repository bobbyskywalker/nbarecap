/*
Copyright © 2026 Aleksander Garbacz <aleksander.garbacz@gmail.com>
*/
package main

import (
	"fmt"
	"nbarecap/pkg/nba_api/clients"
)

//func main() {
//	cmd.Execute()
//}

func main() {
	client := clients.NewNbaApiClient()
	response, err := client.FetchScoreboardV2("2021-09-20")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
}

/*
Copyright © 2026 Aleksander Garbacz <aleksander.garbacz@gmail.com>
*/
package main

import (
	"encoding/json"
	"fmt"
	"nbarecap/pkg/nba_api/clients"
	"nbarecap/pkg/nba_api/models"
)

//func main() {
//	cmd.Execute()
//}

// TODO: refactor and move to proper method in nba package
func main() {
	client := clients.NewNbaApiClient()
	bxResponseBody, err := client.FetchBoxScoreTraditionalV3Json("0022500558")
	if err != nil {
		panic(err)
	}
	var resp models.BoxScoreTraditionalV3Response
	err1 := json.Unmarshal(bxResponseBody, &resp)
	if err1 != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")

	fmt.Println(string(b))
}

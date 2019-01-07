package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xplodwild/realmdefensecheat/realmdefense"
	"os"
)

func main() {
	userAgent := flag.String("useragent", "Dalvik/2.1.0 (Linux; U; Android 8.1.0; ONEPLUS A3003 Build/OPM4.171019.021.Y1)", "The user agent string")

	cli := realmdefense.NewClient(realmdefense.EndpointBabeltimeUS, *userAgent)

	// Get our cloud save data
	res, err := cli.POST(realmdefense.ApiLoadSave, toJson(realmdefense.LoadSaveRequest{
		Data: "",
		Id:   os.Args[1],
		Seq:  1,
	}), false, true)

	if err != nil {
		fmt.Printf("failed to load cloud save: %s\n", err)
		os.Exit(1)
	}

	// Decode the response
	var responseData realmdefense.MessageLoadSaveResponse
	err = json.Unmarshal(res, &responseData)
	if err != nil {
		fmt.Printf("failed to unmarshal loadsave response: %s\n", err)
		fmt.Printf("raw data: %x\n", res)
		os.Exit(1)
	}

	// Decode the game data
	var gameData realmdefense.SaveData
	err = json.Unmarshal([]byte(responseData.Data), &gameData)
	if err != nil {
		fmt.Printf("failed to unmarshal game data: %s\n", err)
		os.Exit(1)
	}

	// Print out the amount of gems we currently have
	gems := realmdefense.GetValueFromIVEntry(gameData.Iv.G)
	fmt.Printf("Amount of gems: %d\n", gems)

	// Add ourselves 1000 gems
	gems += 1000
	gameData.Iv.G = realmdefense.MakeIVEntryValue(gems)

	// Update the sequence number and ET in the save
	gameData.Seq++

	fmt.Printf("New amount of gems: %d\n", gems)

	// Send that as our new save
	cli.POST(realmdefense.ApiSave, toJson(realmdefense.SaveRequest{
		Data: string(toJson(gameData)),
		Id:   gameData.Uid,
		Seq:  gameData.Seq,
	}), false, false)

	fmt.Printf("All done! You can restore now!\n")
}

func toJson(i interface{}) []byte {
	b, _ := json.Marshal(i)
	return b
}

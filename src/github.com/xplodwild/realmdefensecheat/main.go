package main

import (
	"encoding/json"
	"flag"
	"github.com/xplodwild/realmdefensecheat/realmdefense"
	"gopkg.in/abiosoft/ishell.v2"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	setupShell()
	/*
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
		var responseData realmdefense.LoadSaveResponse
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

		fmt.Printf("All done! You can restore now!\n")*/
}

func setupShell() {
	shell := ishell.New()

	shell.Println("<< Realm Defense Cheat Tool >>")
	shell.Println("A fully-featured cheating tool for Babeltime's Realm Defense mobile game.")

	// Read user ID from command line, or ask to type it in
	userAgent := flag.String("useragent", "Dalvik/2.1.0 (Linux; U; Android 8.1.0; ONEPLUS A3003 Build/OPM4.171019.021.Y1)", "The user agent string")
	userIdPtr := flag.String("userid", "", "The user ID")
	flag.Parse()

	var userId string

	if *userIdPtr == "" {
		// Read out the UUID first
		shell.Println("First of all, please type in your User ID. It can be found when pressing")
		shell.Println("the Contact button in the game's settings. It looks like this: abcdef12-1234-ab12-abcd-12345678abcd")
		shell.Print("User ID: ")
		userId = shell.ReadLine()

		if len(userId) != 36 {
			shell.Println("Your user ID should be 36 characters long. Please double-check what you've typed.")
			setupShell()
			return
		}
	} else {
		userId = *userIdPtr

		if len(userId) != 36 {
			shell.Println("Your user ID should be 36 characters long. Please double-check what you've typed.")
			os.Exit(1)
			return
		}
	}

	// Save warning
	shell.Println("**********************")
	shell.Println("******** STOP ********")
	shell.Println("**********************")
	shell.Println("Have you saved your game and exited it? Before doing anything in this app, open the game")
	shell.Println("on your phone, go to the Settings button on the top right corner, press the Backup & Restore")
	shell.Println("button, then Backup. Once the button goes gray, close the game by swiping it off your")
	shell.Println("list of currently open apps on your phone.")
	shell.Println("Press Enter to continue.")

	shell.ReadLine()

	shell.Println("Type \"help\" to get a list of the commands.")

	// Create a global game client
	cli := realmdefense.NewClient(realmdefense.EndpointBabeltimeUS, *userAgent)

	// Register the shell commands
	shell.AddCmd(&ishell.Cmd{
		Name: "backup",
		Help: "Backs up your current game save into the specified file name",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Usage: backup <filename>")
				shell.Println("Example: backup my-save-2019-01-01.json")
				return
			}

			shell.Println("Downloading game save...")
			res, err := cli.LoadSave(realmdefense.LoadSaveRequest{
				Data: "",
				Id:   userId,
				Seq:  1,
			})

			if err != nil {
				shell.Printf("Error while loading game save: %s\n", err)
				return
			}

			shell.Printf("Writing game save to %s...\n", c.Args[0])
			err = ioutil.WriteFile(c.Args[0], []byte(res.Data), 0644)
			if err != nil {
				shell.Printf("Failed to write file: %s\n", err)
				return
			}

			shell.Printf("Game saved to %s!\n", c.Args[0])
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "backup",
		Help: "Backs up your current game save into a file",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Backs up your current game state into the specified filename.")
				shell.Println("Usage: backup <filename>")
				shell.Println("Example: backup my-save-2019-01-01.json")
				return
			}

			shell.Println("Downloading game save...")
			res, err := cli.LoadSave(realmdefense.LoadSaveRequest{
				Data: "",
				Id:   userId,
				Seq:  1,
			})

			if err != nil {
				shell.Printf("Error while loading game save: %s\n", err)
				return
			}

			shell.Printf("Writing game save to %s...\n", c.Args[0])
			err = ioutil.WriteFile(c.Args[0], []byte(res.Data), 0644)
			if err != nil {
				shell.Printf("Failed to write file: %s\n", err)
				return
			}

			shell.Printf("Game saved to %s!\n", c.Args[0])
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "restore",
		Help: "Restores the saved game to the server",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Restores the saved game file to the server.")
				shell.Println("Warning: If you restore bad data, you'll only have bad data and no way to recover anything. Make sure you restore the file you actually want to restore!")
				shell.Println("Warning: The player ID will be kept as-is. If you restore your old save into a new account, make sure to change the \"Uid\" field to your new player ID!")
				shell.Println("Usage: restore <filename>")
				shell.Println("Example: restore my-save-2019-01-01.json")
				return
			}

			// Download the current game save. We'll need it to know the new seq number.
			shell.Println("Downloading current game save...")
			res, err := cli.LoadSave(realmdefense.LoadSaveRequest{
				Data: "",
				Id:   userId,
				Seq:  1,
			})

			if err != nil {
				shell.Printf("Error while loading existing game save: %s\n", err)
				return
			}

			existingSaveData, err := decodeGameData(res.Data)
			if err != nil {
				shell.Println("Failed to decode existing game data:", err)
				return
			}

			// Read and decode the saved file
			saveDataBytes, err := ioutil.ReadFile(c.Args[0])
			if err != nil {
				shell.Println("Failed to read backup file:", err)
				return
			}

			saveData, err := decodeGameData(string(saveDataBytes))
			if err != nil {
				shell.Println("Failed to decode game data from save:", err)
				return
			}

			// Update the sequence number
			saveData.Seq = existingSaveData.Seq + 1
			newSaveDataBytes := realmdefense.ToJson(saveData)

			// Send the save to the server
			err = cli.Save(realmdefense.SaveRequest{
				Data: string(newSaveDataBytes),
				Seq:  saveData.Seq,
				Id:   userId,
			})
			if err != nil {
				shell.Println("Failed to send new save:", err)
				return
			}

			shell.Println("Save sent to server!")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "add-gems",
		Help: "Adds the specified amount of gems to your account",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Adds the specified amount of gems to your account.")
				shell.Println("Usage: add-gems <amount>")
				shell.Println("Example: add-gems 1000")
				return
			}

			amount, err := strconv.Atoi(c.Args[0])
			if err != nil {
				shell.Println(c.Args[0], "is not a valid numerical value")
				return
			}

			// Load our game save
			shell.Println("Loading game data...")
			res, err := cli.LoadSave(realmdefense.LoadSaveRequest{
				Data: "",
				Id:   userId,
				Seq:  1,
			})

			if err != nil {
				shell.Println("Failed to load game data:", err)
				return
			}

			gameData, err := decodeGameData(res.Data)
			if err != nil {
				shell.Println("Failed to decode game data:", err)
				return
			}

			// Change the amount of gems
			gems := realmdefense.GetValueFromIVEntry(gameData.Iv.G)
			shell.Printf("Current amount of gems: %d\n", gems)
			shell.Printf("New amount of gems: %d\n", gems+amount)

			gameData.Iv.G = realmdefense.MakeIVEntryValue(gems + amount)

			// Don't forget to update the sequence number
			gameData.Seq++

			shell.Println("Sending new game save...")
			err = cli.Save(realmdefense.SaveRequest{
				Data: string(realmdefense.ToJson(gameData)),
				Id:   gameData.Uid,
				Seq:  gameData.Seq,
			})

			if err != nil {
				shell.Println("Error while sending game save:", err)
				return
			}

			shell.Println("Game saved!")
		},
	})

	// Let's go!
	shell.Run()
}

func decodeGameData(dataStr string) (realmdefense.SaveData, error) {
	// Decode the game data
	var gameData realmdefense.SaveData
	err := json.Unmarshal([]byte(dataStr), &gameData)
	return gameData, err
}

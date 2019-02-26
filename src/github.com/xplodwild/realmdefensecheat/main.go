package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xplodwild/realmdefensecheat/realmdefense"
	"gopkg.in/abiosoft/ishell.v2"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	setupShell()
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
		Name: "decrypt",
		Help: "Decrypts the specified filepath into filepath.dec",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Usage: decrypt <filepath>")
				return
			}

			file, err := ioutil.ReadFile(c.Args[0])
			if err != nil {
				shell.Printf("Failed to read %s: %s\n", c.Args[0], err)
				return
			}

			err = ioutil.WriteFile(c.Args[0]+".dec", realmdefense.DecryptCFB(file), 0644)
			if err != nil {
				shell.Printf("Failed to write: %s\n", err)
				return
			}

			shell.Println("Decrypted")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "recover",
		Help: "(For Debugging Only) Stores the recovered data for the specified social ID in recover.json",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Usage: recover <socialID>")
				return
			}

			res, err := cli.Recover(realmdefense.RecoverRequest{
				SocialId: c.Args[0],
				Platform: "PlayGames",
			})
			if err != nil {
				shell.Println(err)
				return
			}

			ioutil.WriteFile("recover.json", []byte(res.Data[0]), 0644)
			shell.Println("Done")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "unban",
		Help: "Unbans you from the tournament",
		Func: func(c *ishell.Context) {
			shell.Println("*****************************")
			shell.Println("This will initialize and create you a new player ID, which will reset your")
			shell.Println("tournament profile, while restoring your current game save.")
			shell.Println("*****************************")
			shell.Println("Press Enter to continue.")

			shell.ReadLine()

			// First, load the existing save to push back the progress
			shell.Println("Downloading current game save...")
			res, err := cli.LoadSave(realmdefense.LoadSaveRequest{
				Data: "",
				Id:   userId,
				Seq:  1,
			})

			if err != nil {
				shell.Printf("Error while loading game save: %s\n", err)
				return
			}

			existingSaveData, err := decodeGameData(res.Data)
			if err != nil {
				shell.Println("Failed to decode existing game data:", err)
				return
			}

			// Strip out the existing user ID and social ID so that the server generates a new one. Reset sequence
			// to 0 as well, since it's a new save.
			previousSaveSeq := existingSaveData.Seq
			previousSocialId := existingSaveData.Sid
			existingSaveData.Uid = ""
			existingSaveData.Sid = ""
			existingSaveData.Seq = 0

			shell.Println("Creating a new save...")
			newSaveDataBytes := realmdefense.ToJson(existingSaveData)

			newSave, err := cli.CreateSave(realmdefense.CreateSaveRequest{
				Seq:  0,
				Data: string(newSaveDataBytes),
			})

			if err != nil {
				shell.Println("Failed to create save:", err)
				return
			}

			shell.Println("New save created! Your new player ID:", newSave.Id)

			// And finally, bind back our social account to it
			shell.Println("Binding back account to ", previousSocialId, "...")
			err = cli.Bind(realmdefense.BindRequest{
				Id:          newSave.Id,
				Platform:    "PlayGames",
				SocialAlias: fmt.Sprintf("%d", time.Now().Unix()),
				SocialId:    previousSocialId,
			})

			if err != nil {
				shell.Println("Failed to bind social account:", err)
				return
			}

			shell.Println("Social account bound, pushing this new save as the primary one")
			existingSaveData.Seq = previousSaveSeq + 5
			existingSaveData.Uid = newSave.Id
			existingSaveData.Sid = previousSocialId
			newSaveDataBytes = realmdefense.ToJson(existingSaveData)

			// Send the save to the server
			err = cli.Save(realmdefense.SaveRequest{
				Data: string(newSaveDataBytes),
				Seq:  existingSaveData.Seq,
				Id:   newSave.Id,
			})
			if err != nil {
				shell.Println("Failed to send new save:", err)
				return
			}

			shell.Printf("Saved with Uid=%s, Seq=%d\n", existingSaveData.Uid, existingSaveData.Seq)

			shell.Println("Successfully created and bound back your social account! Open the game, and then")
			shell.Println("go to Settings, Backup & Restore, and press the RESTORE button this time. The game")
			shell.Println("will tell you that 'some progress might be lost', press OK. You will then have a")
			shell.Println("fresh start at the tournament.")
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
				shell.Println("If the amount is 0, this will only load and show your gems amount without touching your save.")
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
			gameData.Iv.G = realmdefense.MakeIVEntryValue(gems + amount)

			shell.Printf("Current amount of gems: %d\n", gems)

			if amount != 0 {
				shell.Printf("New amount of gems: %d\n", gems+amount)

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
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "add-elixir",
		Help: "Adds the specified amount of elixir to your account",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Adds the specified amount of elixir to your account.")
				shell.Println("If the amount is 0, this will only load and show your elixir amount without touching your save.")
				shell.Println("Usage: add-elixir <amount>")
				shell.Println("Example: add-elixir 1000")
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

			// Change the amount of elixir
			elixir := realmdefense.GetValueFromIVEntry(gameData.Iv.X)
			gameData.Iv.X = realmdefense.MakeIVEntryValue(elixir + amount)
			shell.Printf("Current amount of elixir: %d\n", elixir)

			if amount != 0 {
				shell.Printf("New amount of elixir: %d\n", elixir+amount)

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
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "add-tokens",
		Help: "Adds the specified amount of awakening tokens to your account",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 2 {
				shell.Println("Adds the specified amount of awakening tokens to your account.")
				shell.Println("If the amount is 0, this will only load and show your hero's tokens without touching your save.")
				shell.Println("Usage: add-tokens <hero> <amount>")
				shell.Println("List of heroes: bolton, caldera, connie, efrigid, fee, helios, hogan, kaguya (azura), lancelot, leif, mabyn, masamune, narlax, obsidian, raida, sethos, smoulder, yan")
				shell.Println("Example: add-tokens kaguya 80")
				return
			}

			amount, err := strconv.Atoi(c.Args[1])
			if err != nil {
				shell.Println(c.Args[0], "is not a valid numerical value")
				return
			}

			hero := c.Args[0]

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

			// Change the amount of tokens
			origTokens := 0

			switch hero {
			case "bolton":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tbolton)
				gameData.Iv.Tbolton = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "caldera":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tcaldera)
				gameData.Iv.Tcaldera = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "connie":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tconnie)
				gameData.Iv.Tconnie = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "efrigid":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tefrigid)
				gameData.Iv.Tefrigid = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "fee":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tfee)
				gameData.Iv.Tfee = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "helios":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Thelios)
				gameData.Iv.Thelios = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "hogan":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Thogan)
				gameData.Iv.Thogan = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "kaguya":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tkaguya)
				gameData.Iv.Tkaguya = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "lancelot":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tlancelot)
				gameData.Iv.Tlancelot = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "leif":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tleif)
				gameData.Iv.Tleif = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "mabyn":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tmabyn)
				gameData.Iv.Tmabyn = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "masamune":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tmasamune)
				gameData.Iv.Tmasamune = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "narlax":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tnarlax)
				gameData.Iv.Tnarlax = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "obsidian":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tobsidian)
				gameData.Iv.Tobsidian = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "raida":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Traida)
				gameData.Iv.Traida = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "sethos":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tsethos)
				gameData.Iv.Tsethos = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "smoulder":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tsmoulder)
				gameData.Iv.Tsmoulder = realmdefense.MakeIVEntryValue(origTokens + amount)
			case "yan":
				origTokens = realmdefense.GetValueFromIVEntry(gameData.Iv.Tyan)
				gameData.Iv.Tyan = realmdefense.MakeIVEntryValue(origTokens + amount)
			}

			shell.Printf("Current amount of %s tokens: %d\n", hero, origTokens)

			if amount != 0 {
				shell.Printf("New amount of %s tokens: %d\n", hero, origTokens+amount)

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
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "tournament",
		Help: "Sends the specified score as tournament",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 2 {
				shell.Println("Sends the specified score as your tournament score.")
				shell.Println("Make sure it's legit, otherwise you're in big troubles. Heroes will be asked while running the command.")
				shell.Println("Usage: tournament <kills> <time in MILLIseconds>")
				shell.Println("Example: tournament 304 242871")
				return
			}

			kills, err := strconv.ParseUint(c.Args[0], 10, 64)
			if err != nil {
				shell.Println(c.Args[0], "is not a valid numerical value")
				return
			}

			millis, err := strconv.ParseUint(c.Args[1], 10, 64)
			if err != nil {
				shell.Println(c.Args[1], "is not a valid numerical value")
				return
			}

			// Load the user data (for country code, nickname, etc)
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

			// Query the heroes to use
			heroesNames := []string{
				"", "Hogan", "Fee", "Obsidian", "Masamune", "Lancelot", "Bolton", "Smoulder", "Connie",
				"Efrigid", "Helios", "Sethos", "Mabyn", "Yan", "Narlax", "Leif", "Caldera", "Azura", "Raida",
			}
			var heroes []int
			heroesLevels := []int{35, 35, 35}
			heroesRanks := []int{6, 6, 6}

			for heroes == nil || len(heroes) != 3 {
				heroes = shell.Checklist(heroesNames, "Select 3 heroes", []int{})
			}

			// Ask the heroes levels
			for i := 0; i < 3; i++ {
				shell.Println("Please enter", heroesNames[heroes[i]], "level [1-35]")
				heroesLevels[i], err = strconv.Atoi(shell.ReadLine())

				if err != nil || heroesLevels[i] < 1 || heroesLevels[i] > 35 {
					shell.Println("Invalid value")
					i--
				}
			}

			// Ask the heroes ranks
			for i := 0; i < 3; i++ {
				shell.Println("Please enter", heroesNames[heroes[i]], "rank [1-6]")
				heroesRanks[i], err = strconv.Atoi(shell.ReadLine())

				if err != nil || heroesRanks[i] < 1 || heroesRanks[i] > 6 {
					shell.Println("Invalid value")
					i--
				}
			}

			// Load the tournament data
			shell.Println("Loading tournament data...")
			tnQuery, err := cli.TournamentQuery(realmdefense.TournamentQuery{
				Uid:   userId,
				Realm: "Android",
			})

			if err != nil {
				shell.Println("Failed to load tournament data:", err)
				return
			}

			if tnQuery.Banned {
				shell.Println("It looks like you are banned from the tournament!")
				return
			}

			if tnQuery.Status != 1 {
				shell.Println("It doesn't look like the tournament has opened yet")
				return
			}

			shell.Println("Tournament ID:", tnQuery.Tid)
			shell.Println("Current league:", tnQuery.League)

			duration, _ := time.ParseDuration(fmt.Sprintf("%dms", millis))
			shell.Println("Posting score:", kills, " kills in ", duration)

			err = cli.TournamentScore(realmdefense.TournamentScoreRequest{
				Uid: userId,
				Data: string(realmdefense.ToJson(realmdefense.TournamentScoreData{
					CountryCode: gameData.Cc,
					Duration:    millis,
					KillCount:   int(kills),
					Heroes:      heroes,
					Levels:      heroesLevels,
					Ranks:       heroesRanks,
					Skins:       []string{"", "", ""},
					Username:    gameData.Sal,
					W17Bonus:    false,
					W17Stars:    0,
				})),
				Realm: "Android",
				Power: 242871,
				Score: realmdefense.CalculateTournamentScore(kills, millis),
				Tid:   tnQuery.Tid,
			})
			if err != nil {
				shell.Println("Failed to send tournament score:", err)
				return
			}

			shell.Println("Tournament score sent! Hope you win :)")
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

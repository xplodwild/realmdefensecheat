# This tool is currently being updated to the latest protocol version, see issue #25 for more details.

If you wish to donate to allow me to spend more time on this, [click here](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=CKZV4JWZ8BNUY&currency_code=EUR&source=url).

# Realm Defense Cheat Tool

This is a (legit) cheating tool for Babeltime's Realm Defense game, a mobile game 
available on both Android and iOS. This gives you a command-line interface to add yourself
gems, elixir, awakening tokens, set a custom tournament score, get unbanned, etc.

For more technical information on how this came to be, you can [read my Medium article
about how I found out the hash computation method](https://medium.com/@xplodwild/turning-the-frustration-of-a-mobile-game-into-a-reverse-engineering-training-a9887043efdf).

**RUN THE "backup" COMMAND BEFORE DOING ANYTHING! If anything happens to your save
using this tool, it is the only way to recover your progress if it goes wrong!**

## Downloading and running

You will find a prebuilt version for Windows in the Release section.

**Before running any command, start by Saving your game. Go into the game's settings,
press the "Backup & Restore" button, then "Backup". Once the button go grey, close the
game by swiping it off your recent apps menu. Then, you can run the command below. NOT DOING
A SAVE MAY MAKE YOU LOOSE RECENT PROGRESS, OR ISSUES WITH YOUR ENTIRE PROGRESS, SO BE
CAREFUL TO SAVE BEFORE RUNNING THE COMMAND!**

Double-click on the exe to get an interactive prompt. Sorry, no shiny user interface!
Follow the instructions on screen to enter your User ID, make sure your game data
is backed-up to cloud, then type ` help` to get information on the commands available:

* `backup` saves your current game into the specified filename. It is very useful to do a backup
  before running any other command, in case anything goes wrong. Also, this let you restore your
  progress to your account or to another account, as well as let you edit any value you want in
  your game save. You can then apply it by running the `restore` command. **I strongly
  recommend you to run this command first to avoid any problem.**
* `restore` restores your saved game file to the server (that has been obtained through the
  `backup` command). Note that the values are left as-is, so if you're restoring your save to
  a new account, make sure to update the "Uid" field in the JSON file to your new account ID.
* `add-gems` adds the specified amount of gems to your account. Example: `add-gems 1000` 
* `add-elixir` adds the specified amount of elixir to your account. Example: `add-elixir 1000`
* `add-tokens` adds the specified amount of tokens to your hero. Example: `add-tokens bolton 80`
* `tournament` lets you send a custom tournament score to the server. Start by thinking how many
  kills you want, and in how much time (in **milliseconds**), then run `tournament <kills> <duration>`
  for example: `tournament 500 344108` (344108 = 5:44.108). Then, use the arrow keys and space
  bar to select the 3 heroes you want to use, then confirm using Return once you have 3 heroes.
  Then, for each hero, enter their level. Once that's done, enter their awakening rank. And 
  once that's done, your score is sent to the server and in the ladder :) 
* `unban` lets you get unbanned from the tournament. This will create you a completely new
  account, while keeping your game progress untouched. You will be sent back to Bronze league,
  but you will be able to participate again (no more "stuck at 50th"). You will need to reset the
  game data on your phone or tablet before running this command, so that the new profile gets
  applied when you restart the game. If you don't know how to do it: just uninstall then reinstall the
  game after running the command.

## Building the software

These instructions will get you a copy of the project up and running on your local
machine for development and testing purposes.

### Prerequisites

You simply need the Go(lang) SDK and godep tool. Then, set the repository's root as your
GOPATH, go into src/github.com/xplodwild/realmdefensecheat, then run `dep ensure` to get
dependencies.

### Running
```
go run src/github.com/xplodwild/realmdefensecheat/main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Final notes

* This tool should run quite undetected, since it uses the game's regular protocol. It
  contains a default User-Agent for my own phone, but you can set your own phone's
  User-Agent by using the --useragent option. 
* Despite this, I'm not responsible if anything happens to your profile (get banned,
  broken save, or anything). Just don't give yourself too many gems too quicky I guess.
  If anything ever goes wrong, keep your backup file safe (you did one, right?) and open an issue
  so that I can investigate. Generally, errors happen when new fields were added that aren't
  saved by the cheat. Once added, you will be able to restore your save and continue.
* If you get banned, use the "unban" command to reset your tournament score.
* Yes, you can modify the data of random people. Just take their user IDs from
  Facebook events comments.
* If anyone wants to complete the api.go file with some fields descriptions, feel free.

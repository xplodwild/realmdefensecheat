# Realm Defense Cheat Tool

This is a (legit) cheating tool for Babeltime's Realm Defense game, a mobile game 
available on both Android and iOS. By default, running this with your user account
will give you 1000 more gems. You can tune the source code to modify everything else,
such as elixir, awakening tokens, and more.

For more technical information on how this came to be, you can [read my Medium article
about how I found out the hash computation method](https://medium.com/@xplodwild/turning-the-frustration-of-a-mobile-game-into-a-reverse-engineering-training-a9887043efdf).

When I have a bit more time, this will evolve into a fully-featured save editor, as
well as a tournament custom score tool.

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

* `add-gems` adds the specified amount of gems to your account. Example: `add-gems 1000` 
* `backup` saves your current game into the specified filename. It is very useful to do a backup
  before running any other command, in case anything goes wrong. Also, this let you restore your
  progress to your account or to another account, as well as let you edit any value you want in
  your game save. You can then apply it by running the `restore` command.
* `restore` restores your saved game file to the server (that has been obtained through the
  `backup` command). Note that the values are left as-is, so if you're restoring your save to
  a new account, make sure to update the "Uid" field in the JSON file to your new account ID.

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
or anything). Just don't give yourself too many gems too quicky I guess. If anything
ever goes wrong, the tool will soon be able to backup and restore a full game save
to a different account, so that if you get banned, you can recover your progress on
a new account.
* Yes, you can modify the dat arandom people. Just take their user IDs from
Facebook events comments.
* If anyone wants to complete the api.go file with some fields descriptions, feel free.

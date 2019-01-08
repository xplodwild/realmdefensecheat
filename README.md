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

You will find a prebuilt version for Windows in the Release section. Open a command
prompt inside the folder where you've downloaded the .exe.

**Before running the command, start by Saving your game. Go into the game's settings,
press the "Backup & Restore" button, then "Backup". Once the button go grey, close the
game by swiping it off your recents app. Then, you can run the command below. NOT DOING
A SAVE MAY MAKE YOU LOOSE RECENT PROGRESS, OR ISSUES WITH YOUR ENTIRE PROGRESS, SO BE
CAREFUL TO SAVE BEFORE RUNNING THE COMMAND!**

Once your game is saved, run:

```
RealmDefenseCheatTool.exe YOUR_USER_ID
```

Replace `YOUR_USER_ID` with your actual account ID. This can be found by going into
the game's settings, press "Contact", and copy the "ID" contained in the message. It
looks like "abcdef12-1234-ab12-abcd-12345678abcd".

Once the program did its things, open the game again. Your gem counter should now have
increased by 1000.


## Building the software

These instructions will get you a copy of the project up and running on your local
machine for development and testing purposes.

### Prerequisites & Running

You simply need the Go(lang) SDK. Then, set the repository's root as your GOPATH, and
run the main.go file:

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
* Yes, you can also increase the gems of random people. Just take their user IDs from
Facebook events comments.
* If anyone wants to complete the api.go file with some fields descriptions, feel free.

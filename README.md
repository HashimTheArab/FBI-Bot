# FBI-Bot
A multi-purpose, very efficient discord bot written in Go.

# Running the bot
1. Create a bot account at the [Discord Developer Portal](https://discord.com/developers/applications)
2. Invite the bot account to a server
3. Download a release at [releases](https://github.com/Prim69/FBI-Bot/tags)
4. Run the exe and follow the setup
5. The setup will ask for your bots token and create a token.tok file, if you ever need to change it simply edit that file or delete it and re-run the setup

# Features
 - General Commands
   - [x] Help
     - Lists all commands
     - Supports specific commands "help \<command>"
     - Automatic system, command gets added to the help commmand when registered to the command map
 - Fun Commands
   - [x] Snipe
     - Snipes the last deleted message
     - Supports multiple messages
   - [x] Editsnipe
     - Shows changes made to the last edited message
     - Supports multiple messages
 - Music Commands
   - [ ] Play
      - Plays a song
 - User Commands
   - [x] Avatar
     - Works with ID and mentions
     - Works with users not in the server (ID)
 - Bot Commands
   - [x] Ping
     - Displays both API and regular latency
   - [x] Stats
     - Displays uptime and other bot information
 - Utility Commands
   - [x] Nuke
     - Nukes the current channel
     - Keeps channel permissions, and all other stuff
   - [x] Purge
     - Deletes an amount of messages
     - Instant
   - [ ] Backup
     - Saves a server template
   - [ ] Load
     - Loads a server template
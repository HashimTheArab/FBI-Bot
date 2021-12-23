# FBI-Bot
A multi-purpose, very efficient discord bot written in Go.

# Running the bot
1. Create a bot account at the [Discord Developer Portal](https://discord.com/developers/applications)
2. Invite the bot account to a server
3. Download a release at [releases](https://github.com/Prim69/FBI-Bot/tags)
4. Run the exe and follow the setup
5. The setup will ask for your bots token and save it in a file called settings.json, if you ever need to change it simply edit that file

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
 - Web Commands
   - [x] Ask
     - Fetches answer for a question from WolframAlpha
   - [x] Lookup
     - Fetches information on an Xbox account
   - [ ] Urban
     - Fetches a definition from the Urban Dictionary
 - Minecraft Commands
   - [x] Query
     - Query a minecraft server to receive information
     - Both long and short (ping) queries
 - Music Commands
   - [ ] Play
      - Plays a song
 - User Commands
   - [x] Avatar
     - Works with ID and mentions
     - Works with users not in the server (ID)
   - [x] Whois
     - Fetches information about a user
     - Works with ID, mention, and users not in the server
 - Server Commands
   - [x] Serverinfo
     - Displays information about the current server
 - Bot Commands
   - [x] Ping
     - Displays both API and regular latency
     - Can ping a Minecraft server with a Raknet ping (short query) for basic information
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
 - Module System
   - [x] Enable
     - Enable a command/category
   - [x] Disable
     - Disable a command/category
   - [x] List
     - List all disabled commands/categories

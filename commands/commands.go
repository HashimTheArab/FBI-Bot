package commands

import (
	"github.com/Jviguy/SpeedyCmds"
	"github.com/Jviguy/SpeedyCmds/command"
	"github.com/Jviguy/SpeedyCmds/command/commandMap"
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"github.com/prim69/fbi-bot/utils"
	"github.com/prim69/fbi-bot/utils/settings"
	"strings"
	"time"
)

func init() {
	commands = map[string]*command.Command{
		"help":       {"Provides a list of commands", "help <command>", CategoryGeneral, nil, HelpCommand},
		"avatar":     {"Displays a users avatar", "avatar <user>", CategoryUser, []string{"av"}, AvatarCommand},
		"ping":       {"Displays the bots latency", "ping", CategoryBot, []string{"latency"}, PingCommand},
		"snipe":      {"Snipe a deleted message", "snipe [number]", CategoryFun, nil, SnipeCommand},
		"editsnipe":  {"Snipe an edited message", "editsnipe [number]", CategoryFun, nil, EditSnipeCommand},
		"nuke":       {"Nuke a channel", "nuke", CategoryUtility, nil, NukeCommand},
		"about":      {"View information about the bot", "about", CategoryBot, []string{"info"}, StatsCommand},
		"purge":      {"Purge an amount of messages", "purge <amount>", CategoryUtility, []string{"clear"}, PurgeCommand},
		"play":       {"Plays a song", "play <name|link>", CategoryMusic, []string{"p"}, PlayCommand},
		"backup":     {"Backup a server template", "backup <name>", CategoryUtility, nil, BackupCommand},
		"load":       {"Load a server template", "load <name>", CategoryUtility, nil, LoadCommand},
		"query":      {"Query a minecraft server", "query <ip> [port]", CategoryMinecraft, nil, QueryCommand},
		"module":     {"Manage command modules", "module <enable:disable:list>", CategoryModules, []string{"modules", "m"}, ModuleCommand},
		"serverinfo": {"View information on the current server", "serverinfo", CategoryServer, nil, ServerCommand},
		"ask":        {"Answers a question using AI", "ask <question>", CategoryWeb, nil, AskCommand},
		"lookup":     {"View data of a specific Xbox account", "lookup <name>", CategoryWeb, nil, LookupCommand},
		"whois":      {"View information about a user", "whois <user>", CategoryUser, nil, WhoIsCommand},
		"ban":        {"Ban another user", "ban <user> [reason]", CategoryModeration, nil, BanCommand},
	}
}

var commands map[string]*command.Command
var handler *SpeedyCmds.PremadeHandler
var fields []*discordgo.MessageEmbedField

var Categories = [...]string{CategoryGeneral, CategoryServer, CategoryUser, CategoryBot, CategoryFun, CategoryMusic, CategoryMinecraft, CategoryWeb, CategoryModeration, CategoryUtility, CategoryModules}
var UpTime time.Time

const (
	CategoryGeneral    = "General"
	CategoryFun        = "Fun"
	CategoryMusic      = "Music"
	CategoryWeb        = "Web"
	CategoryUser       = "User"
	CategoryServer     = "Server"
	CategoryBot        = "Bot"
	CategoryUtility    = "Utility"
	CategoryModules    = "Modules"
	CategoryModeration = "Moderation"
	CategoryMinecraft  = "Minecraft"
)

func RegisterAll(session *discordgo.Session) {
	UpTime = time.Now()
	handler = SpeedyCmds.New(session, commandMap.New(), true, ">")
	for name, c := range commands {
		c.Usage = handler.Prefix + c.Usage
		handler.GetCommandMap().RegisterCommand(name, *c, true)
	}

	for _, name := range settings.Data.DisabledCommands {
		if _, ok := commands[strings.ToLower(name)]; ok {
			handler.GetCommandMap().Disable(name)
		}
	}

	for _, name := range Categories {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   name,
			Value:  "None",
			Inline: false,
		})
	}

	for cname, Struct := range commands {
		for _, field := range fields {
			if field.Name == Struct.Category {
				if field.Value == "None" {
					field.Value = "> **" + cname + "** `" + Struct.Description + "`\n"
				} else {
					field.Value += "> **" + cname + "** `" + Struct.Description + "`\n"
				}
			}
		}
	}

}

func GetCommand(name string) *command.Command {
	return commands[name]
}

func GetHandler() *SpeedyCmds.PremadeHandler {
	return handler
}

func SendError(ctx ctx.Ctx, session *discordgo.Session, err string) error {
	_, e := session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
		Title:       "Error!",
		Description: err,
		Color:       utils.Red,
	})
	return e
}

func SendEmbed(ctx ctx.Ctx, session *discordgo.Session, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return session.ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)
}

// hasPermission checks if the given ctx has a specific permission.
func hasPermission(ctx ctx.Ctx, session *discordgo.Session, perm int64) bool {
	p, err := session.State.UserChannelPermissions(ctx.GetAuthor().ID, ctx.GetChannel().ID)
	if err != nil {
		_ = SendError(ctx, session, "Failed to retrieve user permissions! Error: "+err.Error())
		return false
	}
	if (p & perm) == 0 {
		_ = SendError(ctx, session, "You do not have permission to use this command.")
		return false
	}
	return true
}

func isOwner(ctx ctx.Ctx, session *discordgo.Session) bool {
	if ctx.GetAuthor().ID != utils.OwnerID {
		_ = SendError(ctx, session, "You do not have permission to use this command.")
		return false
	}
	return true
}

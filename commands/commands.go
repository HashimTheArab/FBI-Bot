package commands

import (
	"github.com/Jviguy/SpeedyCmds"
	"github.com/Jviguy/SpeedyCmds/command"
	"github.com/Jviguy/SpeedyCmds/command/commandMap"
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"time"
)

var Commands = map[string]command.Command{
	"avatar": {"Displays a users avatar", "avatar <user>", CategoryUser, []string{"av"}, AvatarCommand},
	"ping": {"Displays the bots latency", "ping", CategoryBot, []string{"latency"}, PingCommand},
	"snipe": {"Snipe a deleted message", "snipe [number]", CategoryFun, []string{}, SnipeCommand},
	"editsnipe": {"Snipe an edited message", "editsnipe [number]", CategoryFun, []string{}, EditSnipeCommand},
	"nuke": {"Nuke a channel", "nuke", CategoryUtility, []string{}, NukeCommand},
	"stats": {"View information about the bot", "stats", CategoryBot,[]string{"info"}, StatsCommand},
	"purge": {"Purge an amount of messages", "purge <amount>", CategoryUtility, []string{"clear"}, PurgeCommand},
	"play": {"Plays a song", "play <name|link>", CategoryMusic, []string{"p"}, PlayCommand},
	"backup": {"Backup a server template", "backup <name>", CategoryUtility, []string{}, BackupCommand},
	"load": {"Load a server template", "load <name>", CategoryUtility, []string{}, LoadCommand},
}

var Handler *SpeedyCmds.PremadeHandler
var Fields []*discordgo.MessageEmbedField

var Categories = []string{CategoryGeneral, CategoryFun, CategoryUser, CategoryBot, CategoryModeration, CategoryMusic, CategoryUtility}

var prefix = ""

var UpTime time.Time

const (
	CategoryGeneral    = "General"
	CategoryFun        = "Fun"
	CategoryMusic      = "Music"
	CategoryUser	   = "User"
	CategoryServer	   = "Server"
	CategoryBot        = "Bot"
	CategoryUtility    = "Utility"
	CategoryModeration = "Moderation"
)

func RegisterAll(session *discordgo.Session) {
	UpTime = time.Now()
	Commands["help"] = command.Command{Description: "Provides a list of commands", Usage: "help <command>", Category: CategoryGeneral, Aliases: []string{}, Execute: HelpCommand}
	Handler = SpeedyCmds.New(session, commandMap.New(), true, ">")
	prefix = Handler.Prefix
	for name, Struct := range Commands {
		Handler.GetCommandMap().RegisterCommand(name, Struct, true)
	}

	for _, name := range Categories {
		Fields = append(Fields, &discordgo.MessageEmbedField{
			Name:   name,
			Value:  "None",
			Inline: false,
		})
	}

	for cname, Struct := range Commands {
		for _, field := range Fields {
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

func SendInvalidUsage(ctx ctx.Ctx, session *discordgo.Session, usage string) {
	embed := discordgo.MessageEmbed{
		Title:       "Invalid Usage!",
		Description: "Usage: " + prefix + usage,
		Color:       0x9E1F44,
	}
	_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &embed)
}

func SendError(ctx ctx.Ctx, session *discordgo.Session, err string) {
	embed := discordgo.MessageEmbed{
		Title:       "Error!",
		Description: err,
		Color:       0x9E1F44,
	}
	_, _ = session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &embed)
}

func SendEmbed(ctx ctx.Ctx, session *discordgo.Session, embed discordgo.MessageEmbed) (*discordgo.Message, error) {
	return session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &embed)
}
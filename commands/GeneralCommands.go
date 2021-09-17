package commands

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func HelpCommand(ctx ctx.Ctx, session *discordgo.Session) error {
	if len(ctx.GetArgs()) < 1 {
		_, _ = SendEmbed(ctx, session, &discordgo.MessageEmbed{
			Title: "FBI Bot Commands",
			Description: handler.Prefix + "help <command name>",
			Color: 0x93c993,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Bot made by prim#0419",
			},
			Fields: fields,
		})
		return nil
	}

	if c, ok := commands[ctx.GetArgs()[0]]; ok {
		aliases := "None"
		if len(c.Aliases) > 0 {
			aliases = strings.Join(c.Aliases, ", ")
		}
		_, _ = SendEmbed(ctx, session, &discordgo.MessageEmbed{
			Description: "**Command Name:** " + ctx.GetArgs()[0] + "\n**Aliases:**: " + aliases + "\n**Description:** " + c.Description + "\n**Usage:** " + c.Usage,
			Color: 0x00B9CF,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Syntax <> = required, [] = optional",
			},
		})
		return nil
	}
	return SendError(ctx, session, "The command `" + ctx.GetArgs()[0] + "` does not exist!")
}

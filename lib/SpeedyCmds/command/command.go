package command

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Description string
	Usage       string
	Category    string
	Aliases     []string
	Execute     func(ctx ctx.Ctx, session *discordgo.Session) error
}

func (c *Command) SendUsage(ctx ctx.Ctx, session *discordgo.Session) error {
	_, err := session.ChannelMessageSendEmbed(ctx.GetChannel().ID, &discordgo.MessageEmbed{
		Title:       "Invalid Usage!",
		Description: "Usage: " + c.Usage,
		Color:       16711680,
	})
	return err
}
